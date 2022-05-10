package helper

import "strings"

// Response is used for static shape json return
type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

// Empty object is used when data doesnt want to be nill on json
type EmptyObj struct{}

func BuildResponse(status bool, message string, data interface{}) Response {
	res := Response{
		Status:  status,
		Message: message,
		Errors:  nil,
		Data:    data,
	}

	return res
}

func BuildErrorResponse(message string, errors string, data interface{}) Response {
	splitedError := strings.Split(errors, "\n")
	res := Response{
		Status:  false,
		Message: message,
		Errors:  splitedError,
		Data:    nil,
	}

	return res
}
