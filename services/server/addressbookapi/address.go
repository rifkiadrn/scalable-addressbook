package addressbookapi

type address struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
}

type addressView struct {
	Status  int                  `json:"status"`
	Message string               `json:"message"`
	Result  map[string][]address `json:"result"`
}
