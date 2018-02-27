import { Component, OnInit } from '@angular/core';

import { Service } from './service';
import { DirectoryService } from './directory.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent  implements OnInit {
  services: Service[];
  interval: any;

  constructor(private directoryService: DirectoryService) { }

  ngOnInit() {
    this.getServices();
    this.interval = setInterval(() => {
      this.getServices();
    }, 5000);
  }

  ngOnDestroy() {
    clearInterval(this.interval);
  }

  getServices(): void {
    this.directoryService.getServices()
    .subscribe(services => this.services = services);
  }
}
