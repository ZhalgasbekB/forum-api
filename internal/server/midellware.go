package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"gitea.com/lzhuk/forum/internal/helpers/cookies"
)

func (h *Handler) IsAuthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := cookies.Cookie(r)
		if err != nil {
			fmt.Println(err)
			return
		}
		
		uuid := cookie.Value
		session, err := h.Services.SessionService.GetSessionByUUIDService(uuid)
		if err != nil {
			fmt.Println(err)
			return
		}
		
		if session.ExpireAt.Before(time.Now()) {
			cookies.DeleteCookie(w)
			next.ServeHTTP(w, r)
			return
		}
		user, err := h.Services.UserService.UserByIDService(session.UserID)
		
		if err != nil {
			fmt.Println(err)
			return
		}

		ctx := context.WithValue(r.Context(), "key", user)
		next.ServeHTTP(w, r.WithContext(ctx))
	
	})
}

func (h *Handler) RequiredAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	})
}
