package customerror

type CustomError struct {
	Code  string      `json:"code"`
	Data  interface{} `json:"data"`
	Msg   string      `json:"message"`
	Extra interface{} `json:"extra"`
}

func New(code, msg string) *CustomError {
	return &CustomError{Code: code, Msg: msg}
}
