// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/h4x4d/crypto-market/main/internal/models"
)

// CreateBidHandlerFunc turns a function with the right signature into a create bid handler
type CreateBidHandlerFunc func(CreateBidParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateBidHandlerFunc) Handle(params CreateBidParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// CreateBidHandler interface for that can handle valid create bid params
type CreateBidHandler interface {
	Handle(CreateBidParams, *models.User) middleware.Responder
}

// NewCreateBid creates a new http.Handler for the create bid operation
func NewCreateBid(ctx *middleware.Context, handler CreateBidHandler) *CreateBid {
	return &CreateBid{Context: ctx, Handler: handler}
}

/*
	CreateBid swagger:route POST /bid createBid

Create bid
*/
type CreateBid struct {
	Context *middleware.Context
	Handler CreateBidHandler
}

func (o *CreateBid) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewCreateBidParams()
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

// CreateBidBody create bid body
//
// swagger:model CreateBidBody
type CreateBidBody struct {

	// amount to buy
	// Required: true
	// Minimum: 0
	AmountToBuy *float32 `json:"amount_to_buy"`

	// buy speed
	// Minimum: 0
	BuySpeed *float32 `json:"buy_speed,omitempty"`

	// from currency
	// Required: true
	// Enum: ["USDT","BTC"]
	FromCurrency *string `json:"from_currency"`

	// max price
	// Required: true
	// Minimum: 0
	MaxPrice *float32 `json:"max_price"`

	// min price
	// Required: true
	// Minimum: 0
	MinPrice *float32 `json:"min_price"`

	// to currency
	// Required: true
	// Enum: ["USDT","BTC"]
	ToCurrency *string `json:"to_currency"`
}

// Validate validates this create bid body
func (o *CreateBidBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAmountToBuy(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateBuySpeed(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateFromCurrency(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMaxPrice(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMinPrice(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateToCurrency(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateBidBody) validateAmountToBuy(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"amount_to_buy", "body", o.AmountToBuy); err != nil {
		return err
	}

	if err := validate.Minimum("body"+"."+"amount_to_buy", "body", float64(*o.AmountToBuy), 0, false); err != nil {
		return err
	}

	return nil
}

func (o *CreateBidBody) validateBuySpeed(formats strfmt.Registry) error {
	if swag.IsZero(o.BuySpeed) { // not required
		return nil
	}

	if err := validate.Minimum("body"+"."+"buy_speed", "body", float64(*o.BuySpeed), 0, false); err != nil {
		return err
	}

	return nil
}

var createBidBodyTypeFromCurrencyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["USDT","BTC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createBidBodyTypeFromCurrencyPropEnum = append(createBidBodyTypeFromCurrencyPropEnum, v)
	}
}

const (

	// CreateBidBodyFromCurrencyUSDT captures enum value "USDT"
	CreateBidBodyFromCurrencyUSDT string = "USDT"

	// CreateBidBodyFromCurrencyBTC captures enum value "BTC"
	CreateBidBodyFromCurrencyBTC string = "BTC"
)

// prop value enum
func (o *CreateBidBody) validateFromCurrencyEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createBidBodyTypeFromCurrencyPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *CreateBidBody) validateFromCurrency(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"from_currency", "body", o.FromCurrency); err != nil {
		return err
	}

	// value enum
	if err := o.validateFromCurrencyEnum("body"+"."+"from_currency", "body", *o.FromCurrency); err != nil {
		return err
	}

	return nil
}

func (o *CreateBidBody) validateMaxPrice(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"max_price", "body", o.MaxPrice); err != nil {
		return err
	}

	if err := validate.Minimum("body"+"."+"max_price", "body", float64(*o.MaxPrice), 0, false); err != nil {
		return err
	}

	return nil
}

func (o *CreateBidBody) validateMinPrice(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"min_price", "body", o.MinPrice); err != nil {
		return err
	}

	if err := validate.Minimum("body"+"."+"min_price", "body", float64(*o.MinPrice), 0, false); err != nil {
		return err
	}

	return nil
}

var createBidBodyTypeToCurrencyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["USDT","BTC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createBidBodyTypeToCurrencyPropEnum = append(createBidBodyTypeToCurrencyPropEnum, v)
	}
}

const (

	// CreateBidBodyToCurrencyUSDT captures enum value "USDT"
	CreateBidBodyToCurrencyUSDT string = "USDT"

	// CreateBidBodyToCurrencyBTC captures enum value "BTC"
	CreateBidBodyToCurrencyBTC string = "BTC"
)

// prop value enum
func (o *CreateBidBody) validateToCurrencyEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createBidBodyTypeToCurrencyPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *CreateBidBody) validateToCurrency(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"to_currency", "body", o.ToCurrency); err != nil {
		return err
	}

	// value enum
	if err := o.validateToCurrencyEnum("body"+"."+"to_currency", "body", *o.ToCurrency); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create bid body based on context it is used
func (o *CreateBidBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateBidBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateBidBody) UnmarshalBinary(b []byte) error {
	var res CreateBidBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// CreateBidOKBody create bid o k body
//
// swagger:model CreateBidOKBody
type CreateBidOKBody struct {

	// id
	// Example: bid_123
	ID string `json:"id,omitempty"`
}

// Validate validates this create bid o k body
func (o *CreateBidOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create bid o k body based on context it is used
func (o *CreateBidOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateBidOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateBidOKBody) UnmarshalBinary(b []byte) error {
	var res CreateBidOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
