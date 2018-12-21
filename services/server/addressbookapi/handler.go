package addressbookapi

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/addressbook/services/server/appcontext"
	"github.com/addressbook/services/server/helper"
)

var proc Processor

func getProcessor() Processor {
	if proc == nil {
		proc = &addressProcessor{repo: initRepository(appcontext.GetHelper())}
	}
	return proc
}

func GetAddressListHandler(w http.ResponseWriter, r *http.Request) {
	addressListJSON, err := getProcessor().getAddressListProcessor()
	if err != nil {
		helper.ResponseError("GetAddressList", w, helper.NewHttpError(http.StatusInternalServerError, err.Error()))
		return
	}
	helper.ResponseHelper(w, http.StatusOK, helper.ContentJson, string(addressListJSON))
}
func CreateNewAddressHandler(w http.ResponseWriter, r *http.Request) {
	var addressData address
	err := json.NewDecoder(r.Body).Decode(&addressData)
	if err != nil {
		helper.ResponseError("CreateNewAddress", w, helper.NewHttpError(http.StatusBadRequest, err.Error()))
		return
	}
	err = getProcessor().createNewAddressProcessor(addressData)
	if err != nil {
		helper.ResponseError("CreateNewAddress", w, err)
		return
	}
	helper.ResponseHelper(w, http.StatusCreated, helper.ContentText, "201 Created")
}
func UpdateAddressHandler(w http.ResponseWriter, r *http.Request) {
	var addressData address
	err := json.NewDecoder(r.Body).Decode(&addressData)
	if err != nil {
		helper.ResponseError("UpdateAddress", w, helper.NewHttpError(http.StatusBadRequest, err.Error()))
		return
	}
	err = getProcessor().updateAddressProcessor(addressData)
	if err != nil {
		helper.ResponseError("UpdateAddress", w, err)
		return
	}
	helper.ResponseHelper(w, http.StatusCreated, helper.ContentText, "201 Created")
}
func DeleteAddressHandler(w http.ResponseWriter, r *http.Request) {
	var address address
	err := json.NewDecoder(r.Body).Decode(&address)
	if err != nil {
		helper.ResponseError("DeleteAddress", w, helper.NewHttpError(http.StatusBadRequest, err.Error()))
		return
	}
	log.Printf("DeleteUser: Get Request to delete address of id: %d", address.ID)
	defer r.Body.Close()

	err = getProcessor().deleteAddressProcessor(address.ID)
	if err != nil {
		helper.ResponseError("DeleteAddress", w, err)
		return
	}
	helper.ResponseHelper(w, http.StatusOK, helper.ContentText, "200 OK")
}
