package middleware

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/soliloquyx/food-delivery-eda/internal/gateway/adapters/in/http/middleware/requestid"
)

func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.Header.Get("X-Request-Id")
		if id == "" {
			id = uuid.NewString()
		}

		w.Header().Set("X-Request-Id", id)

		ctx := requestid.With(r.Context(), id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
