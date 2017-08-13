import { Injectable } from '@angular/core';
import { Http } from '@angular/http';

import 'rxjs/add/operator/toPromise';

import { Home } from "./home";


@Injectable()
export class TabServiceService {

  private homeURL = '/api/unauth/home';  // URL to web api

  constructor(private http: Http) { }
  getContents(): Promise<Home> {
    return this.http.get(this.homeURL)
               .toPromise()
               .then(response => response.json() as Home)
               .catch(this.handleError);
  }

  private handleError(error: any): Promise<any> {
    console.error('An error occurred', error); // for demo purposes only
    return Promise.reject(error.message || error);
  }
}
