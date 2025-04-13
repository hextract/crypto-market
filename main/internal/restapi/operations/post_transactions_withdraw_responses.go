// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/h4x4d/crypto-market/main/internal/models"
)

// PostTransactionsWithdrawOKCode is the HTTP code returned for type PostTransactionsWithdrawOK
const PostTransactionsWithdrawOKCode int = 200

/*
PostTransactionsWithdrawOK Successful operation

swagger:response postTransactionsWithdrawOK
*/
type PostTransactionsWithdrawOK struct {

	/*
	  In: Body
	*/
	Payload *PostTransactionsWithdrawOKBody `json:"body,omitempty"`
}

// NewPostTransactionsWithdrawOK creates PostTransactionsWithdrawOK with default headers values
func NewPostTransactionsWithdrawOK() *PostTransactionsWithdrawOK {

	return &PostTransactionsWithdrawOK{}
}

// WithPayload adds the payload to the post transactions withdraw o k response
func (o *PostTransactionsWithdrawOK) WithPayload(payload *PostTransactionsWithdrawOKBody) *PostTransactionsWithdrawOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post transactions withdraw o k response
func (o *PostTransactionsWithdrawOK) SetPayload(payload *PostTransactionsWithdrawOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostTransactionsWithdrawOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostTransactionsWithdrawBadRequestCode is the HTTP code returned for type PostTransactionsWithdrawBadRequest
const PostTransactionsWithdrawBadRequestCode int = 400

/*
PostTransactionsWithdrawBadRequest Incorrect data or insufficient balance

swagger:response postTransactionsWithdrawBadRequest
*/
type PostTransactionsWithdrawBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostTransactionsWithdrawBadRequest creates PostTransactionsWithdrawBadRequest with default headers values
func NewPostTransactionsWithdrawBadRequest() *PostTransactionsWithdrawBadRequest {

	return &PostTransactionsWithdrawBadRequest{}
}

// WithPayload adds the payload to the post transactions withdraw bad request response
func (o *PostTransactionsWithdrawBadRequest) WithPayload(payload *models.Error) *PostTransactionsWithdrawBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post transactions withdraw bad request response
func (o *PostTransactionsWithdrawBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostTransactionsWithdrawBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostTransactionsWithdrawUnauthorizedCode is the HTTP code returned for type PostTransactionsWithdrawUnauthorized
const PostTransactionsWithdrawUnauthorizedCode int = 401

/*
PostTransactionsWithdrawUnauthorized Unauthorized

swagger:response postTransactionsWithdrawUnauthorized
*/
type PostTransactionsWithdrawUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostTransactionsWithdrawUnauthorized creates PostTransactionsWithdrawUnauthorized with default headers values
func NewPostTransactionsWithdrawUnauthorized() *PostTransactionsWithdrawUnauthorized {

	return &PostTransactionsWithdrawUnauthorized{}
}

// WithPayload adds the payload to the post transactions withdraw unauthorized response
func (o *PostTransactionsWithdrawUnauthorized) WithPayload(payload *models.Error) *PostTransactionsWithdrawUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post transactions withdraw unauthorized response
func (o *PostTransactionsWithdrawUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostTransactionsWithdrawUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
