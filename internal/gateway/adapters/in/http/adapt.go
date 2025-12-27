package http

import (
	stdhttp "net/http"
	"time"

	"github.com/soliloquyx/food-delivery-eda/internal/gateway/adapters/in/http/middleware/requestid"
	"go.uber.org/zap"
)

type endpoint func(w stdhttp.ResponseWriter, r *stdhttp.Request) error

func Adapt(logger *zap.Logger, next endpoint) stdhttp.HandlerFunc {
	return func(w stdhttp.ResponseWriter, r *stdhttp.Request) {
		start := time.Now()
		reqID, _ := requestid.From(r.Context())

		l := logger.With(
			zap.String("request_id", reqID),
			zap.String("method", r.Method),
			zap.String("path", r.URL.Path),
		)

		defer func() {
			if rec := recover(); rec != nil {
				l.Error(
					"panic",
					zap.Any("recover", rec),
					zap.String("code", codeInternal),
					zap.Duration("dur", time.Since(start)),
				)

				writeJSON(
					w,
					stdhttp.StatusInternalServerError,
					errorResponse{
						Code:      codeInternal,
						Message:   stdhttp.StatusText(stdhttp.StatusInternalServerError),
						RequestID: reqID,
					},
				)
			}
		}()

		if err := next(w, r); err != nil {
			status, code, msg := mapError(err)

			l.Error(
				"http request failed",
				zap.Int("status", status),
				zap.String("code", code),
				zap.Duration("dur", time.Since(start)),
				zap.Error(err),
			)

			writeJSON(
				w,
				status,
				errorResponse{
					Code:      code,
					Message:   msg,
					RequestID: reqID,
				},
			)
		}
	}
}
