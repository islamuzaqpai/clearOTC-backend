package handler

import (
	"ClearOTC/internal/api/httpx"
	"ClearOTC/internal/models"
	"context"
	"encoding/json"
	"net/http"
)

type UserService interface {
	Register(ctx context.Context, req *models.CreateUserParams) (*models.UserRes, error)
}

type UserHandler struct {
	service UserService
}

func NewUserHandler(service UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req models.CreateUserParams

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpx.WriteError(w, http.StatusBadRequest, "Bad Request") //ю  знуть кастомные ошибки
		return
	}

	res, err := h.service.Register(r.Context(), &req)
	if err != nil {
		httpx.WriteError(w, http.StatusInternalServerError, "failed to register")
		return
	}

	httpx.WriteJSON(w, http.StatusCreated, res)
}
