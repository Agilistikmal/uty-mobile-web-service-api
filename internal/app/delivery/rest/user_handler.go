package rest

import (
	"encoding/json"
	"net/http"

	"github.com/agilistikmal/uty-mobile-web-service-api/internal/app/model"
	"github.com/agilistikmal/uty-mobile-web-service-api/internal/app/service"
	"github.com/agilistikmal/uty-mobile-web-service-api/internal/pkg"
)

type UserHandler struct {
	service    *service.UserService
	otpService *service.OTPService
}

func NewUserHandler(service *service.UserService, otpService *service.OTPService) *UserHandler {
	return &UserHandler{
		service:    service,
		otpService: otpService,
	}
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var user *model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		pkg.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	user, err = h.service.Register(user)
	if err != nil {
		pkg.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	_, err = h.otpService.Generate(user.Username)
	if err != nil {
		pkg.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	pkg.SendSuccess(w, user)
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var user *model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		pkg.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	user, err = h.service.Login(user.Username, user.Password)
	if err != nil {
		pkg.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	pkg.SendSuccess(w, user)
}
