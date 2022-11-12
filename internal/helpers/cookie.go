package helpers

import (
	"net/http"
)

func CreateCookie(name, value string, secure, httpOnly bool) *http.Cookie {
	return &http.Cookie{
		Name:  name,
		Value: value,
		Path:  "/",
		Secure: secure,
		HttpOnly: httpOnly,
	}
}