import { Injectable } from '@angular/core';
import { Http } from '@angular/http';

import 'rxjs/add/operator/toPromise';

import { MCQ } from './mcq';

@Injectable()
export class MCQService {
  private listUrl = '/api/admin/mcq/list';

  constructor(private http: Http) { }

  getMcqs(): Promise<MCQ[]> {
    return this.http.get(this.listUrl)
    .toPromise()
    .then(response => response.json().data as MCQ[])
    .catch(this.handleError);
  }

  private handleError(error: any): Promise<any> {
    console.error('An error occurred', error); // for demo purposes only
    return Promise.reject(error.message || error);
  }
}
