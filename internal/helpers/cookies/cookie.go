package cookies

import (
	"net/http"
	"time"

	"gitea.com/lzhuk/forum/internal/model"
)

const (
	cookieName = "CookieUUID"
)

func Cookie(r *http.Request) (*http.Cookie, error) {
	cookie, err := r.Cookie(cookieName)
	if err != nil {
		return nil, err
	}
	return cookie, nil
}

func CreateCookie(w http.ResponseWriter, session *model.Sessinon) http.Cookie {
	cookie := http.Cookie{
		Name:    cookieName,
		Value:   session.UUID,
		Path:    "/", // NEED TO CHANGE ???
		Expires: session.ExpireAt,
		MaxAge:  int(time.Until(session.ExpireAt).Seconds()),
	}
	// http.SetCookie(w, &cookie)

	return cookie
}

func DeleteCookie(w http.ResponseWriter) {
	cookie := http.Cookie{
		Name:   cookieName,
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, &cookie)
}
