import { ComponentFixture, TestBed } from '@angular/core/testing';

import { LoginEmailComponent } from './login-email.component';

describe('LoginEmailComponent', () => {
  let component: LoginEmailComponent;
  let fixture: ComponentFixture<LoginEmailComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [LoginEmailComponent]
    });
    fixture = TestBed.createComponent(LoginEmailComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
