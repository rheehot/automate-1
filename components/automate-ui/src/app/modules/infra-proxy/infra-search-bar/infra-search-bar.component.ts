import {
  Component,
  EventEmitter,
  Output,
  OnInit,
  ViewChild,
  ElementRef
} from '@angular/core';

@Component({
  selector: 'app-infra-search-bar',
  templateUrl: './infra-search-bar.component.html',
  styleUrls: ['./infra-search-bar.component.scss']
})

export class InfraSearchBarComponent implements OnInit {

  inputText = '';

  @Output() searchButtonClick: EventEmitter<any> = new EventEmitter<any>();
  @ViewChild('search_box', { static: true }) inputField: ElementRef;

  ngOnInit() {}

  handleFiltersClick(currentText: string): void {
    this.searchButtonClick.emit(currentText);
  }

  pressEnter(currentText: string): void {
    this.searchButtonClick.emit(currentText);
  }

  handleInput(key, currentText): void {
    switch (key.toLowerCase()) {
      case 'enter':
        this.pressEnter(currentText);
        break;
    }
  }

  getFilterText(): string {
    return 'Search';
  }

}
