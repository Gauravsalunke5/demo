import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { FormGroup, FormBuilder } from '@angular/forms';
import { Certificate, Record } from './shared/models';
@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})


export class AppComponent implements OnInit {

  public queryForm: FormGroup;
  public recordForm: FormGroup;
  public transferForm: FormGroup;
  public currentCertificate: Certificate = new Certificate({});
  public all_cert: Certificate[] = [];
  public PRno: any;
  public cert_CName: any;
  public cert_Seatno: any;
  public cert_examination: any;
  public cert_YOP: any;
  public cert_sububject: any;
  public data: string;


  title = 'SecureCert Demo';

  constructor(
    private _httpClient: HttpClient,
    private _fb: FormBuilder
  ) { }

  ngOnInit() {
    this.queryAllCerts();

    this.queryForm = this._fb.group({
      seatNo: null
    });

    this.recordForm = this._fb.group({
      cert_PRno: null,
      cert_CName: null,
      cert_Seatno: null,
      cert_examination: null,
      cert_YOP: null,
      cert_sububject: null
    });


    this.transferForm = this._fb.group({
      seatNo:null,
      studentName:null
    });

  }





  queryAllCerts() {
   // localStorage.setItem('username', 'amogh');
    this._httpClient.get(`http://localhost:8000/certificates`).subscribe((res: any[]) => {
      this.all_cert = res.map(cert => new Certificate(cert));
    });
  }

  queryCert() {
    console.log("jdbad");
    this._httpClient.get(`http://localhost:8000/certificates/${this.queryForm.get('seatNo').value}`).subscribe(res => {
      this.currentCertificate.record = new Record(res);

    });

  }

  recordCert() {
    this._httpClient.post(`http://localhost:8000/certificates`, this.recordForm.value).subscribe((data: any) => {
      console.log(data);
      this.recordForm.reset();
    });
  }




  transferCertificate() {
    this._httpClient.post(`http://localhost:8000/certificates/${this.transferForm.get('seatNo').value}/transferName`, this.transferForm.value).subscribe((data: any) => {
      console.log(data);
    });
  }

  

}
