package server

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"gitea.com/lzhuk/forum/internal/helpers/cookies"
	"gitea.com/lzhuk/forum/internal/helpers/response"
)

func (h *Handler) IsAuthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var cookie http.Cookie
		if err := json.NewDecoder(r.Body).Decode(cookie); err != nil {
			log.Println(err)
			next.ServeHTTP(w, r)
			return
		}

		uuid := cookie.Value
		session, err := h.Services.SessionService.GetSessionByUUIDService(uuid)
		if err != nil {
			log.Println(err)
			return
		}

		if session.ExpireAt.Before(time.Now()) {
			response.WriteJSON(w, http.StatusSeeOther, cookies.DeleteCookie()) // ??? can be work
			// next.ServeHTTP(w, r) // Come on the client like a expired Cookie
			return
		}

		user, err := h.Services.UserService.UserByIDService(session.UserID)
		if err != nil {
			log.Println(err)
			return
		}

		ctx := context.WithValue(r.Context(), key, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (h *Handler) RequiredAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := contextUser(r)
		if user != nil {
			response.WriteJSON(w, 401, user)
			return
		}
		next.ServeHTTP(w, r)
	})
}
