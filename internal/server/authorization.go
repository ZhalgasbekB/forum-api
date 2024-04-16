package server

import (
	"gitea.com/lzhuk/forum/internal/convert"
	"gitea.com/lzhuk/forum/internal/errors"
	"gitea.com/lzhuk/forum/internal/helpers/cookies"
	json1 "gitea.com/lzhuk/forum/internal/helpers/json"
	"gitea.com/lzhuk/forum/internal/model"
	"log"
	"net/http"
)

func (h *Handler) Authenticate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		return
	}

	user, err := convert.AuthenticateUserDTO(r)
	if err != nil {
		log.Println(err)
		return
	}

	userByEmail, err := h.Services.UserService.UserByEmailService(user.Email, user.Password)
	if userByEmail == nil {
		if err != errors.ErrInvalidCredentials {
			log.Println(err)
			return
		}
		if err := h.Services.UserService.CreateUserService(user); err != nil {
			log.Println(err)
			errors.ErrorSend(w, http.StatusInternalServerError, err.Error())
			return
		}
		userByEmail.Name = user.Name
		userByEmail.Email = user.Email
	}

	session, err := h.Services.SessionService.CreateSessionService(userByEmail.ID)
	if err != nil {
		log.Println(err)
		return
	}

	cookies.CreateCookie(w, session)
	json1.WriteJSON(w, http.StatusOK, model.UserResponseDTO{Name: userByEmail.Name, Email: userByEmail.Email})
}
