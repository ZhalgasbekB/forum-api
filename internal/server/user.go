package server

import (
	"fmt"
	"net/http"

	"gitea.com/lzhuk/forum/internal/convert"
	"gitea.com/lzhuk/forum/internal/errors"
	"gitea.com/lzhuk/forum/internal/helpers/cookies"
)

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	userReq, err := convert.UserLoginRequestBody(r)
	user, err := h.Services.UserService.UserByEmailService(userReq.Email, userReq.Password)
	// ???
	if err != nil {
		if err == errors.ErrSQLNoRows {
			errors.NewError(http.StatusInternalServerError, err.Error())
			return
		}
		return
	}
	// ???

	session, err := h.Services.SessionService.CreateSessionService(user.ID)
	if err != nil {
		fmt.Println(err)
		return
	}
	cookies.CreateCookie(w, session)
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	user, err := convert.UserRegisterRequestBody(r)
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := h.Services.UserService.CreateUserService(user); err != nil {
		fmt.Println(err)
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
		fmt.Println(err)
		return
	}

	if err := h.Services.SessionService.DeleteSessionService(cookie.Value); err != nil {
		return
	}
	cookies.DeleteCookie(w)
	w.WriteHeader(http.StatusSeeOther)
}
