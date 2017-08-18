import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { HttpModule }    from '@angular/http';

import { AppComponent } from './app.component';
import { TabServiceService } from './tab-service.service';
import { SpaceReplacePipe } from './space-replace.pipe';

@NgModule({
  declarations: [
    AppComponent,
    SpaceReplacePipe
  ],
  imports: [
    BrowserModule,
    HttpModule
  ],
  providers: [TabServiceService],
  bootstrap: [AppComponent]
})
export class AppModule { }
