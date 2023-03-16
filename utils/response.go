package utils

type Response struct {
	Status     bool   `json:"status"`
	Message    string `json:"message"`
	HTTPStatus uint   `json:"http status"`
	Data       any    `json:"data"`
}

type ResponseError struct {
	Status     bool  `json:"status"`
	Message    error `json:"message"`
	HTTPStatus uint  `json:"http status"`
	Data       any   `json:"data"`
}

func SuccessResponse(message string, status uint, data any) Response {
	res := Response{
		Status:     true,
		Message:    message,
		HTTPStatus: status,
		Data:       data,
	}
	return res
}

func ErrorResponse(message string, status uint) Response {
	res := Response{
		Status:     false,
		Message:    message,
		HTTPStatus: status,
		Data:       nil,
	}
	return res
}

func ErrorResponseInput(message error, status uint) ResponseError {
	res := ResponseError{
		Status:     false,
		Message:    message,
		HTTPStatus: status,
		Data:       nil,
	}
	return res
}
