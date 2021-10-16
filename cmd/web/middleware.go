package main

import (
	"github.com/justinas/nosurf"
	"net/http"
)


// NoSurf adds CSRF protection to all POST requests
func NoSurf(next http.Handler) http.Handler{
	csrfHamdler := nosurf.New(next)

	csrfHamdler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path: "/",
		Secure: app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHamdler

}
// SessionLoad loads and saves the session on every request
func SessionLoad(next http.Handler) http.Handler{
	return session.LoadAndSave(next)
}

