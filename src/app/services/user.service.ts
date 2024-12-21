import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { User } from '../interfaces/user.interface';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class UserService {
  private apiUrl = 'http://localhost:8080/api/users';

  constructor(private http: HttpClient) {}

  createUser(userData: string): Observable<User> {
    const user = {
      data: userData
    };
    return this.http.post<User>(this.apiUrl, user);
  }
}