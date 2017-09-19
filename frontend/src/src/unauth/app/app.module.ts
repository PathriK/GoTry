import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { HttpModule }    from '@angular/http';
import { RouterModule }   from '@angular/router';

import { AppComponent } from './app.component';
import { SpaceReplacePipe } from './space-replace.pipe';

import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { MdTabsModule } from '@angular/material';

import { HomeComponent } from './home/home.component';
import { GalleryComponent } from './gallery/gallery.component';
import { CoursesComponent } from './courses/courses.component';
import { ContactComponent } from './contact/contact.component';
import { RedirectComponent } from './redirect/redirect.component';

@NgModule({
  declarations: [
    AppComponent,
    SpaceReplacePipe,
    HomeComponent,
    GalleryComponent,
    CoursesComponent,
    ContactComponent,
    RedirectComponent
  ],
  imports: [
    BrowserModule,
    HttpModule,
    RouterModule.forRoot([
      {
        path: 'home',
        component: HomeComponent
      },
      {
        path: 'gallery',
        component: GalleryComponent
      },
      {
        path: 'courses',
        component: CoursesComponent
      },
      {
        path: 'contact_us',
        component: ContactComponent
      },
      {
        path: 'test_demo',
        component: RedirectComponent
      },
      {
        path: '**',
        redirectTo: '/home',
        pathMatch: 'full'
      }
    ]),
    BrowserAnimationsModule,
    MdTabsModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
