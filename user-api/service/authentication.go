package service

import (
	"fmt"
	"github.com/OVillas/user-api/model"
	"github.com/OVillas/user-api/util"
	"log/slog"
	"time"
)

var confirmationsCodes map[string]model.ConfirmationCode

func init() {
	confirmationsCodes = make(map[string]model.ConfirmationCode)
}

type authenticationService struct {
	userRepository model.UserRepository
	emailService   model.EmailService
}

func NewAuthenticationService(userRepository model.UserRepository, emailService model.EmailService) model.AuthenticationService {
	return &authenticationService{
		userRepository: userRepository,
		emailService:   emailService,
	}
}

func (a *authenticationService) Login(login model.Login) (string, error) {
	log := slog.With(
		slog.String("func", "Login"),
		slog.String("service", "authentication"))

	user, err := a.userRepository.GetByEmail(login.Email)
	if err != nil {
		log.Warn("Failed to obtain user by email")
		return "", model.ErrGetUser
	}

	if user == nil {
		log.Warn("User not found with this email: " + login.Email)
		return "", model.ErrUserNotFound
	}

	if err := CheckPassword(user.Password, login.Password); err != nil {
		log.Warn("invalid password for email: " + user.Email)
		return "", model.ErrPasswordNotMatch
	}

	token, err := util.CreateToken(*user)
	if err != nil {
		log.Error("error trying create token jwt. Error: ", err)
		return "", model.ErrGenToken
	}

	return token, nil
}

func (a *authenticationService) UpdatePassword(id string, updatePassword model.UpdatePassword) error {
	log := slog.With(
		slog.String("func", "Login"),
		slog.String("service", "authentication"))

	user, err := a.userRepository.GetById(id)
	if err != nil {
		log.Error("failed to get user by id")
		return model.ErrGetUser
	}

	if user == nil {
		log.Warn("User not found with this id")
		return model.ErrUserNotFound
	}

	if err := CheckPassword(user.Password, updatePassword.Current); err != nil {
		log.Warn("current password not match ")
		return model.ErrPasswordNotMatch
	}

	newHashedPassword, err := Hash(updatePassword.New)
	if err != nil {
		log.Error("Error trying to hashed password")
		return model.ErrHashPassword
	}

	if err := a.userRepository.UpdatePassword(id, string(newHashedPassword)); err != nil {
		log.Error("Error: ", err)
		return model.ErrUpdatePassword
	}

	log.Info("Password updated successfully")
	return nil
}

func (a *authenticationService) SendConfirmationEmailCode(email string) error {
	log := slog.With(
		slog.String("func", "SendConfirmationEmailCode"),
		slog.String("service", "authentication"))

	otp := model.ConfirmationCode{
		Code:       util.GenerateOTP(6),
		ExpiryTime: time.Now().Add(time.Hour),
	}

	a.addOrUpdateConfirmationCode(email, otp)

	subject := "Confirmação de cadastro"
	content := fmt.Sprintf("<h1>Olá!</h1><p>Seu código de confirmação é: <h2><b>%s</b></h2></p>", otp.Code)
	to := []string{email}

	err := a.emailService.SendEmail(subject, content, to)
	if err != nil {
		log.Error("Errors: ", err)
		return model.ErrToSendConfirmationCode
	}
	log.Info("Confirmation send successfully")

	return nil
}

func (a *authenticationService) ConfirmEmail(confirmCodeEmail model.ConfirmCodeEmail) error {
	log := slog.With(
		slog.String("func", "ConfirmEmail"),
		slog.String("service", "authentication"))

	user, err := a.userRepository.GetByEmail(confirmCodeEmail.Email)
	if err != nil {
		log.Warn("Failed to obtain user by email")
		return model.ErrGetUser
	}

	if user == nil {
		log.Warn("User not found with this email: " + confirmCodeEmail.Email)
		return model.ErrUserNotFound
	}

	confirmationCode, ok := confirmationsCodes[confirmCodeEmail.Email]
	if !ok {
		log.Error("OTP not found with this email: " + confirmCodeEmail.Email)
		return model.ErrOTPNotFound
	}

	if time.Now().After(confirmationCode.ExpiryTime) {
		log.Warn("Token expired")
		return model.ErrInvalidOTP
	}

	if confirmationCode.Code != confirmCodeEmail.Code {
		log.Warn("incorrect token")
		return model.ErrInvalidOTP
	}

	if err := a.userRepository.UpdateConfirmedEmail(user.Id); err != nil {
		log.Error("Error: ", err)
		return err
	}

	log.Info("Confirmed email successfully")
	return nil
}

func (a *authenticationService) addOrUpdateConfirmationCode(email string, code model.ConfirmationCode) {
	if existingCode, ok := confirmationsCodes[email]; ok {
		existingCode.Code = code.Code
		existingCode.ExpiryTime = code.ExpiryTime
		confirmationsCodes[email] = existingCode
	} else {
		confirmationsCodes[email] = code
	}
}
