package controllers

import (
	"fmt"
	"net/http"
)

func Test(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/test" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusMethodNotAllowed)
		return
	}

	if _, err := fmt.Fprintf(w, "Hello"); err != nil {
		http.Error(w, "failed to write response", http.StatusInternalServerError)
		return
	}
}
