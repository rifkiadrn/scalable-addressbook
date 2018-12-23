package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/addressbook/services/server/appcontext"
	"github.com/addressbook/services/server/router"
)

func main() {
	appcontext.InitContext()
	router := router.CreateRouter()
	log.Fatal(http.ListenAndServe(":5000", router))
}
