import { Login } from './../../models/authentication/login';
import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { environment } from 'src/environments/environments';

@Injectable({
  providedIn: 'root'
})
export class AuthenticationService {
  private userAPI = environment.API

  constructor(private http: HttpClient) { }

  public Login(login: Login): Observable<string> {
    return this.http.post<string>(`${this.userAPI}/authentication/login`, login)
  }

}
