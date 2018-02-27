import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';

import { Observable } from 'rxjs/Observable';
import { of } from 'rxjs/observable/of';
import { catchError, map, tap } from 'rxjs/operators';

import { Service } from './service';

const httpOptions = {
  headers: new HttpHeaders({ 'Content-Type': 'application/json' })
};

interface ServerData {
  services: Service[];
}

@Injectable()
export class DirectoryService {

  private apiUrl = 'api/directory/list-services';

  constructor(
    private http: HttpClient
  ) { }

  getServices(): Observable<Service[]> {
    return this.http.get<ServerData>(this.apiUrl)
      .pipe(
        map(res => res.services),
        catchError(this.handleError('getServices', []))
      );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T> (operation = 'operation', result?: T) {
    return (error: any): Observable<T> => {
      console.error(error); // log to console instead

      // Let the app keep running by returning an empty result.
      return of(result as T);
    };
  }
}
