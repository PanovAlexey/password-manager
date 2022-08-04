package trace

import (
	"api-gw/internal/application/service"
	"context"
	"fmt"
	"google.golang.org/grpc/metadata"
	"net/http"
	"time"
)

type Logger interface {
	Debug(args ...interface{})
}

func Trace(logger Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			var traceId string
			traceId = fmt.Sprintf("%d", time.Now().UTC().UnixNano())
			ctx := context.WithValue(r.Context(), service.TraceIdKey, traceId)

			ctx = metadata.NewOutgoingContext(
				ctx,
				metadata.New(map[string]string{
					service.TraceIdKey: traceId,
				}),
			)

			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)

			logger.Debug(fmt.Sprintf(
				"Call:%v, method:%v, traceId: %v, time: %v",
				r.Host+r.RequestURI, r.Method, traceId, time.Since(start),
			))
		})
	}
}
