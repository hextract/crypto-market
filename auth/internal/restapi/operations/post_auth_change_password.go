// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PostAuthChangePasswordHandlerFunc turns a function with the right signature into a post auth change password handler
type PostAuthChangePasswordHandlerFunc func(PostAuthChangePasswordParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostAuthChangePasswordHandlerFunc) Handle(params PostAuthChangePasswordParams) middleware.Responder {
	return fn(params)
}

// PostAuthChangePasswordHandler interface for that can handle valid post auth change password params
type PostAuthChangePasswordHandler interface {
	Handle(PostAuthChangePasswordParams) middleware.Responder
}

// NewPostAuthChangePassword creates a new http.Handler for the post auth change password operation
func NewPostAuthChangePassword(ctx *middleware.Context, handler PostAuthChangePasswordHandler) *PostAuthChangePassword {
	return &PostAuthChangePassword{Context: ctx, Handler: handler}
}

/*
	PostAuthChangePassword swagger:route POST /auth/change-password postAuthChangePassword

Change password
*/
type PostAuthChangePassword struct {
	Context *middleware.Context
	Handler PostAuthChangePasswordHandler
}

func (o *PostAuthChangePassword) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostAuthChangePasswordParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PostAuthChangePasswordBody post auth change password body
//
// swagger:model PostAuthChangePasswordBody
type PostAuthChangePasswordBody struct {

	// login
	Login string `json:"login,omitempty"`

	// new password
	NewPassword string `json:"newPassword,omitempty"`

	// old password
	OldPassword string `json:"oldPassword,omitempty"`
}

// Validate validates this post auth change password body
func (o *PostAuthChangePasswordBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post auth change password body based on context it is used
func (o *PostAuthChangePasswordBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostAuthChangePasswordBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostAuthChangePasswordBody) UnmarshalBinary(b []byte) error {
	var res PostAuthChangePasswordBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostAuthChangePasswordOKBody post auth change password o k body
//
// swagger:model PostAuthChangePasswordOKBody
type PostAuthChangePasswordOKBody struct {

	// token
	Token string `json:"token,omitempty"`
}

// Validate validates this post auth change password o k body
func (o *PostAuthChangePasswordOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post auth change password o k body based on context it is used
func (o *PostAuthChangePasswordOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostAuthChangePasswordOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostAuthChangePasswordOKBody) UnmarshalBinary(b []byte) error {
	var res PostAuthChangePasswordOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
