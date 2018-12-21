package router

import (
	"github.com/addressbook/services/server/addressbookapi"
	"github.com/gorilla/mux"
)

func CreateRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/address/list", addressbookapi.GetAddressListHandler).Methods("GET")
	router.HandleFunc("/address/create", addressbookapi.CreateNewAddressHandler).Methods("POST")
	router.HandleFunc("/address/update", addressbookapi.UpdateAddressHandler).Methods("POST")
	router.HandleFunc("/address/delete", addressbookapi.DeleteAddressHandler).Methods("DELETE")

	return router
}
