package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

func RegisterController() {
	uc := newUserController()

	// http.Handle("/users", *uc)
	// http.Handle("/users/", *uc)

	handler := http.HandlerFunc(uc.ServeHttp)

	http.Handle("/users", handler)
	http.Handle("/users/", handler)
}

func encodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
