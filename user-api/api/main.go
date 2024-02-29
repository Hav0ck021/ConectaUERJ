package main

import (
	"fmt"
	"github.com/OVillas/user-api/api/handler"
	"github.com/OVillas/user-api/config"
	"github.com/OVillas/user-api/middleware"
	"github.com/OVillas/user-api/repository"
	"github.com/OVillas/user-api/service"
	"github.com/labstack/echo/v4"
	Middleware "github.com/labstack/echo/v4/middleware"
)

func main() {
	config.Load()
	e := echo.New()

	e.Use(Middleware.CORSWithConfig(Middleware.CORSConfig{
		AllowOrigins: []string{config.FrontendURL},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	configureUserRoutes(e)
	configureAuthenticationRoutes(e)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.Port)))
}

func configureUserRoutes(e *echo.Echo) {
	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository)
	emailService := service.NewEmailService("cineZuka", config.EmailSender, config.EMailSenderPassword)
	authenticationService := service.NewAuthenticationService(userRepository, emailService)
	userHandler := handler.NewUserHandler(userService, authenticationService)

	group := e.Group("v1/user")
	group.POST("", userHandler.Create)
	group.GET("", userHandler.GetAll)
	group.GET("/:id", userHandler.GetById)
	group.GET("/name", userHandler.GetByName)
	group.GET("/email", userHandler.GetByEmail)
	group.PUT("/:id", userHandler.Update, middleware.CheckLoggedIn)
	group.DELETE("/:id", userHandler.Delete, middleware.CheckLoggedIn)
}

func configureAuthenticationRoutes(e *echo.Echo) {
	userRepository := repository.NewUserRepository()
	emailService := service.NewEmailService("cineZuka", config.EmailSender, config.EMailSenderPassword)
	authenticationService := service.NewAuthenticationService(userRepository, emailService)
	authenticationHandler := handler.NewAuthenticationHandler(authenticationService)

	group := e.Group("v1/authentication")
	group.POST("/login", authenticationHandler.Login)
	group.PATCH("/user/:userId/password", authenticationHandler.UpdatePassword, middleware.CheckLoggedIn)
	group.PATCH("/ConfirmEmail", authenticationHandler.ConfirmEmail)

}
