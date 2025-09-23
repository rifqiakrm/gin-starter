// Package response provides basic API response
package response

import (
	"net/http"
	"os"
)

// APIResponseList defines a generic interface for API responses.
// It provides methods to retrieve response code, message, and data.
type APIResponseList interface {
	GetCode() int
	GetMessage() string
	GetData() interface{}
}

// apiResponseList is the concrete implementation of APIResponseList.
// It encapsulates the HTTP status code, response message, and optional data payload.
type apiResponseList struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// GetCode returns the HTTP status code of the response.
func (a apiResponseList) GetCode() int {
	return a.Code
}

// GetMessage returns the response message.
func (a apiResponseList) GetMessage() string {
	return a.Message
}

// GetData returns the response data payload.
// It can be nil if no data is associated with the response.
func (a apiResponseList) GetData() interface{} {
	return a.Data
}

// SuccessAPIResponse creates a standardized success response.
// Parameters:
//   - code: HTTP status code (e.g., 200).
//   - message: A success message.
//   - data: Optional data payload to include in the response.
//
// Returns an APIResponseList that can be serialized into JSON.
func SuccessAPIResponse(code int, message string, data interface{}) APIResponseList {
	if data == nil {
		data = map[string]interface{}{}
	}

	return &apiResponseList{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

// ErrorAPIResponse creates a standardized error response.
// In production mode (APP_ENV=production), certain messages are sanitized:
//   - 400 -> "bad request"
//   - 500 -> "something wrong on the server. contact server admin."
//
// Parameters:
//   - code: HTTP status code (e.g., 400, 500).
//   - message: Error message (maybe overridden in production).
//
// Returns an APIResponseList without a data payload.
func ErrorAPIResponse(code int, message string) APIResponseList {
	if os.Getenv("APP_ENV") == "production" {
		switch code {
		case http.StatusBadRequest:
			message = "bad request"
		case http.StatusInternalServerError:
			message = "something wrong on the server. contact server admin."
		}
	}

	return &apiResponseList{
		Code:    code,
		Message: message,
		Data:    map[string]interface{}{},
	}
}
