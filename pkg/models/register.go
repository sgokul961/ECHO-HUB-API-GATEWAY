package models

type Response struct {
	Status  int
	Message string
	Data    interface{}
	Error   interface{}
}

func MakeResponse(status int, message string, data interface{}, error interface{}) *Response {
	return &Response{
		Status:  status,
		Message: message,
		Data:    data,
		Error:   error,
	}
}
