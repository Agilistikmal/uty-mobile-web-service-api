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
	"github.com/spf13/viper"
	"github.com/xendit/xendit-go/v6"
)

func main() {
	config.NewConfig()

	db := database.NewDatabase()
	validate := validator.New()
	xenditClient := xendit.NewClient(viper.GetString("xendit.secret_key"))

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository, validate)

	otpRepository := repository.NewOTPRepository(db)
	otpService := service.NewOTPService(otpRepository)

	paymentRepository := repository.NewPaymentRepository(db)
	paymentService := service.NewPaymentService(xenditClient, paymentRepository, userRepository, validate)

	// REST Handler
	userHandler := rest.NewUserHandler(userService, otpService)
	otpHandler := rest.NewOTPHandler(otpService, userService)
	paymentHandler := rest.NewPaymentHandler(paymentService)

	routes := route.NewRoutes(userHandler, otpHandler, paymentHandler)
	routes.Init()

	log.Println("Running on http://localhost:8080")
	http.ListenAndServe(":8080", routes.Mux)
}
