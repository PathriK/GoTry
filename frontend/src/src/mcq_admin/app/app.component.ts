import { Component, OnInit } from '@angular/core';
import { Admin } from './admin';
@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent  implements OnInit {
  isReady = false;
  admin: Admin;

  constructor() {}

  getAdminContents(): void {
    this.admin = {
      title: 'MCQ Admin',
      tabs: [
        {
          name: 'MCQ List',
          url: 'list'
        },
        {
          name: 'MCQ Add',
          url: 'add'
        }
      ]
    };
    this.isReady = true;
  }

  ngOnInit() {
    this.getAdminContents();
  }
}
