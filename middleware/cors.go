package middleware

import "net/http"

type CorsOptions struct {
	Origin  string
	Methods string
	Headers string
}

func Cors(next http.Handler, options *CorsOptions) http.Handler {
	var defaultOrigin = "*"
	var defaultMethods = "*"
	var defaultHeaders = "*"

	if options.Origin != "" {
		defaultOrigin = options.Origin
	}
	if options.Methods != "" {
		defaultMethods = options.Methods
	}
	if options.Headers != "" {
		defaultHeaders = options.Headers
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", defaultOrigin)
		w.Header().Set("Access-Control-Allow-Methods", defaultMethods)
		w.Header().Set("Access-Control-Allow-Headers", defaultHeaders)
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
