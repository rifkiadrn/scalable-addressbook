package addressbookapi

import (
	"encoding/json"
)

type Processor interface {
	getAddressListProcessor() ([]byte, error)
	createNewAddressProcessor(address address) error
	updateAddressProcessor(newAddress address) error
	deleteAddressProcessor(id int) error
}

type addressProcessor struct {
	repo Repository
}

func (s *addressProcessor) getAddressListProcessor() ([]byte, error) {
	addressList, err := s.repo.getAddressList()
	if err != nil {
		return nil, err
	}
	addressListJSON, err := json.Marshal(addressList)
	if err != nil {
		return nil, err
	}
	return addressListJSON, nil
}

func (s *addressProcessor) createNewAddressProcessor(address address) error {
	err := s.repo.createNewAddress(address.Name, address.PhoneNumber)
	if err != nil {
		return err
	}
	return nil
}

func (s *addressProcessor) updateAddressProcessor(newAddress address) error {
	err := s.repo.updateAddress(newAddress.ID, newAddress.Name, newAddress.PhoneNumber)
	if err != nil {
		return err
	}

	return nil
}

func (s *addressProcessor) deleteAddressProcessor(id int) error {
	err := s.repo.deleteAddress(id)
	if err != nil {
		return err
	}

	return nil
}
