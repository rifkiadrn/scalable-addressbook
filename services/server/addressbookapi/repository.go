package addressbookapi

import (
	"net/http"

	"github.com/addressbook/services/server/helper"
)

type Repository interface {
	getAddressList() ([]address, error)
	createNewAddress(name, phoneNumber string) error
	updateAddress(id int, newName string, newPhoneNumber string) error
	deleteAddress(id int) error
}

type addressRepository struct {
	execer helper.QueryExecuter
}

const (
	getAddressListQuery   = "SELECT * FROM addresslist"
	createNewAddressQuery = "INSERT INTO addresslist(name, phone_number) VALUES($1,$2)"
	updateAddressQuery    = "UPDATE addresslist SET name=$2 and phone_number=$3 WHERE id=$1"
	deleteAddressQuery    = "DELETE FROM addresslist WHERE id=$1"
)

func initRepository(execer helper.QueryExecuter) *addressRepository {
	return &addressRepository{
		execer: execer,
	}
}

func (s *addressRepository) getAddressList() ([]address, error) {
	var addressList []address
	data, err := s.execer.DoQueryRow(getAddressListQuery)
	if err != nil {
		return nil, err
	}
	err = helper.ParseMap(data, &addressList)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, helper.NewHttpError(http.StatusNotFound, "Data not found")
	}
	return addressList, nil
}

func (s *addressRepository) createNewAddress(name, phoneNumber string) error {
	_, err := s.execer.DoQuery(createNewAddressQuery, name, phoneNumber)
	if err != nil {
		return err
	}
	return nil
}

func (s *addressRepository) updateAddress(id int, newName, newPhoneNumber string) error {

	_, err := s.execer.DoQuery(updateAddressQuery, id, newName, newPhoneNumber)
	if err != nil {
		return err
	}
	return nil
}

func (s *addressRepository) deleteAddress(id int) error {
	_, err := s.execer.DoQuerySlice(deleteAddressQuery, id)
	if err != nil {
		return err
	}
	return nil
}
