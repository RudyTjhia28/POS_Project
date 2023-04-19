package models

import "net/http"

type ResponseBody struct {
	Success           bool        `json:"success"`
	HttpStatusCode    int         `json:"-"`
	HttpStatusMessage string      `json:"-"`
	Result            interface{} `json:"result,omitempty"`
	ErrorMessage      string      `json:"errorMessage,omitempty"`
}

func SuccessResponse(result interface{}, httpStatusCode int) ResponseBody {
	return ResponseBody{
		Success:           true,
		HttpStatusCode:    httpStatusCode,
		HttpStatusMessage: http.StatusText(httpStatusCode),
		Result:            result,
	}
}

func ErrorResponse(httpStatusCode int, errorMessage string) ResponseBody {
	return ResponseBody{
		Success:           false,
		HttpStatusCode:    httpStatusCode,
		HttpStatusMessage: http.StatusText(httpStatusCode),
		ErrorMessage:      errorMessage,
	}
}
