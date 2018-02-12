import { TestBed, inject } from '@angular/core/testing';

import { MCQService } from './MCQ.service';

describe('MCQService', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [MCQService]
    });
  });

  it('should be created', inject([MCQService], (service: MCQService) => {
    expect(service).toBeTruthy();
  }));
});
