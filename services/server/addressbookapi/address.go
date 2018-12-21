package addressbookapi

type address struct {
	ID          int    `json:id`
	Name        string `json:name`
	PhoneNumber string `json:phone_number`
}
