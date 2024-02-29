import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { CookieService } from 'ngx-cookie-service';
import { Login } from 'src/models/authentication/login';
import { AuthenticationService } from 'src/services/authentication/authentication.service';

@Component({
  selector: 'app-login-password',
  templateUrl: './login-password.component.html',
  styleUrls: ['./login-password.component.scss']
})
export class LoginPasswordComponent implements OnInit {
  private email!: string;
  public form!: FormGroup;
  public invalidPassword: boolean = false
  public invalidData: boolean = false

  get f(): any {
    return this.form.controls
  }


  constructor(
    private router: Router,
    private authService: AuthenticationService,
    private fb: FormBuilder,
    private cookieService: CookieService) { }

  ngOnInit(): void {
    this.validation()

    let email = localStorage.getItem('email')
    if (email) {
      this.email = email;
    } else {
      this.router.navigate(['/auth/login/enter-email']);
    }

    this.cookieService.deleteAll();
  }

  public getEmail(): string {
    return this.email
  }

  public validation(): void {
    this.form = this.fb.group({
      password: ['', [Validators.required]]
    })
  }


  public login(): void {
    this.invalidPassword = false;
    if (this.form.valid) {
      const password = this.form.get('password')?.value;
      const login = new Login(this.email, password)
      this.authService.Login(login)
        .subscribe({
          next: (response: string) => {
            this.cookieService.set("token", response)
            this.router.navigate(['/home']);
          },
          error: (error: any) => {
            if (error.status === 403) {
              this.invalidPassword = true;
            }

          }
        })

    }
  }

  public cssValidator(campoForm: FormControl): any {
    return { 'is-invalid': (campoForm?.errors && (campoForm?.touched || campoForm?.dirty)) || this.invalidPassword };
  }

}
