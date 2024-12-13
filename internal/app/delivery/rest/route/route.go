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
	PostHandler    *rest.PostHandler
}

func NewRoutes(userHandler *rest.UserHandler, otpHandler *rest.OTPHandler, paymentHandler *rest.PaymentHandler, postHandler *rest.PostHandler) *Route {
	return &Route{
		Mux:            http.NewServeMux(),
		UserHandler:    userHandler,
		OTPHandler:     otpHandler,
		PaymentHandler: paymentHandler,
		PostHandler:    postHandler,
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
	r.Mux.HandleFunc("PATCH /user/{username}", r.UserHandler.Update)
	r.Mux.HandleFunc("DELETE /user/{username}", r.UserHandler.Delete)

	r.Mux.HandleFunc("POST /auth/otp", r.OTPHandler.Verify)

	r.Mux.HandleFunc("POST /payment", r.PaymentHandler.Create)
	r.Mux.HandleFunc("GET /payment/id/{id}", r.PaymentHandler.FindByID)
	r.Mux.HandleFunc("GET /payment/reference_id/{reference_id}", r.PaymentHandler.FindByReferenceID)

	r.Mux.HandleFunc("POST /post", r.PostHandler.Create)
	r.Mux.HandleFunc("PATCH /post/{id}", r.PostHandler.Update)
	r.Mux.HandleFunc("DELETE /post/{id}", r.PostHandler.Delete)
	r.Mux.HandleFunc("GET /post/{id}", r.PostHandler.FindByID)
	r.Mux.HandleFunc("GET /post", r.PostHandler.FindMany)
}
