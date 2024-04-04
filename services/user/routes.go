package user

import (
	"fmt"
	"fun_project/services/auth"
	"fun_project/types"
	"fun_project/utils"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	store types.UserStore
}

func NewHander(store types.UserStore) *Handler {
	return &Handler{
		store: store,
	}
}

func (h *Handler) RegisterRouter(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
	router.HandleFunc("/user/{login}", h.handleGetUser).Methods("GET")

}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) handleGetUser(w http.ResponseWriter, r *http.Request) {
	user, err := h.store.GetUserByLogin(mux.Vars(r)["login"])
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	utils.WriteJSON(w, user, http.StatusOK)
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {

	var payload types.RegisterUserPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	_, err := h.store.GetUserByEmail(payload.Email)
	if err == nil {
		utils.WriteError(w, fmt.Errorf("user with this email does not exists"), http.StatusBadRequest)
		return
	}

	hashedPassword, err := auth.HashPassword(payload.Password)
	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
	}

	err = h.store.CreateUser(types.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  hashedPassword,
	})

	if err != nil {
		utils.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, nil, http.StatusCreated)
}
