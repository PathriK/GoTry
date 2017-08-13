import { Component } from '@angular/core';
import { Home } from "./home"
import { TabServiceService } from './tab-service.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  isReady:boolean = false;
  home:Home;

  constructor(private tabService: TabServiceService) { }

  getHomeContents(): void {
    this.tabService
    .getContents()
    .then(home => {
      this.home = home;
      this.isReady = true;
    });
  }

  ngOnInit(): void {
    this.getHomeContents();
  }
}
