package rest

import (
	"encoding/json"
	"net/http"

	"github.com/agilistikmal/uty-mobile-web-service-api/internal/app/model"
	"github.com/agilistikmal/uty-mobile-web-service-api/internal/app/service"
	"github.com/agilistikmal/uty-mobile-web-service-api/internal/pkg"
)

type OTPHandler struct {
	service     *service.OTPService
	UserService *service.UserService
}

func NewOTPHandler(service *service.OTPService, userService *service.UserService) *OTPHandler {
	return &OTPHandler{
		service:     service,
		UserService: userService,
	}
}

func (h *OTPHandler) Verify(w http.ResponseWriter, r *http.Request) {
	var otp *model.OTP
	err := json.NewDecoder(r.Body).Decode(&otp)
	if err != nil {
		pkg.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.service.Verify(otp.Username, otp.Code)
	if err != nil {
		pkg.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	userResponse, err := h.UserService.Find(otp.Username)
	if err != nil {
		pkg.SendError(w, http.StatusNotFound, err.Error())
		return
	}

	pkg.SendSuccess(w, userResponse)
}
