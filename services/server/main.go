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
	fmt.Println(os.Getenv("POSTGRES_HOST") + os.Getenv("POSTGRES_PORT") + os.Getenv("POSTGRES_USER") +
		os.Getenv("POSTGRES_PASSWORD") + os.Getenv("POSTGRES_DB"))
	appcontext.InitContext()
	router := router.CreateRouter()
	log.Fatal(http.ListenAndServe(":5000", router))
}
