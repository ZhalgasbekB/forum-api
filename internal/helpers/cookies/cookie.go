package cookies

import (
	"net/http"
	"time"

	"gitea.com/lzhuk/forum/internal/model"
)

const (
	cookieName = "CookieUUID"
)

func CreateCookie(session *model.Session) http.Cookie {
	cookie := http.Cookie{
		Name:    cookieName,
		Value:   session.UUID,
		Path:    "/",
		Expires: session.ExpireAt,
		MaxAge:  int(time.Until(session.ExpireAt).Seconds()),
	}
	return cookie
}

func DeleteCookie() http.Cookie {
	cookie := http.Cookie{
		Name:   cookieName,
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	return cookie
}
