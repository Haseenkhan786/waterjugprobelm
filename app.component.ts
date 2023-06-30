import { HttpClient,HttpClientModule } from '@angular/common/http';
import { Component, ViewChild } from '@angular/core';
import { NgForm } from '@angular/forms';


@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})

export class AppComponent {


  constructor(private http: HttpClient) {}


  name: string='';
  gender:string='';
  fromDate:string='';
  toDate:string='';
  phone: string='';
  selectedFile:any= File;
  email:string='';
  onFileSelected(event: any) {
    this.selectedFile = event.target.files[0];
  }

  @ViewChild('formData', { static: false }) formData:any;



  submitForm()
   {
     const formData = new FormData();
    formData.append('full_name', this.name);
    formData.append('gender', this.gender);
    formData.append('from_date', this.fromDate);
    formData.append('to_date', this.toDate);
    formData.append('phone', this.phone);
    formData.append('upload_resume', this.selectedFile.name);
    formData.append('email', this.email);




    if (this.formData.valid) 
    {
      // Perform actions when the form is valid and submitted
      console.log('Form submitted successfully');
      console.log('Name:', this.name);
      console.log('Gender:', this.gender);
      console.log('From Date:', this.fromDate);
      console.log('To Date:', this.toDate);
      console.log('Phone:', this.phone);
      console.log('File:', this.selectedFile);
      console.log('Email:', this.email);

      this.http.post<any>('http://localhost:8080/postemployees', formData)
      .subscribe(
        response => {
          // Handle the response from the server
          console.log('Form submitted successfully', response);
        },
        error => {
          // Handle any errors that occur during the request
          console.error('Form submission failed', error);
        }
      );
    }
    
    else
     {
      // Handle form validation errors
      console.log('Form has validation errors');
    }

// to check the resume file

 
}




} 