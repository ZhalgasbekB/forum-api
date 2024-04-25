package server

import (
	"log"
	"net/http"

	"gitea.com/lzhuk/forum/internal/convert"
	"gitea.com/lzhuk/forum/internal/helpers/json"

	"gitea.com/lzhuk/forum/internal/errors"
)

// FOR CHECKING  ADMIN
// user := contextUser(r)
// if user.Role != roles.ADMIN {
// 	log.Println("Not Admin")
// 	errors.ErrorSend(w, http.StatusInternalServerError, "Because You Are Not Admin")
// 	return
// }
//

func (h *Handler) Admin(w http.ResponseWriter, r *http.Request) {
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

func (h *Handler) AdminChangeRole(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	uRole, err := convert.UpdateRole(r)
	if err != nil {
		log.Println(err)
		errors.ErrorSend(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := h.Services.Admin.UpdateUserService(uRole); err != nil {
		log.Println(err)
		errors.ErrorSend(w, http.StatusInternalServerError, err.Error())
		return
	}
	json.WriteJSON(w, http.StatusOK, uRole.Role)
}

func (h *Handler) AdminDeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	id, err := convert.DeleteUser(r)
	if err != nil {
		log.Println(err)
		errors.ErrorSend(w, http.StatusInternalServerError, err.Error())
		return
	}
	if err := h.Services.Admin.DeleteUserService(id); err != nil {
		log.Println(err)
		errors.ErrorSend(w, http.StatusInternalServerError, err.Error())
		return
	}
	json.WriteJSON(w, http.StatusOK, id)
}

func (h *Handler) AdminUpdateAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	user, err := convert.UpdateUserAdmin(r)
	if err != nil {
		log.Println(err)
		errors.ErrorSend(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := h.Services.Admin.UpdateUserNewDateService(user); err != nil {
		log.Println(err)
		errors.ErrorSend(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}
