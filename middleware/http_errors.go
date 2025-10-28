package middleware

import (
	"net/http"
)

type HandlerWithError = func(writer http.ResponseWriter, req *http.Request) error
type ErrorHandler = func(writer http.ResponseWriter, req *http.Request, err error) error

func HttpErrorWrapper(next HandlerWithError, errHandler ErrorHandler) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := next(w, r)
		if err == nil {
			return
		}

		errHandler(w, r, err)
	}
}
