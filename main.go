package main

import (
	"net/http"

	"github.com/nyae44/webservice/controllers"
)

func main() {
	controllers.RegisterController()
	http.ListenAndServe(":3000", nil)
}
