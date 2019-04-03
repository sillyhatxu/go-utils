package errors

import (
	"encoding/json"
	"fmt"
	"gopintar-utils/response"
)

type ShopintarError struct {
	Code response.ResponseCode `json:"code"`
	Data interface{}           `json:"data"`
	Msg  string                `json:"message"`
}

func (se ShopintarError) Error() string {
	seJSON, err := json.Marshal(se)
	if err != nil {
		return fmt.Sprintf(`{ "code" : "%s", "data" : "%s" , "message" : %s }`, se.Code, se.Data, se.Msg)
	}
	return string(seJSON)
}

func NewError(msg string) error {
	return ShopintarError{Code: response.ERROR, Data: nil, Msg: msg}
}

func NewCodeError(code response.ResponseCode, msg string) error {
	return ShopintarError{Code: code, Data: nil, Msg: msg}
}
