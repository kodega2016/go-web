package handlers

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "home page")
}

func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "about page")
}
