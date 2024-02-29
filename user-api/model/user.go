package model

import (
	"errors"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

var (
	ErrHashPassword             = errors.New("error trying hashed password")
	ErrUserAlreadyRegistered    = errors.New("there is already a registered user with this email")
	ErrCreateUser               = errors.New("error to create user")
	ErrGetUser                  = errors.New("error to get user")
	ErrConvertUserPayLoadToUser = errors.New("error to create id from new user")
	ErrInvalidId                = errors.New("the id passed is invalid")
	ErrUserNotFound             = errors.New("user not found")
	ErrDeleteUser               = errors.New("error to delete user")
	ErrSameEmail                = errors.New("the email cannot be the same as the previous one")
)

type User struct {
	Id               string    `gorm:"column:Id"`
	Name             string    `gorm:"column:Name"`
	Email            string    `gorm:"column:Email"`
	Password         string    `gorm:"column:Password"`
	IsEmailConfirmed bool      `gorm:"column:IsEmailConfirmed"`
	CreatedAt        time.Time `gorm:"column:CreatedAt"`
	LastModified     time.Time `gorm:"column:LastModified"`
}

type UserPayLoad struct {
	Name     string `json:"name,omitempty" validate:"required,min=1,max=75"`
	Email    string `json:"email,omitempty" validate:"required,email"`
	Password string `json:"password,omitempty" validate:"required,min=6,containsany=!@#&?"`
}

type UserUpdatePayLoad struct {
	Name  string `json:"name,omitempty" validate:"min=1,max=75"`
	Email string `json:"email,omitempty"`
}

type UserResponse struct {
	Id               string
	Name             string
	Email            string
	IsEmailConfirmed bool
	CreatedAt        string
	LastModified     string
}

type UserHandler interface {
	Create(c echo.Context) error
	GetById(c echo.Context) error
	GetByName(c echo.Context) error
	GetByEmail(c echo.Context) error
	GetAll(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}

type UserService interface {
	Create(userPayLoad UserPayLoad) error
	GetById(id string) (*UserResponse, error)
	GetByName(name string) ([]UserResponse, error)
	GetByEmail(email string) (*UserResponse, error)
	GetAll() ([]UserResponse, error)
	Update(id string, userUpdate UserUpdatePayLoad) error
	Delete(id string) error
}

type UserRepository interface {
	Create(user User) error
	GetById(id string) (*User, error)
	GetByName(name string) ([]User, error)
	GetByEmail(email string) (*User, error)
	GetAll() ([]User, error)
	Update(id string, user User) error
	Delete(id string) error
	UpdatePassword(id string, password string) error
	UpdateConfirmedEmail(id string) error
}

func (upl *UserPayLoad) Validate() error {
	validate := validator.New()
	return validate.Struct(upl)
}

func (uu *UserUpdatePayLoad) Validate() error {
	validate := validator.New()
	return validate.Struct(uu)
}

func (upl *UserPayLoad) ToUser(hashedPassword string) (*User, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	return &User{
		Id:       id.String(),
		Name:     upl.Name,
		Email:    upl.Email,
		Password: hashedPassword,
	}, nil
}

func (uu *UserUpdatePayLoad) ToUser() *User {
	return &User{
		Name:  uu.Name,
		Email: uu.Email,
	}
}

func (u *User) ToUserResponse() *UserResponse {
	return &UserResponse{
		Id:               u.Id,
		Name:             u.Name,
		Email:            u.Email,
		IsEmailConfirmed: u.IsEmailConfirmed,
		CreatedAt:        u.CreatedAt.Format("2006-01-02 15:04:05"),
		LastModified:     u.LastModified.Format("2006-01-02 15:04:05"),
	}
}
