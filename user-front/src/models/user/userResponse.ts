export class UserResponse {
  id: string;
  name: string;
  email: string;
  isEmailConfirmed: boolean;
  createdAt: string;
  lastModified: string;

  constructor(
    id: string,
    name: string,
    email: string,
    isEmailConfirmed: boolean,
    createdAt: string,
    lastModified: string
  ) {
    this.id = id;
    this.name = name;
    this.email = email;
    this.isEmailConfirmed = isEmailConfirmed;
    this.createdAt = createdAt;
    this.lastModified = lastModified;
  }
}
