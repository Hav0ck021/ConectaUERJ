package model

import (
	"errors"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var (
	ErrPasswordNotMatch        = errors.New("invalid password")
	ErrGenToken                = errors.New("error to generate new token jwt")
	ErrUnexpectedSigningMethod = errors.New("unexpected signature method")
	ErrInvalidToken            = errors.New("token invalid")
	ErrIdNotFoundInPermissions = errors.New("error to get id in token")
	ErrIdIsNotAString          = errors.New("'id' field value is not a string")
	ErrUpdatePassword          = errors.New("error to update password")
	ErrToSendConfirmationCode  = errors.New("error to send confirmation code")
	ErrInvalidOTP              = errors.New("Wrong or expired OTP")
	ErrOTPNotFound             = errors.New("Not found OTP from email")
)

type RequestBody struct {
	Email string `json:"email,omitempty" validate:"required,email"`
}

type Login struct {
	Email    string `json:"email,omitempty" validate:"required,email"`
	Password string `json:"password,omitempty" validate:"required"`
}

type UpdatePassword struct {
	Password        string `json:"password,omitempty" validate:"required,min=6,containsany=!@#&?"`
	ConfirmPassword string `json:"confirmPassword,omitempty" validate:"required,min=6,containsany=!@#&?"`
}

type ConfirmationCode struct {
	Code       string
	ExpiryTime time.Time
}

type ConfirmCodeEmail struct {
	Email string `json:"email,omitempty" validate:"required,email"`
	Code  string `json:"code,omitempty" validate:"required"`
}

func CustomValidation(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	confirmPassword := fl.Parent().FieldByName("ConfirmPassword").String()
	return password == confirmPassword
}

func (l *Login) Validate() error {
	validate := validator.New()
	return validate.Struct(l)
}

func (up *UpdatePassword) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("custom", CustomValidation)
	return validate.Struct(up)
}

func (ce *ConfirmCodeEmail) Validate() error {
	validate := validator.New()
	return validate.Struct(ce)
}

func (rb *RequestBody) Validate() error {
	validate := validator.New()
	return validate.Struct(rb)
}

type AuthenticationHandler interface {
	Login(c echo.Context) error
	UpdatePassword(c echo.Context) error
	ConfirmEmail(c echo.Context) error
	ForgotPassword(c echo.Context) error
	ConfirmPasswordResetOtp(c echo.Context) error
}

type AuthenticationService interface {
	Login(login Login) (string, error)
	UpdatePassword(id string, updatePassword UpdatePassword) error
	SendOneTimePassword(email string) error
	CheckOneTimePassword(confirmCodeEmail ConfirmCodeEmail) error
	ConfirmEmail(confirmCodeEmail ConfirmCodeEmail) error
	ForgotPassword(email string) (string, error)
}
