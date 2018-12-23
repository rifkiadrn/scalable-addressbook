package main

import (
	"log"
	"net/http"

	"github.com/addressbook/services/server/appcontext"
	"github.com/addressbook/services/server/router"
)

func main() {
	appcontext.InitContext()
	router := router.CreateRouter()
	log.Fatal(http.ListenAndServe(":5000", router))
}
