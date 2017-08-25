import { Pipe, PipeTransform } from '@angular/core';

@Pipe({
  name: 'spaceReplace'
})
export class SpaceReplacePipe implements PipeTransform {

  transform(value: string, withVal=''): string {
    let newValue = value.replace(/ /g, withVal);
    return `${newValue}`;
  }

}
