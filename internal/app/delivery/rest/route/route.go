package route

import (
	"net/http"

	"github.com/agilistikmal/uty-mobile-web-service-api/internal/app/delivery/rest"
)

type Route struct {
	Mux *http.ServeMux

	UserHandler    *rest.UserHandler
	OTPHandler     *rest.OTPHandler
	PaymentHandler *rest.PaymentHandler
}

func NewRoutes(userHandler *rest.UserHandler, otpHandler *rest.OTPHandler, PaymentHandler *rest.PaymentHandler) *Route {
	return &Route{
		Mux:            http.NewServeMux(),
		UserHandler:    userHandler,
		OTPHandler:     otpHandler,
		PaymentHandler: PaymentHandler,
	}
}

func (r *Route) Init() {
	r.ProductRoutes()
}

func (r *Route) ProductRoutes() {
	// Mendaftarkan user handler dengan endpoint
	r.Mux.HandleFunc("POST /auth/register", r.UserHandler.Register)
	r.Mux.HandleFunc("POST /auth/login", r.UserHandler.Login)
	r.Mux.HandleFunc("GET /user/{username}", r.UserHandler.Find)

	r.Mux.HandleFunc("POST /auth/otp", r.OTPHandler.Verify)

	r.Mux.HandleFunc("POST /payment", r.PaymentHandler.Create)
	r.Mux.HandleFunc("GET /payment/id/{id}", r.PaymentHandler.FindByID)
	r.Mux.HandleFunc("GET /payment/reference_id/{reference_id}", r.PaymentHandler.FindByReferenceID)
}
