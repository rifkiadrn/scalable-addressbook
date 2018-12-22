package helper

import (
	"fmt"
)

type SuccessResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Result  string `json:"result"`
}

func NewSuccessResponse(status int, message string, result []byte) *SuccessResponse {
	return &SuccessResponse{
		Message: message,
		Status:  status,
		Result:  string(result),
	}
}

func (e SuccessResponse) Response() string {
	return fmt.Sprintf(`{"status:": "%v",  "message": "%v", "result": %v}`, e.Status, e.Message, string(e.Result))
}
