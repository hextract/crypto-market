// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"generator/internal/infrastructure/dto"
)

// SendContinuousOrderReader is a Reader for the SendContinuousOrder structure.
type SendContinuousOrderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SendContinuousOrderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSendContinuousOrderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewSendContinuousOrderInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /create-order] sendContinuousOrder", response, response.Code())
	}
}

// NewSendContinuousOrderOK creates a SendContinuousOrderOK with default headers values
func NewSendContinuousOrderOK() *SendContinuousOrderOK {
	return &SendContinuousOrderOK{}
}

/*
SendContinuousOrderOK describes a response with status code 200, with default header values.

Order successfully accepted
*/
type SendContinuousOrderOK struct {
	Payload *dto.SuccessResponse
}

// IsSuccess returns true when this send continuous order o k response has a 2xx status code
func (o *SendContinuousOrderOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this send continuous order o k response has a 3xx status code
func (o *SendContinuousOrderOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this send continuous order o k response has a 4xx status code
func (o *SendContinuousOrderOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this send continuous order o k response has a 5xx status code
func (o *SendContinuousOrderOK) IsServerError() bool {
	return false
}

// IsCode returns true when this send continuous order o k response a status code equal to that given
func (o *SendContinuousOrderOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the send continuous order o k response
func (o *SendContinuousOrderOK) Code() int {
	return 200
}

func (o *SendContinuousOrderOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /create-order][%d] sendContinuousOrderOK %s", 200, payload)
}

func (o *SendContinuousOrderOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /create-order][%d] sendContinuousOrderOK %s", 200, payload)
}

func (o *SendContinuousOrderOK) GetPayload() *dto.SuccessResponse {
	return o.Payload
}

func (o *SendContinuousOrderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dto.SuccessResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendContinuousOrderInternalServerError creates a SendContinuousOrderInternalServerError with default headers values
func NewSendContinuousOrderInternalServerError() *SendContinuousOrderInternalServerError {
	return &SendContinuousOrderInternalServerError{}
}

/*
SendContinuousOrderInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type SendContinuousOrderInternalServerError struct {
	Payload *dto.Error
}

// IsSuccess returns true when this send continuous order internal server error response has a 2xx status code
func (o *SendContinuousOrderInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this send continuous order internal server error response has a 3xx status code
func (o *SendContinuousOrderInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this send continuous order internal server error response has a 4xx status code
func (o *SendContinuousOrderInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this send continuous order internal server error response has a 5xx status code
func (o *SendContinuousOrderInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this send continuous order internal server error response a status code equal to that given
func (o *SendContinuousOrderInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the send continuous order internal server error response
func (o *SendContinuousOrderInternalServerError) Code() int {
	return 500
}

func (o *SendContinuousOrderInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /create-order][%d] sendContinuousOrderInternalServerError %s", 500, payload)
}

func (o *SendContinuousOrderInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /create-order][%d] sendContinuousOrderInternalServerError %s", 500, payload)
}

func (o *SendContinuousOrderInternalServerError) GetPayload() *dto.Error {
	return o.Payload
}

func (o *SendContinuousOrderInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dto.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
