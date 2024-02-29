import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { CookieService } from 'ngx-cookie-service';
import { UserResponse } from 'src/models/user/userResponse';
import { UserService } from 'src/services/user/user.service';

@Component({
  selector: 'app-login-email',
  templateUrl: './login-email.component.html',
  styleUrls: ['./login-email.component.scss']
})
export class LoginEmailComponent implements OnInit {
  public form!: FormGroup;
  public emailNotFound = false;

  get f(): any {
    return this.form.controls
  }


  constructor(
    private userService: UserService,
    private fb: FormBuilder,
    private router: Router,
    private cookieService: CookieService) { }

  ngOnInit(): void {
    this.validation()
    localStorage.clear();
    this.cookieService.deleteAll()

  }

  public searchUserByEmail() {
    this.emailNotFound = false;
    if (this.form.valid) {
      const email = this.form.get('email')?.value;

      this.userService.getUserByEmail(email)
        .subscribe({
            next: (response: UserResponse) => {
              localStorage.setItem('email', email);
              this.router.navigate(['/auth/login/enter-password']);
            },
            error: (error: any) => {
              if (error.status === 404) {
                this.emailNotFound = true;
              }
            }

        })
    }
  }


  public validation(): void {
    this.form = this.fb.group({
      email: ['', [Validators.required, Validators.email]]
    })
  }

  public cssValidator(campoForm: FormControl): any {
    return { 'is-invalid': (campoForm?.errors && (campoForm?.touched || campoForm?.dirty)) || this.emailNotFound };
  }

}
