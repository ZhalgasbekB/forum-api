package server

import (
	"context"
	"log"
	"net/http"
	"time"

	"gitea.com/lzhuk/forum/internal/errors"
	"gitea.com/lzhuk/forum/internal/helpers/cookies"
	"gitea.com/lzhuk/forum/internal/helpers/roles"
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
			log.Println("Time Expired")
			cookies.DeleteCookie(w)
			next.ServeHTTP(w, r)
			return
		}

		user, err := h.Services.UserService.UserByIDService(session.UserID)
		if err != nil {
			log.Println(err)
			errors.ErrorSend(w, http.StatusInternalServerError, err.Error())
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
			log.Println("No Authenticated User: Please Authenticate")
			errors.ErrorSend(w, http.StatusSeeOther, "No Authenticated User: Please Authenticate")
			return
		}
		next.ServeHTTP(w, r)
	})
}

func RateLimitMiddleware(limit int, interval time.Duration) func(http.Handler) http.Handler {
	tokens := make(chan struct{}, limit)

	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				select {
				case tokens <- struct{}{}:
				default:
				}
			}
		}
	}()

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			select {
			case <-tokens:
				next.ServeHTTP(w, r)
			default:
				errors.ErrorSend(w, http.StatusTooManyRequests, "Too many requests")
				return
			}
		})
	}
}

/// ADMIN and MODERATOR
func (h *Handler) AdminVerification(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := contextUser(r)
		if user == nil {
			log.Println("No Authenticated User: Please Authenticate")
			errors.ErrorSend(w, http.StatusBadRequest, "No Authenticated User: Please Authenticate")
			return
		}
		if user.Role != roles.ADMIN {
			log.Println("No Authenticated Admin: You aren't admin")
			errors.ErrorSend(w, http.StatusBadRequest, "No Authenticated Admin: You Aren't Admin")
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (h *Handler) ModeratorVerification(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := contextUser(r)
		if user == nil {
			log.Println("No Authenticated User: Please Authenticate")
			errors.ErrorSend(w, http.StatusBadRequest, "No Authenticated User: Please Authenticate")
			return
		}
		if user.Role != roles.MODERATOR {
			log.Println("No Authenticated Moderator: You aren't moderator")
			errors.ErrorSend(w, http.StatusBadRequest, "No Authenticated Moderator: You Aren't Moderator")
			return

		}
		next.ServeHTTP(w, r)
	})
}
