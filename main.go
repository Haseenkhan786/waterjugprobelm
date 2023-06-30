package main

import ( 
	"github.com/gin-contrib/cors"
	"database/sql"
	"fmt"
	"net/http"
//    "mime/multipart"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
  //  "os"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "haseen"
	password = "1302001"
	dbname   = "postgres"
)

type EmployeeDetail struct {
	FullName      string `form:"full_name" json:"full_name"`
	Gender        string `form:"gender" json:"gender"`
	From_date     string `form:"from_date" json:"from_date"`
	To_date       string `form:"to_date" json:"to_date"`
	Phone         int `form:"phone" json:"phone"`
	// Upload_resume *multipart.FileHeader `form:"upload_resume" json:"upload_resume"`
	Upload_resume string `form:"upload_resume" json:"upload_resume"`
	Email         string `form:"email" json:"email"`
}
type SickEmployeeDetails struct {
	FullName      string `json:"full_name"`
	Gender        string `json:"gender"`
	From_date     string `json:"from_date"`
	To_date       string `json:"to_date"`
	Phone         int `json:"phone"`
	Sickfile	string `json:"upload_resume"`
	Email         string `json:"email"`
}

var employees = []EmployeeDetail{}

func main() {

	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:4200"} // Replace with the URL of your Angular application
	config.AllowHeaders = []string{"Origin", "Content-Type"}
	router.Use(cors.New(config))
	
	router.GET("/getemployees", getEmployees)

	router.POST("/postemployees", postEmployees)

	router.Run("localhost:8080")
}

// POST
func postEmployees(c *gin.Context) {
	var newEmployee EmployeeDetail

	if err := c.Bind(&newEmployee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	// employees = append(employees, newEmployee)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to the database"})
		return
	}
	defer db.Close()
	// fileLink := ""

    // if newEmployee.Upload_resume != nil {

    //     // Handle file upload and save the file to the desired folder (myfolder)

    //     filePath := "static/myfolder/" + newEmployee.Upload_resume.Filename

    //     if err := c.SaveUploadedFile(newEmployee.Upload_resume, filePath); err != nil {

    //         c.String(http.StatusInternalServerError, "Failed to save file on the server")

    //         return

    //     }

    //     fileLink = filePath

    // }

	insertDynStmt := `INSERT INTO "employee_detail" ("full_name", "gender",  "from_date", "to_date","phone", "upload_resume", "email") VALUES ($1, $2, $3, $4, $5, $6,$7)`

	_, err = db.Exec(insertDynStmt, newEmployee.FullName, newEmployee.Gender, newEmployee.From_date, newEmployee.To_date, newEmployee.Phone, newEmployee.Upload_resume, newEmployee.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert employee data into the database"})
		return
	}

	c.JSON(http.StatusCreated, newEmployee)
}

// GET
func getEmployees(c *gin.Context) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to the database"})
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM employee_detail")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch employee data from the database"})
		return
	}
	defer rows.Close()

	var employees []SickEmployeeDetails
	for rows.Next() {
		var emp SickEmployeeDetails
		err := rows.Scan(&emp.FullName, &emp.Gender, &emp.From_date, &emp.To_date, &emp.Phone,&emp.Sickfile , &emp.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan employee data from the database"})
			return
		}
		employees = append(employees, emp)
	}

	c.IndentedJSON(http.StatusOK, employees)
}

