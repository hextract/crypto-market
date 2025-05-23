// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetLastMarketDataHandlerFunc turns a function with the right signature into a get last market data handler
type GetLastMarketDataHandlerFunc func(GetLastMarketDataParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetLastMarketDataHandlerFunc) Handle(params GetLastMarketDataParams) middleware.Responder {
	return fn(params)
}

// GetLastMarketDataHandler interface for that can handle valid get last market data params
type GetLastMarketDataHandler interface {
	Handle(GetLastMarketDataParams) middleware.Responder
}

// NewGetLastMarketData creates a new http.Handler for the get last market data operation
func NewGetLastMarketData(ctx *middleware.Context, handler GetLastMarketDataHandler) *GetLastMarketData {
	return &GetLastMarketData{Context: ctx, Handler: handler}
}

/*
	GetLastMarketData swagger:route GET /get_market_data getLastMarketData

Get last market data
*/
type GetLastMarketData struct {
	Context *middleware.Context
	Handler GetLastMarketDataHandler
}

func (o *GetLastMarketData) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetLastMarketDataParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
