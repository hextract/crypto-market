// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/h4x4d/crypto-market/stack_connector/internal/models"
)

// GetTransactionsPurchaseOKCode is the HTTP code returned for type GetTransactionsPurchaseOK
const GetTransactionsPurchaseOKCode int = 200

/*
GetTransactionsPurchaseOK Successful operation

swagger:response getTransactionsPurchaseOK
*/
type GetTransactionsPurchaseOK struct {

	/*
	  In: Body
	*/
	Payload []*GetTransactionsPurchaseOKBodyItems0 `json:"body,omitempty"`
}

// NewGetTransactionsPurchaseOK creates GetTransactionsPurchaseOK with default headers values
func NewGetTransactionsPurchaseOK() *GetTransactionsPurchaseOK {

	return &GetTransactionsPurchaseOK{}
}

// WithPayload adds the payload to the get transactions purchase o k response
func (o *GetTransactionsPurchaseOK) WithPayload(payload []*GetTransactionsPurchaseOKBodyItems0) *GetTransactionsPurchaseOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get transactions purchase o k response
func (o *GetTransactionsPurchaseOK) SetPayload(payload []*GetTransactionsPurchaseOKBodyItems0) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTransactionsPurchaseOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*GetTransactionsPurchaseOKBodyItems0, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetTransactionsPurchaseBadRequestCode is the HTTP code returned for type GetTransactionsPurchaseBadRequest
const GetTransactionsPurchaseBadRequestCode int = 400

/*
GetTransactionsPurchaseBadRequest Incorrect data

swagger:response getTransactionsPurchaseBadRequest
*/
type GetTransactionsPurchaseBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetTransactionsPurchaseBadRequest creates GetTransactionsPurchaseBadRequest with default headers values
func NewGetTransactionsPurchaseBadRequest() *GetTransactionsPurchaseBadRequest {

	return &GetTransactionsPurchaseBadRequest{}
}

// WithPayload adds the payload to the get transactions purchase bad request response
func (o *GetTransactionsPurchaseBadRequest) WithPayload(payload *models.Error) *GetTransactionsPurchaseBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get transactions purchase bad request response
func (o *GetTransactionsPurchaseBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTransactionsPurchaseBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetTransactionsPurchaseUnauthorizedCode is the HTTP code returned for type GetTransactionsPurchaseUnauthorized
const GetTransactionsPurchaseUnauthorizedCode int = 401

/*
GetTransactionsPurchaseUnauthorized Unauthorized

swagger:response getTransactionsPurchaseUnauthorized
*/
type GetTransactionsPurchaseUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetTransactionsPurchaseUnauthorized creates GetTransactionsPurchaseUnauthorized with default headers values
func NewGetTransactionsPurchaseUnauthorized() *GetTransactionsPurchaseUnauthorized {

	return &GetTransactionsPurchaseUnauthorized{}
}

// WithPayload adds the payload to the get transactions purchase unauthorized response
func (o *GetTransactionsPurchaseUnauthorized) WithPayload(payload *models.Error) *GetTransactionsPurchaseUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get transactions purchase unauthorized response
func (o *GetTransactionsPurchaseUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTransactionsPurchaseUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
