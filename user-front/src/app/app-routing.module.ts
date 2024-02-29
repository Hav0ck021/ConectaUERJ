import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AuthenticationComponent } from './components/authentication/authentication.component';
import { RegisterComponent } from './components/authentication/register/register.component';
import { HomeComponent } from './components/home/home.component';
import { LoginEmailComponent } from './components/authentication/login-email/login-email.component';
import { LoginPasswordComponent } from './components/authentication/login-password/login-password.component';

const routes: Routes = [

  { path: "home", component: HomeComponent},

  {
    path: "auth",
    component: AuthenticationComponent,

    children: [
      { path: "login/enter-email", component:LoginEmailComponent},
      { path: "login/enter-password", component: LoginPasswordComponent},
      { path: "register", component: RegisterComponent}
    ]
  }

];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
