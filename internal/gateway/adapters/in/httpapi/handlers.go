package httpapi

import "net/http"

func (h *handler) PlaceOrder(w http.ResponseWriter, r *http.Request) {
	if err := h.Order.Place(); err != nil {
		writeError(w, http.StatusInternalServerError, "")
	}

	writeJSON(w, http.StatusOK, nil)
}
