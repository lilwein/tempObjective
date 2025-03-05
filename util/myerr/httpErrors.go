package myerr

import (
	"fmt"
	"net/http"
)

var nil_p = HttpErr{}

// Nil value
var Nil *HttpErr = &nil_p

// Struct for http errors
type HttpErr struct {
	Message   string
	Error     error
	HttpError int
}

// Return a new httpErr
func NewHttpErr(msg string, err error, httpError int) *HttpErr {
	var new = &HttpErr{
		Message:   msg,
		Error:     err,
		HttpError: httpError,
	}
	fmt.Println(new)

	return new
}

// Error message
func (err *HttpErr) message() string {
	var message string = ""
	if err.Message != "" {
		message = message + err.Message
	}
	if err.Error != nil {
		message = message + err.Error.Error()
	}
	return message
}

// Returns true if err is not nil
func (err *HttpErr) HTTPErrors(w http.ResponseWriter) bool {
	if !err.isNull() {
		msg := err.message()
		http.Error(w, msg, err.HttpError)
		return true
	}
	return false
}

// Returns true if the error is null
func (err *HttpErr) isNull() bool {
	return err.Message == "" && err.Error == nil && err.HttpError == 0
}
