import { Component } from '@angular/core';
import { Home } from "./home"

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  isReady:boolean = false;
  home:Home;

  constructor() { }

  getHomeContents(): void {
    this.home = {
      title: 'My Page',
      tabs: [
        'Home', 'Gallery', 'Courses', 'Contact Us', 'Test Demo'
      ]
    };
    this.isReady = true;
  }

  ngOnInit(): void {
    this.getHomeContents();
  }
}
