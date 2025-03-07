// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetBidByIDParams creates a new GetBidByIDParams object
//
// There are no default values defined in the spec.
func NewGetBidByIDParams() GetBidByIDParams {

	return GetBidByIDParams{}
}

// GetBidByIDParams contains all the bound params for the get bid by id operation
// typically these are obtained from a http.Request
//
// swagger:parameters get_bid_by_id
type GetBidByIDParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*ID of bid to return
	  Required: true
	  In: path
	*/
	BidID int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetBidByIDParams() beforehand.
func (o *GetBidByIDParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rBidID, rhkBidID, _ := route.Params.GetOK("bid_id")
	if err := o.bindBidID(rBidID, rhkBidID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindBidID binds and validates parameter BidID from path.
func (o *GetBidByIDParams) bindBidID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("bid_id", "path", "int64", raw)
	}
	o.BidID = value

	return nil
}
