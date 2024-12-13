package rest

import (
	"encoding/json"
	"net/http"

	"github.com/agilistikmal/uty-mobile-web-service-api/internal/app/model"
	"github.com/agilistikmal/uty-mobile-web-service-api/internal/app/service"
	"github.com/agilistikmal/uty-mobile-web-service-api/internal/pkg"
)

type PostHandler struct {
	service *service.PostService
}

func NewPostHandler(service *service.PostService) *PostHandler {
	return &PostHandler{
		service: service,
	}
}

func (h *PostHandler) Create(w http.ResponseWriter, r *http.Request) {
	var request *model.PostCreateRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		pkg.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	post, err := h.service.Create(request)
	if err != nil {
		pkg.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	pkg.SendSuccess(w, post)
}

func (h *PostHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	var request *model.PostUpdateRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		pkg.SendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	post, err := h.service.Update(id, request)
	if err != nil {
		pkg.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	pkg.SendSuccess(w, post)
}

func (h *PostHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	post, err := h.service.Delete(id)
	if err != nil {
		pkg.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	pkg.SendSuccess(w, post)
}

func (h *PostHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	post, err := h.service.FindByID(id)
	if err != nil {
		pkg.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	pkg.SendSuccess(w, post)
}

func (h *PostHandler) FindMany(w http.ResponseWriter, r *http.Request) {
	posts := h.service.FindMany()

	pkg.SendSuccess(w, posts)
}
