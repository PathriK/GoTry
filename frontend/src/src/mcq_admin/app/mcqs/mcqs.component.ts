import { Component, OnInit } from '@angular/core';
import {MCQ} from './mcq'

@Component({
  selector: 'app-mcqs',
  templateUrl: './mcqs.component.html',
  styleUrls: ['./mcqs.component.css']
})
export class McqsComponent implements OnInit {

  isReady:boolean = false;
  mcqs:MCQ[];

  constructor() { }

  getMCQs(){
    this.mcqs = [{
      question: "Sample Question 1",
      choices: [
        "Ch1",
        "ch2"
      ],
      answer: 1
    },
    {
      question: "Sample Question 2",
      choices: [
        "Ch1",
        "ch2"
      ],
      answer: 2
    }];
    this.isReady = true;
  }

  ngOnInit() {
    this.getMCQs();
  }

}
