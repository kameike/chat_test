// Code generated by go-swagger; DO NOT EDIT.

package chat_rooms

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	apimodel "github.com/kameike/chat_api/swggen/apimodel"
)

// GetChatroomsIDMessagesOKCode is the HTTP code returned for type GetChatroomsIDMessagesOK
const GetChatroomsIDMessagesOKCode int = 200

/*GetChatroomsIDMessagesOK ok

swagger:response getChatroomsIdMessagesOK
*/
type GetChatroomsIDMessagesOK struct {

	/*
	  In: Body
	*/
	Payload []*apimodel.Chat `json:"body,omitempty"`
}

// NewGetChatroomsIDMessagesOK creates GetChatroomsIDMessagesOK with default headers values
func NewGetChatroomsIDMessagesOK() *GetChatroomsIDMessagesOK {

	return &GetChatroomsIDMessagesOK{}
}

// WithPayload adds the payload to the get chatrooms Id messages o k response
func (o *GetChatroomsIDMessagesOK) WithPayload(payload []*apimodel.Chat) *GetChatroomsIDMessagesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get chatrooms Id messages o k response
func (o *GetChatroomsIDMessagesOK) SetPayload(payload []*apimodel.Chat) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetChatroomsIDMessagesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make([]*apimodel.Chat, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}