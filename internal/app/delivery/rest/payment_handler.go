package rest

import (
	"encoding/json"
	"net/http"

	"github.com/agilistikmal/uty-mobile-web-service-api/internal/app/model"
	"github.com/agilistikmal/uty-mobile-web-service-api/internal/app/service"
	"github.com/agilistikmal/uty-mobile-web-service-api/internal/pkg"
)

type PaymentHandler struct {
	service *service.PaymentService
}

func NewPaymentHandler(service *service.PaymentService) *PaymentHandler {
	return &PaymentHandler{
		service: service,
	}
}

func (h *PaymentHandler) Create(w http.ResponseWriter, r *http.Request) {
	var payment *model.Payment
	err := json.NewDecoder(r.Body).Decode(&payment)
	if err != nil {
		pkg.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	payment, err = h.service.Create(payment)
	if err != nil {
		pkg.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	pkg.SendSuccess(w, payment)
}

func (h *PaymentHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	payment, err := h.service.FindByID(id)
	if err != nil {
		pkg.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	pkg.SendSuccess(w, payment)
}

func (h *PaymentHandler) FindByReferenceID(w http.ResponseWriter, r *http.Request) {
	referenceID := r.PathValue("reference_id")
	payment, err := h.service.FindByReferenceID(referenceID)
	if err != nil {
		pkg.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	pkg.SendSuccess(w, payment)
}
