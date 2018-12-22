package addressbookapi

import (
	"fmt"
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
	createNewAddressQuery = "INSERT INTO addresslist(name, phone_number) VALUES($1,$2) RETURNING *"
	updateAddressQuery    = "UPDATE addresslist SET name=$2, phone_number=$3 WHERE id=$1 RETURNING *"
	deleteAddressQuery    = "DELETE FROM addresslist WHERE id=$1 RETURNING *"
)

func initRepository(execer helper.QueryExecuter) *addressRepository {
	return &addressRepository{
		execer: execer,
	}
}

func (s *addressRepository) getAddressList() ([]address, error) {
	var addressList []address
	data, err := s.execer.DoQuery(getAddressListQuery)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return []address{}, nil
	}
	err = helper.ParseMap(data, &addressList)
	if err != nil {
		return nil, err
	}
	return addressList, nil
}

func (s *addressRepository) createNewAddress(name, phoneNumber string) error {
	data, err := s.execer.DoQueryRow(createNewAddressQuery, name, phoneNumber)
	if err != nil {
		return err
	}
	if data == nil {
		return helper.NewHttpError(http.StatusBadRequest, "Failed to insert address")
	}
	return nil
}

func (s *addressRepository) updateAddress(id int, newName, newPhoneNumber string) error {

	data, err := s.execer.DoQueryRow(updateAddressQuery, id, newName, newPhoneNumber)
	if err != nil {
		return err
	}
	if data == nil {
		return helper.NewHttpError(http.StatusNoContent, "Data is not exist")
	}
	fmt.Println(data)
	return nil
}

func (s *addressRepository) deleteAddress(id int) error {
	data, err := s.execer.DoQueryRow(deleteAddressQuery, id)
	if err != nil {
		return err
	}
	if data == nil {
		return helper.NewHttpError(http.StatusNoContent, "Data is not exist")
	}
	return nil
}
