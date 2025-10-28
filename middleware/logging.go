package middleware

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type wrappedWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *wrappedWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.statusCode = statusCode
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		writer := &wrappedWriter{
			w,
			http.StatusOK,
		}

		next.ServeHTTP(writer, r)

		log.Println(writer.statusCode, r.Method, r.URL.Path, fmt.Sprintf("%vms", time.Since(start).Milliseconds()))
	})
}
