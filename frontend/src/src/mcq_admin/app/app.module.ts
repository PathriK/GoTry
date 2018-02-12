import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { HttpModule } from '@angular/http';
import { RouterModule } from '@angular/router';

import { AppComponent } from './app.component';

import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { MdTabsModule } from '@angular/material';
import { MdListModule } from '@angular/material';
import { MdIconModule } from '@angular/material';
import { MdButtonModule } from '@angular/material';
import { MdTooltipModule } from '@angular/material';
import { MdSidenavModule } from '@angular/material';
import { MdToolbarModule } from '@angular/material';
import { MatInputModule } from '@angular/material';
import { MatRadioModule } from '@angular/material';
import { FormsModule } from '@angular/forms';

import { McqsComponent } from './mcqs/mcqs.component';
import { AddComponent } from './add/add.component';
import { RedirectComponent } from './redirect/redirect.component';

import { MCQService } from './mcq/mcq.service';

@NgModule({
  declarations: [
    AppComponent,
    McqsComponent,
    AddComponent,
    RedirectComponent
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
    MdListModule,
    MdIconModule,
    MdButtonModule,
    MdTooltipModule,
    MdSidenavModule,
    MdToolbarModule,
    MatInputModule,
    MatRadioModule,
    FormsModule
  ],
  providers: [MCQService],
  bootstrap: [AppComponent]
})
export class AppModule { }
