package main

import (
	"fmt"

	"github.com/nyae44/webservice/models"
)

func main() {
	u := models.User{
		ID:        2,
		FirstName: "Daley",
		LastName:  "Nyae",
	}
	fmt.Println(u)
}
