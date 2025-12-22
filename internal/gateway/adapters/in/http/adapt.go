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

		defer func() {
			if rec := recover(); rec != nil {
				logger.Error(
					"panic",
					zap.String("request_id", reqID),
					zap.Any("recover", rec),
					zap.String("method", r.Method),
					zap.String("path", r.URL.Path),
					zap.Duration("dur", time.Since(start)),
				)

				writeError(
					w,
					stdhttp.StatusInternalServerError,
					codeInternal,
					stdhttp.StatusText(stdhttp.StatusInternalServerError),
					reqID,
				)
			}
		}()

		if err := next(w, r); err != nil {
			status, code, msg := mapError(err)

			logger.Error(
				"http request failed",
				zap.String("request_id", reqID),
				zap.String("method", r.Method),
				zap.String("path", r.URL.Path),
				zap.Int("status", status),
				zap.String("code", code),
				zap.Duration("dur", time.Since(start)),
				zap.Error(err),
			)

			writeError(
				w,
				status,
				code,
				msg,
				reqID,
			)
		}
	}
}
