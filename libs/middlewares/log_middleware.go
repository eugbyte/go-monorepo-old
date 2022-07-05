package middlewares

import (
	"net/http"

	"github.com/go-monorepo/libs/util"
)

type LogMiddleware struct{}

func (mw LogMiddleware) Wrap(handler Handler) Handler {
	return func(response http.ResponseWriter, request *http.Request) {
		// pre-process request here
		util.Trace("LogMiddleware", "pre-processing request...")
		handler(response, request)
		// post-process reponse here
	}
}
