package http

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/soliloquyx/food-delivery-eda/internal/gateway/adapters/in/http/placeorder"
)

func (h *handler) PlaceOrder(w http.ResponseWriter, r *http.Request) error {
	var body placeorder.Request
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return errors.Join(errInvalidJSON, err)
	}

	in, err := placeorder.ToInput(body)
	if err != nil {
		return errors.Join(errValidation, err)
	}

	result, err := h.order.PlaceOrder(r.Context(), in)
	if err != nil {
		return err
	}

	resp := placeorder.ToResponse(result)

	writeJSON(w, http.StatusOK, resp)
	return nil
}
