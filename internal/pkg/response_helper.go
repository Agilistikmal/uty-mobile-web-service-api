package pkg

import (
	"encoding/json"
	"net/http"

	"github.com/agilistikmal/uty-mobile-web-service-api/internal/app/model"
)

func SendSuccess(w http.ResponseWriter, data any) {
	resp := &model.Response{
		Success: true,
		Code:    http.StatusOK,
		Message: "ok",
		Data:    data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func SendError(w http.ResponseWriter, code int, message string) {
	resp := &model.Response{
		Success: false,
		Code:    code,
		Message: message,
		Data:    nil,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(resp)
}
