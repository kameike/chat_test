// Code generated by go-swagger; DO NOT EDIT.

package account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	apimodel "github.com/kameike/chat_api/swggen/apimodel"
)

// PostAuthOKCode is the HTTP code returned for type PostAuthOK
const PostAuthOKCode int = 200

/*PostAuthOK ok

swagger:response postAuthOK
*/
type PostAuthOK struct {

	/*
	  In: Body
	*/
	Payload *apimodel.AuthInfo `json:"body,omitempty"`
}

// NewPostAuthOK creates PostAuthOK with default headers values
func NewPostAuthOK() *PostAuthOK {

	return &PostAuthOK{}
}

// WithPayload adds the payload to the post auth o k response
func (o *PostAuthOK) WithPayload(payload *apimodel.AuthInfo) *PostAuthOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post auth o k response
func (o *PostAuthOK) SetPayload(payload *apimodel.AuthInfo) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostAuthOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}