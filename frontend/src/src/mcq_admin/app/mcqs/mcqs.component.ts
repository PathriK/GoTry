import { Component, OnInit } from '@angular/core';

import { MCQ } from '../mcq/mcq';
import { MCQService } from '../mcq/mcq.service';

@Component({
  selector: 'app-mcqs',
  templateUrl: './mcqs.component.html',
  styleUrls: ['./mcqs.component.css']
})
export class McqsComponent implements OnInit {

  isReady = false;
  mcqs: MCQ[];

  constructor(
    private mcqService: MCQService ) { }

  getMCQs() {
    this.mcqService.getMcqs().then(
      mcqs => {
        this.mcqs = mcqs;
        this.isReady = true;
      });
  }

  ngOnInit() {
    this.getMCQs();
  }

}
