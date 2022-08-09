package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// func SetMimeType(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

//         rw.Header().Add("Content-Type","")
// 	})
// }

func Nosurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction, //In production, you have to switch to yes
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler

}
