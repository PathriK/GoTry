import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { HttpModule }    from '@angular/http';
import { RouterModule }   from '@angular/router';

import { AppComponent } from './app.component';

import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { MdTabsModule } from '@angular/material';
import {MdListModule} from '@angular/material';

import {McqsComponent} from './mcqs/mcqs.component'
import {AddComponent} from './add/add.component'

@NgModule({
  declarations: [
    AppComponent,
    McqsComponent,
    AddComponent
  ],
  imports: [
    BrowserModule,
    HttpModule,
    RouterModule.forRoot([
      {
        path: 'list',
        component: McqsComponent
      },
      {
        path: 'add',
        component: AddComponent
      },
      {
        path: '**',
        redirectTo: '/list',
        pathMatch: 'full'
      }
    ]),
    BrowserAnimationsModule,
    MdTabsModule,
    MdListModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
