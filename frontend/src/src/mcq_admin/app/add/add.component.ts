import { Component, OnInit } from '@angular/core';
import { MCQ } from '../mcq/mcq';

@Component({
  selector: 'app-add',
  templateUrl: './add.component.html',
  styleUrls: ['./add.component.css']
})
export class AddComponent implements OnInit {

  mcq: MCQ;

  constructor() {
   }

   onSubmit(data) {
     data = JSON.stringify(data, null, 2);
     console.log(data);
   }
   trackByFn(index, item) {
return index;
   }
  ngOnInit() {
    this.mcq = new MCQ();
  }

}
