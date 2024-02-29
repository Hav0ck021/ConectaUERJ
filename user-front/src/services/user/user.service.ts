import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { environment } from 'src/environments/environments';
import { UserResponse } from 'src/models/user/userResponse';

@Injectable({
  providedIn: 'root'
})
export class UserService {

  private userAPI = environment.API


  constructor(private http: HttpClient) { }

  public getUserByEmail(email: string): Observable<UserResponse> {
    return this.http.get<UserResponse>(`${this.userAPI}/user/email?e=${email}`)
  }

}
