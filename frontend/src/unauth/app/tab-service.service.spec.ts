import { TestBed, inject } from '@angular/core/testing';

import { TabServiceService } from './tab-service.service';

describe('TabServiceService', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [TabServiceService]
    });
  });

  it('should be created', inject([TabServiceService], (service: TabServiceService) => {
    expect(service).toBeTruthy();
  }));
});
