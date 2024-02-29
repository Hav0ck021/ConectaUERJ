import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule } from '@angular/common/http';
import { ReactiveFormsModule } from '@angular/forms';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { AuthenticationComponent } from './components/authentication/authentication.component';
import { RegisterComponent } from './components/authentication/register/register.component';
import { HomeComponent } from './components/home/home.component';
import { LoginPasswordComponent } from './components/authentication/login-password/login-password.component';
import { LoginEmailComponent } from './components/authentication/login-email/login-email.component';
import { NavComponent } from './components/shared/nav/nav.component';
import { FooterComponent } from './components/shared/footer/footer.component';

@NgModule({
  declarations: [
    AppComponent,
    AuthenticationComponent,
    RegisterComponent,
    HomeComponent,
    LoginPasswordComponent,
    LoginEmailComponent,
    NavComponent,
    FooterComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    ReactiveFormsModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
