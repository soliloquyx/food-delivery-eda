package http

import (
	"encoding/json"
	"net/http"

	"github.com/soliloquyx/food-delivery-eda/internal/gateway/adapters/in/http/placeorder"
)

func (h *handler) PlaceOrder(w http.ResponseWriter, r *http.Request) {
	var body placeorder.Request
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}

	in, err := placeorder.ToInput(body)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}

	result, err := h.order.Place(r.Context(), in)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "")
		return
	}

	resp := placeorder.ToResponse(result)

	writeJSON(w, http.StatusOK, resp)
}
