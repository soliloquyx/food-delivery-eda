package httpapi

import (
	"encoding/json"
	"net/http"

	"github.com/soliloquyx/food-delivery-eda/internal/gateway/adapters/in/httpapi/placeorder"
)

func (h *handler) PlaceOrder(w http.ResponseWriter, r *http.Request) {
	var body placeorder.Request
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}

	in := placeorder.ToInput(body)

	if err := h.order.Place(r.Context(), in); err != nil {
		writeError(w, http.StatusInternalServerError, "")
		return
	}

	writeJSON(w, http.StatusOK, nil)
}
