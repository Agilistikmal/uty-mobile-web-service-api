package main

import (
	"log"
	"net/http"

	"github.com/agilistikmal/uty-mobile-web-service-api/internal/app/delivery/rest"
	"github.com/agilistikmal/uty-mobile-web-service-api/internal/app/delivery/rest/route"
	"github.com/agilistikmal/uty-mobile-web-service-api/internal/app/repository"
	"github.com/agilistikmal/uty-mobile-web-service-api/internal/app/service"
	"github.com/agilistikmal/uty-mobile-web-service-api/internal/infrastructure/config"
	"github.com/agilistikmal/uty-mobile-web-service-api/internal/infrastructure/database"
	"github.com/go-playground/validator/v10"
)

func main() {
	config.NewConfig()

	db := database.NewDatabase()
	validate := validator.New()

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository, validate)

	otpRepository := repository.NewOTPRepository(db)
	otpService := service.NewOTPService(otpRepository, userRepository)

	// REST Handler
	userHandler := rest.NewUserHandler(userService, otpService)
	otpHandler := rest.NewOTPHandler(otpService)

	routes := route.NewRoutes(userHandler, otpHandler)
	routes.Init()

	log.Println("Running on http://localhost:8080")
	http.ListenAndServe(":8080", routes.Mux)
}
