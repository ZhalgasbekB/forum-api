package server

import (
	"context"
	"log"
	"net/http"
	"time"

	"gitea.com/lzhuk/forum/internal/errors"
	"gitea.com/lzhuk/forum/internal/helpers/cookies"
	"gitea.com/lzhuk/forum/internal/helpers/response"
)

func (h *Handler) IsAuthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := cookies.Cookie(r)
		if err != nil {
			log.Println(err)
			next.ServeHTTP(w, r)
			return
		}
		uuid := cookie.Value
		session, err := h.Services.SessionService.GetSessionByUUIDService(uuid)
		if err != nil {
			log.Println(err)
			cookies.DeleteCookie(w)
			next.ServeHTTP(w, r)
			return
		}

		if session.ExpireAt.Before(time.Now()) {
			log.Println("Expired Session:")
			cookies.DeleteCookie(w)
			next.ServeHTTP(w, r)
			return
		}

		user, err := h.Services.UserService.UserByIDService(session.UserID)
		if err != nil {
			log.Println(err)
			return
		}

		ctx := context.WithValue(r.Context(), key, user)

		ctxWithTimeout, cancel := context.WithTimeout(ctx, time.Second*5)
		defer cancel()

		next.ServeHTTP(w, r.WithContext(ctxWithTimeout))
	})
}

func (h *Handler) RequiredAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := contextUser(r)
		if user == nil {
			response.WriteJSON(w, http.StatusSeeOther, errors.NewError(http.StatusSeeOther, "No Authenticated User: Please Authenticate"))
			return
		}
		next.ServeHTTP(w, r)
	})
}
