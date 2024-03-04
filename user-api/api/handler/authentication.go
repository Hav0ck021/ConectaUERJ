package handler

import (
	"errors"
	"github.com/OVillas/user-api/util"
	"log/slog"
	"net/http"

	"github.com/OVillas/user-api/model"
	"github.com/labstack/echo/v4"
)

type authenticationHandler struct {
	authenticationService model.AuthenticationService
}

func NewAuthenticationHandler(authService model.AuthenticationService) model.AuthenticationHandler {
	return &authenticationHandler{
		authenticationService: authService,
	}
}

func (a *authenticationHandler) Login(c echo.Context) error {
	log := slog.With(
		slog.String("func", "Login"),
		slog.String("handler", "authentication"))

	var login model.Login
	if err := c.Bind(&login); err != nil {
		log.Warn("Failed to bind user data to model")
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	if err := login.Validate(); err != nil {
		log.Warn("Invalid login data")
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	token, err := a.authenticationService.Login(login)
	if err != nil && errors.Is(err, model.ErrPasswordNotMatch) {
		log.Warn("email or password invalid")
		return c.NoContent(http.StatusForbidden)

	}

	if err != nil && errors.Is(err, model.ErrUserNotFound) {
		log.Warn("User not found with email: " + login.Email)
		return c.NoContent(http.StatusNotFound)

	}

	if err != nil {
		log.Error("Error trying to call login service.")
		return c.JSON(http.StatusInternalServerError, err)
	}

	log.Info("login executed successfully")
	return c.JSON(http.StatusOK, token)
}

func (a *authenticationHandler) UpdatePassword(c echo.Context) error {
	log := slog.With(
		slog.String("func", "UpdatePassword"),
		slog.String("handler", "authentication"))

	userId, err := util.ExtractUserIdFromToken(c)
	if err != nil {
		log.Warn("err to get user if from token")
		return c.JSON(http.StatusUnauthorized, err)
	}

	var updatePassword model.UpdatePassword
	if err := c.Bind(&updatePassword); err != nil {
		log.Warn("Failed to bind user data to model")
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	if err := updatePassword.Validate(); err != nil {
		log.Warn("invalid user data")
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	err = a.authenticationService.UpdatePassword(userId, updatePassword)

	if err != nil && errors.Is(err, model.ErrUserNotFound) {
		log.Error("Error: ", err)
		return c.JSON(http.StatusNotFound, err)
	}

	if err != nil && errors.Is(err, model.ErrPasswordNotMatch) {
		log.Error("Error: ", err)
		return c.JSON(http.StatusUnauthorized, err)
	}

	log.Info("UpdatePassword executed successfully")
	return c.NoContent(http.StatusNoContent)
}

func (a *authenticationHandler) ForgotPassword(c echo.Context) error {
	log := slog.With(
		slog.String("func", "ForgotPassword"),
		slog.String("handler", "authentication"))

	var requestBody model.RequestBody
	if err := c.Bind(&requestBody); err != nil {
		log.Warn("Failed to bind user data to model")
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	email := requestBody.Email

	if err := a.authenticationService.SendOneTimePassword(email); err != nil {
		log.Error("Errors: ", err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	log.Info("One time password send successfully")
	return c.NoContent(http.StatusOK)
}

func (a *authenticationHandler) ConfirmPasswordResetOtp(c echo.Context) error {
	log := slog.With(
		slog.String("func", "ConfirmPasswordResetOtp"),
		slog.String("handler", "authentication"))

	var confirmCodeEmail model.ConfirmCodeEmail
	if err := c.Bind(&confirmCodeEmail); err != nil {
		log.Warn("Failed to bind confirmCodeData data to model")
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	if err := confirmCodeEmail.Validate(); err != nil {
		log.Warn("Invalid confirmCodeEmail data")
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	err := a.authenticationService.CheckOneTimePassword(confirmCodeEmail)

	if err != nil && errors.Is(err, model.ErrInvalidOTP) {
		log.Warn("Expired token or wrong token")
		return c.NoContent(http.StatusUnauthorized)
	}

	if err != nil {
		log.Error("Errors: ", err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	token, err := a.authenticationService.ForgotPassword(confirmCodeEmail.Email)
	if err != nil {
		log.Error("Error to generate reset password token", err)
		return c.NoContent(http.StatusInternalServerError)
	}

	log.Info("Token verified successfully")
	return c.JSON(http.StatusOK, token)
}

func (a *authenticationHandler) ConfirmEmail(c echo.Context) error {
	log := slog.With(
		slog.String("func", "ConfirmEmail"),
		slog.String("handler", "authentication"))

	var confirmCodeEmail model.ConfirmCodeEmail
	if err := c.Bind(&confirmCodeEmail); err != nil {
		log.Warn("Failed to bind confirmCodeData data to model")
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	if err := confirmCodeEmail.Validate(); err != nil {
		log.Warn("Invalid confirmCodeEmail data")
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	err := a.authenticationService.ConfirmEmail(confirmCodeEmail)

	if err != nil && errors.Is(err, model.ErrInvalidOTP) {
		log.Warn("Expired token or wrong token")
		return c.NoContent(http.StatusUnauthorized)
	}

	if err != nil {
		log.Error("Errors: ", err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	log.Info("e-mail confirmed successfully")
	return c.NoContent(http.StatusOK)
}
