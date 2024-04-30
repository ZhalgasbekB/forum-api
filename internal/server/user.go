package server

import (
	"log"
	"net/http"

	"gitea.com/lzhuk/forum/internal/convert"
	"gitea.com/lzhuk/forum/internal/errors"
	"gitea.com/lzhuk/forum/internal/helpers/cookies"
	"gitea.com/lzhuk/forum/internal/helpers/json"
	"gitea.com/lzhuk/forum/internal/model"
)

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	userReq, err := convert.UserLoginRequestBody(r)
	if err != nil {
		log.Println(err)
		errors.ErrorSend(w, http.StatusInternalServerError, err.Error())
		return
	}

	user, err := h.Services.UserService.UserByEmailService(userReq.Email, userReq.Password)
	if err != nil {
		log.Println(err)
		errors.ErrorSend(w, http.StatusInternalServerError, err.Error())
		return
	}

	session, err := h.Services.SessionService.CreateSessionService(user.ID)
	if err != nil {
		log.Println(err)
		errors.ErrorSend(w, http.StatusInternalServerError, err.Error())
		return
	}

	cookies.CreateCookie(w, session)
	json.WriteJSON(w, http.StatusOK, model.UserResponseDTO{Name: user.Name, Email: user.Email})
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	user, err := convert.UserRegisterRequestBody(r)
	if err != nil {
		log.Println(err)
		errors.ErrorSend(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := h.Services.UserService.CreateUserService(user); err != nil {
		log.Println(err)
		errors.ErrorSend(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	cookie, err := cookies.Cookie(r)
	if err != nil {
		log.Println(err)
		errors.ErrorSend(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := h.Services.SessionService.DeleteSessionService(cookie.Value); err != nil {
		log.Println(err)
		errors.ErrorSend(w, http.StatusInternalServerError, err.Error())
		return
	}

	cookies.DeleteCookie(w)
	w.WriteHeader(http.StatusSeeOther)
}

func (h *Handler) Activity(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	user := contextUser(r)

	json.WriteJSON(w, http.StatusOK, user)
}
