package controllers

import (
	"net/http"
	"regexp"
)

type userController struct {
	userIDPattern *regexp.Regexp
}

func (uc userController) ServeHttp(w http.ResponseWriter, r http.Request) {
	w.Write([]byte("Hello from the user controller!"))
}

// The regex looks for paths i.e /users/1
func newUserController() *userController {
	return &userController{
		userIDPattern: regexp.MustCompile(`^/users(\d+)/?`),
	}
}
