package http

import (
	"errors"
	"net/http"
)

const (
	codeInvalidJSON      = "invalid_json"
	codeValidationFailed = "validation_failed"
	codeInternal         = "internal"
)

var (
	errInvalidJSON = errors.New(codeInvalidJSON)
	errValidation  = errors.New(codeValidationFailed)
)

type errorResponse struct {
	Code      string `json:"code"`
	Message   string `json:"message"`
	RequestID string `json:"request_id,omitempty"`
}

func mapError(err error) (status int, code, msg string) {
	switch {
	case errors.Is(err, errInvalidJSON):
		return http.StatusBadRequest, codeInvalidJSON, "Invalid JSON body"
	case errors.Is(err, errValidation):
		return http.StatusBadRequest, codeValidationFailed, "Validation failed"
	default:
		return http.StatusInternalServerError, codeInternal, http.StatusText(http.StatusInternalServerError)
	}
}
