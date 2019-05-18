// Code generated by go-swagger; DO NOT EDIT.

package apimodel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ChatroomRequest chatroom request
// swagger:model chatroomRequest
type ChatroomRequest struct {

	// chatrooms
	Chatrooms []string `json:"chatrooms"`
}

// Validate validates this chatroom request
func (m *ChatroomRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ChatroomRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChatroomRequest) UnmarshalBinary(b []byte) error {
	var res ChatroomRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
