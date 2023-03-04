package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"template/internal/interfaces"
)

type UserControllerImpl struct {
	userService interfaces.UserService
}

func NewUserController(userService interfaces.UserService) interfaces.UserController {
	return &UserControllerImpl{
		userService: userService,
	}
}

func (uc *UserControllerImpl) CreateUser(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	var req struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Create user
	user, err := uc.userService.CreateUser(req.Name, req.Email, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (uc *UserControllerImpl) GetUser(w http.ResponseWriter, r *http.Request) {
	// Parse ID parameter
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Get user
	user, err := uc.userService.GetUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if user == nil {
		http.NotFound(w, r)
		return
	}

	// Return response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
