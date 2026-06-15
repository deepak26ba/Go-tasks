package handlers

import (
	"encoding/json"
	"http/models"
	"http/service"
	"net/http"
	"strconv"
)

var userRoleRow models.UserRoles

type Handler struct {
	service service.Service
}

func New(service service.Service) *Handler {

	return &Handler{service: service}

}

func (h *Handler) CreateUserRole(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&userRoleRow); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := h.service.CreateUserRoles(userRoleRow)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		err := models.Error{
			Message:    err.Error(),
			Error:      err,
			StatusCode: http.StatusInternalServerError,
		}
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)

}

func (h *Handler) GetUserRole(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var rows []models.UserRoles

	result, err := h.service.GetUserRoles(rows)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		err := models.Error{
			Message:    err.Error(),
			Error:      err,
			StatusCode: http.StatusInternalServerError,
		}
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)

}

func (h *Handler) GetByIdUserRole(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "ID is missing from query parameters", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	result, err := h.service.GetByIdUserRoles(userRoleRow, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		err := models.Error{
			Message:    err.Error(),
			Error:      err,
			StatusCode: http.StatusInternalServerError,
		}
		json.NewEncoder(w).Encode(err)
		return
	}

	json.NewEncoder(w).Encode(result)

}

func (h *Handler) DeleteByIdUserRole(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodDelete {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "ID is missing from query parameters", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	result, err := h.service.GetByIdUserRoles(userRoleRow, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		err := models.Error{
			Message:    err.Error(),
			Error:      err,
			StatusCode: http.StatusInternalServerError,
		}
		json.NewEncoder(w).Encode(err)
		return
	}

	json.NewEncoder(w).Encode(result)

}

func (h *Handler) UpdateUserRole(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPut {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "ID is missing from query parameters", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&userRoleRow); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := h.service.GetByIdUserRoles(userRoleRow, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		err := models.Error{
			Message:    err.Error(),
			Error:      err,
			StatusCode: http.StatusInternalServerError,
		}
		json.NewEncoder(w).Encode(err)
		return
	}

	json.NewEncoder(w).Encode(result)

}

func (h *Handler) UpdateByIdUserRole(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPatch {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "ID is missing from query parameters", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&userRoleRow); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := h.service.GetByIdUserRoles(userRoleRow, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		err := models.Error{
			Message:    err.Error(),
			Error:      err,
			StatusCode: http.StatusInternalServerError,
		}
		json.NewEncoder(w).Encode(err)
		return
	}

	json.NewEncoder(w).Encode(result)

}
