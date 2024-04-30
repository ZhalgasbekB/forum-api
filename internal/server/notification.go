package server

import (
	"log"
	"net/http"

	"gitea.com/lzhuk/forum/internal/errors"
	"gitea.com/lzhuk/forum/internal/helpers/json"
)

func (h *Handler) Notifications(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	user := contextUser(r)
	notifications, err := h.Services.Nothification.ReadService(user.ID)
	if err != nil {
		log.Println(err)
		errors.ErrorSend(w, http.StatusInternalServerError, err.Error())
		return
	}

	json.WriteJSON(w, http.StatusOK, notifications)
}

func (h *Handler) Notification(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	user := contextUser(r)
	if err := h.Services.Nothification.UpdateService(user.ID); err != nil {
		log.Println(err)
		errors.ErrorSend(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}
