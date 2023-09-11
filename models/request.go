package models

import (
	"encoding/json"
	"fmt"
)

type Response struct {
	Body []byte `json:"response_body"`
	Code int    `json:"response_code"`
}

func (r *Response) Error() error {
	if r.Code/100 == 2 {
		return nil
	}

	if r.Body == nil {
		return fmt.Errorf("error with the response body. code: %v", r.Code)
	}

	var responseMessage ErrorResponse
	err := json.Unmarshal(r.Body, &responseMessage)
	if err != nil {
		return fmt.Errorf("error while unmarshaling error response body. http code: %v, unmarshal error: %v", r.Code, err)
	}

	return fmt.Errorf("error occurred. http code: %v, error msg: %v", r.Code, responseMessage.Error)
}
