import { Component } from '@angular/core';
import { Admin } from "./admin"
@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  isReady:boolean = false;
  admin:Admin;

  constructor() {}

  getAdminContents(): void {
    this.admin = {
      title: 'MCQ Admin',
      tabs: [
        {
          name: 'MCQ List',
          url: 'list'
        }
      ]
    };
    this.isReady = true;
  }

  ngOnInit(): void {
    this.getAdminContents();
  }
}
