import { ComponentFixture, TestBed } from '@angular/core/testing';

import { LoginPasswordComponent } from './login-password.component';

describe('LoginPasswordComponent', () => {
  let component: LoginPasswordComponent;
  let fixture: ComponentFixture<LoginPasswordComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [LoginPasswordComponent]
    });
    fixture = TestBed.createComponent(LoginPasswordComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
