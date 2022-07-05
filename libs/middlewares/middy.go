package middlewares

import "net/http"

type Handler = func(response http.ResponseWriter, request *http.Request)

type MiddlewareImpl interface {
	Wrap(handler Handler) Handler
}

func Middy(handler Handler, middlewares ...MiddlewareImpl) Handler {
	current := handler
	for _, mw := range middlewares {
		current = mw.Wrap(current)
	}

	return func(response http.ResponseWriter, request *http.Request) {
		current(response, request)
	}
}
