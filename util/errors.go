package util

type HTTPError string

const HTTPError400 HTTPError = "bad request"
const HTTPError404 HTTPError = "not found"
const HTTPError405 HTTPError = "method not allowed"

type LoginFailedError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"invalid email or password"`
}

const LoginFailed HTTPError = "invalid email or password"
