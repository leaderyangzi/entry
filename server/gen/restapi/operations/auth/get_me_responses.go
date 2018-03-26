// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/laincloud/entry/server/gen/models"
)

// GetMeOKCode is the HTTP code returned for type GetMeOK
const GetMeOKCode int = 200

/*GetMeOK the user corresponding to the access_token

swagger:response getMeOK
*/
type GetMeOK struct {

	/*
	  In: Body
	*/
	Payload *models.User `json:"body,omitempty"`
}

// NewGetMeOK creates GetMeOK with default headers values
func NewGetMeOK() *GetMeOK {

	return &GetMeOK{}
}

// WithPayload adds the payload to the get me o k response
func (o *GetMeOK) WithPayload(payload *models.User) *GetMeOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get me o k response
func (o *GetMeOK) SetPayload(payload *models.User) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMeOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetMeDefault generic error response

swagger:response getMeDefault
*/
type GetMeDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetMeDefault creates GetMeDefault with default headers values
func NewGetMeDefault(code int) *GetMeDefault {
	if code <= 0 {
		code = 500
	}

	return &GetMeDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get me default response
func (o *GetMeDefault) WithStatusCode(code int) *GetMeDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get me default response
func (o *GetMeDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get me default response
func (o *GetMeDefault) WithPayload(payload *models.Error) *GetMeDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get me default response
func (o *GetMeDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMeDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
