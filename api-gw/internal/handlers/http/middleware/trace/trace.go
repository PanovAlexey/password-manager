package trace

import (
	"context"
	"fmt"
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
			traceIdName := "trace-id" //@toDo move it to service

			ctx := context.WithValue(r.Context(), traceIdName, traceId)
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)

			logger.Debug(fmt.Sprintf(
				"Call:%v, method:%v, traceId: %v, time: %v",
				r.Host+r.RequestURI, r.Method, traceId, time.Since(start),
			))
		})
	}
}
