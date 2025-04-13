// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/h4x4d/crypto-market/main/internal/models"
)

// GetTransactionsTransfersHandlerFunc turns a function with the right signature into a get transactions transfers handler
type GetTransactionsTransfersHandlerFunc func(GetTransactionsTransfersParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn GetTransactionsTransfersHandlerFunc) Handle(params GetTransactionsTransfersParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// GetTransactionsTransfersHandler interface for that can handle valid get transactions transfers params
type GetTransactionsTransfersHandler interface {
	Handle(GetTransactionsTransfersParams, *models.User) middleware.Responder
}

// NewGetTransactionsTransfers creates a new http.Handler for the get transactions transfers operation
func NewGetTransactionsTransfers(ctx *middleware.Context, handler GetTransactionsTransfersHandler) *GetTransactionsTransfers {
	return &GetTransactionsTransfers{Context: ctx, Handler: handler}
}

/*
	GetTransactionsTransfers swagger:route GET /transactions/transfers getTransactionsTransfers

# Get withdrawal and deposit history

Returns the user's withdrawal and deposit history with optional filters
*/
type GetTransactionsTransfers struct {
	Context *middleware.Context
	Handler GetTransactionsTransfersHandler
}

func (o *GetTransactionsTransfers) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetTransactionsTransfersParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal *models.User
	if uprinc != nil {
		principal = uprinc.(*models.User) // this is really a models.User, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
