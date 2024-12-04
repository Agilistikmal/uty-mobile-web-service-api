package rest

import (
	"encoding/json"
	"fmt"
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
	// Melakukan konversi request body JSON ke struct model user
	var user *model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		pkg.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Melakukan registrasi user
	userResponse, err := h.service.Register(user)
	if err != nil {
		pkg.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	_, err = h.otpService.Generate(userResponse)
	if err != nil {
		pkg.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	pkg.SendSuccess(w, fmt.Sprintf("OTP code has been sent to your WhatsApp (%s)", userResponse.Phone))
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	// Melakukan konversi request body JSON ke struct model user
	var user *model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		pkg.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Melakukan login
	userResponse, err := h.service.Login(user.Username, user.Password)
	if err != nil {
		pkg.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	_, err = h.otpService.Generate(userResponse)
	if err != nil {
		pkg.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	pkg.SendSuccess(w, fmt.Sprintf("OTP code has been sent to your WhatsApp (%s)", userResponse.Phone))
}

func (h *UserHandler) Find(w http.ResponseWriter, r *http.Request) {
	username := r.PathValue("username")

	// Mencari user
	user, err := h.service.Find(username)
	if err != nil {
		pkg.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	pkg.SendSuccess(w, user)
}

func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	username := r.PathValue("username")

	var request *model.UserUpdateRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		pkg.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Mengupdate user
	user, err := h.service.Update(username, request)
	if err != nil {
		pkg.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	pkg.SendSuccess(w, user)
}

func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	username := r.PathValue("username")

	_, err := h.service.Delete(username)
	if err != nil {
		pkg.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	pkg.SendSuccess(w, nil)
}
