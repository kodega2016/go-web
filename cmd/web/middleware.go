package main

import (
	"fmt"
	"net/http"
)

// WriteToConsole simply write a console message when a request is performed
func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("processing...", r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

// SessionLoad loads the session for the next request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
