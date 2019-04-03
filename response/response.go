package response

type ResponseEntity struct {
	Code ResponseCode `json:"code"`
	Data interface{}  `json:"data"`
	Msg  string       `json:"message"`
}

type ResponseCode string

const (
	SUCCESS ResponseCode = "SUCCESS"

	ERROR ResponseCode = "Server error"

	PARAMS_VALIDATE_ERROR ResponseCode = "system.validate.error"

	NOT_FOUND_ERROR ResponseCode = "system.error.not_found"

	IMPROPER_OPERATION_ERROR ResponseCode = "This operation is not appropriate."

	UNAUTHORIZED_ERROR ResponseCode = "You are not authorized to access this page."
)

func Success(data interface{}) *ResponseEntity {
	return &ResponseEntity{Code: SUCCESS, Data: data, Msg: "Success"}
}

func Error(data interface{}, msg string) *ResponseEntity {
	return &ResponseEntity{Code: ERROR, Data: data, Msg: msg}
}

func ErrorParamsValidate(data interface{}, msg string) *ResponseEntity {
	return &ResponseEntity{Code: PARAMS_VALIDATE_ERROR, Data: data, Msg: msg}
}

func ErrorNotFoundError(data interface{}, msg string) *ResponseEntity {
	return &ResponseEntity{Code: NOT_FOUND_ERROR, Data: data, Msg: msg}
}

func ErrorImproperOperationError(data interface{}, msg string) *ResponseEntity {
	return &ResponseEntity{Code: IMPROPER_OPERATION_ERROR, Data: data, Msg: msg}
}

func ErrorUnauthorizedErrore(data interface{}, msg string) *ResponseEntity {
	return &ResponseEntity{Code: UNAUTHORIZED_ERROR, Data: data, Msg: msg}
}
