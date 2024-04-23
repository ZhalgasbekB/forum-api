package server

import (
	"log"
	"net/http"

	"gitea.com/lzhuk/forum/internal/helpers/json"

	"gitea.com/lzhuk/forum/internal/errors"
)

func (h *Handler) AdminUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	users, err := h.Services.Admin.UsersService()
	if err != nil {
		log.Println(err)
		errors.ErrorSend(w, http.StatusInternalServerError, err.Error())
		return
	}
	json.WriteJSON(w, http.StatusOK, users)
}
