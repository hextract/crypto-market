// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/h4x4d/crypto-market/stack_connector/internal/models"
)

// GetTransactionsTransfersOKCode is the HTTP code returned for type GetTransactionsTransfersOK
const GetTransactionsTransfersOKCode int = 200

/*
GetTransactionsTransfersOK Success operation

swagger:response getTransactionsTransfersOK
*/
type GetTransactionsTransfersOK struct {

	/*
	  In: Body
	*/
	Payload []*GetTransactionsTransfersOKBodyItems0 `json:"body,omitempty"`
}

// NewGetTransactionsTransfersOK creates GetTransactionsTransfersOK with default headers values
func NewGetTransactionsTransfersOK() *GetTransactionsTransfersOK {

	return &GetTransactionsTransfersOK{}
}

// WithPayload adds the payload to the get transactions transfers o k response
func (o *GetTransactionsTransfersOK) WithPayload(payload []*GetTransactionsTransfersOKBodyItems0) *GetTransactionsTransfersOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get transactions transfers o k response
func (o *GetTransactionsTransfersOK) SetPayload(payload []*GetTransactionsTransfersOKBodyItems0) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTransactionsTransfersOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*GetTransactionsTransfersOKBodyItems0, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetTransactionsTransfersUnauthorizedCode is the HTTP code returned for type GetTransactionsTransfersUnauthorized
const GetTransactionsTransfersUnauthorizedCode int = 401

/*
GetTransactionsTransfersUnauthorized Unauthorized

swagger:response getTransactionsTransfersUnauthorized
*/
type GetTransactionsTransfersUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetTransactionsTransfersUnauthorized creates GetTransactionsTransfersUnauthorized with default headers values
func NewGetTransactionsTransfersUnauthorized() *GetTransactionsTransfersUnauthorized {

	return &GetTransactionsTransfersUnauthorized{}
}

// WithPayload adds the payload to the get transactions transfers unauthorized response
func (o *GetTransactionsTransfersUnauthorized) WithPayload(payload *models.Error) *GetTransactionsTransfersUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get transactions transfers unauthorized response
func (o *GetTransactionsTransfersUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTransactionsTransfersUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetTransactionsTransfersConflictCode is the HTTP code returned for type GetTransactionsTransfersConflict
const GetTransactionsTransfersConflictCode int = 409

/*
GetTransactionsTransfersConflict Incorrect data

swagger:response getTransactionsTransfersConflict
*/
type GetTransactionsTransfersConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetTransactionsTransfersConflict creates GetTransactionsTransfersConflict with default headers values
func NewGetTransactionsTransfersConflict() *GetTransactionsTransfersConflict {

	return &GetTransactionsTransfersConflict{}
}

// WithPayload adds the payload to the get transactions transfers conflict response
func (o *GetTransactionsTransfersConflict) WithPayload(payload *models.Error) *GetTransactionsTransfersConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get transactions transfers conflict response
func (o *GetTransactionsTransfersConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTransactionsTransfersConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
