package wrapper

import (
	"net/http"
)

type HandlerWithError func(writer http.ResponseWriter, req *http.Request) error
type HandleError func(writer http.ResponseWriter, req *http.Request, err error) error

func HandleHttpError(next HandlerWithError, handleError HandleError) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := next(w, r)
		if err == nil {
			return
		}

		handleError(w, r, err)
	}
}
