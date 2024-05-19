package controllers

import "net/http"

func RegisterController() {
	uc := newUserController()

	// Wrap userController's ServeHttp method with handlerfunc instead of
	/*
		http.Handle("/users", *c)
		http.Handle("/users/", *c)
	*/
	handler := http.HandlerFunc(uc.ServeHttp)

	http.Handle("/users", handler)
	http.Handle("/users/", handler)
}
