package route

import (
	"net/http"

	"github.com/agilistikmal/uty-mobile-web-service-api/internal/app/delivery/rest"
)

type Route struct {
	Mux *http.ServeMux

	UserHandler *rest.UserHandler
}

func NewRoutes(userHandler *rest.UserHandler) *Route {
	return &Route{
		Mux:         http.NewServeMux(),
		UserHandler: userHandler,
	}
}

func (r *Route) Init() {
	r.ProductRoutes()
}

func (r *Route) ProductRoutes() {
	r.Mux.HandleFunc("POST /auth/register", r.UserHandler.Register)
	r.Mux.HandleFunc("POST /auth/login", r.UserHandler.Login)
}
