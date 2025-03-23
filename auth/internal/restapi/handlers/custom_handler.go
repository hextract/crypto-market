package handlers

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"net/http"
)

type CustomResponder func(http.ResponseWriter, runtime.Producer)

func (c CustomResponder) WriteResponse(w http.ResponseWriter, p runtime.Producer) {
	c(w, p)
}
func NewCustomResponder(r *http.Request, h http.Handler) middleware.Responder {
	return CustomResponder(func(w http.ResponseWriter, _ runtime.Producer) {
		h.ServeHTTP(w, r)
	})
}
