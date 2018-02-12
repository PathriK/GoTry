export class MCQ {
  question: string;
  options: string[];
  answer: number;

  constructor() {
    this.question = '';
    this.answer = 0;
    this.options = Array(4).fill('');
  }
}
