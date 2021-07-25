package models

// UrlProcessingError is the response when there is error
type UrlProcessingError struct {
	Code uint16 `json:"code"`
	Msg  string `json:"msg"`
}

// NoError
var NoError = UrlProcessingError{Code: 00000, Msg: "no error"}

// NotFoundError shows the error response body when url data not found
var NotFoundError = UrlProcessingError{Code: 20401, Msg: "Shortening url not found"}

// BadRequestError shows the error response body when bad request happens
var BadRequestError = UrlProcessingError{Code: 40001, Msg: "Bad request"}

// InternalServerError shows the error response body if internal server error
var InternalServerError = UrlProcessingError{Code: 50001, Msg: "Internal server error"}
