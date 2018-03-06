import { Component, OnInit, OnDestroy } from '@angular/core';

import { Service } from './service';
import { DirectoryService } from './directory.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent  implements OnInit, OnDestroy {
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
    .subscribe(services => {
      this.services = services;
      this.services.sort(
        function(a,b){
           if (a.organization_name !== b.organization_name) {
              return (a.organization_name > b.organization_name) ? 1 : -1;
           } else {
              return (a.name > b.name) ? 1 : -1;
           }
        });
      });
  }
}
