package http

import (
	"errors"
	"net/http"

	"github.com/soliloquyx/food-delivery-eda/internal/gateway/app/order"
)

const (
	codeInvalidJSON     = "invalid_json"
	codeInvalidArgument = "invalid_argument"
	codeInternal        = "internal"
)

var (
	errInvalidJSON = errors.New("invalid json")
)

type errorResponse struct {
	Code      string `json:"code"`
	RequestID string `json:"request_id,omitempty"`
}

func mapError(err error) (status int, code string) {
	switch {
	case errors.Is(err, errInvalidJSON):
		return http.StatusBadRequest, codeInvalidJSON
	case errors.Is(err, order.ErrInvalidUUID), errors.Is(err, order.ErrInvalidFulfillmentType):
		return http.StatusBadRequest, codeInvalidArgument
	default:
		return http.StatusInternalServerError, codeInternal
	}
}
