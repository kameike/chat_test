// Code generated by go-swagger; DO NOT EDIT.

package chatrooms

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPostChatroomsChannelHashReadParams creates a new PostChatroomsChannelHashReadParams object
// no default values defined in spec.
func NewPostChatroomsChannelHashReadParams() PostChatroomsChannelHashReadParams {

	return PostChatroomsChannelHashReadParams{}
}

// PostChatroomsChannelHashReadParams contains all the bound params for the post chatrooms channel hash read operation
// typically these are obtained from a http.Request
//
// swagger:parameters PostChatroomsChannelHashRead
type PostChatroomsChannelHashReadParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*長いハッシュ値
	  Required: true
	  In: path
	*/
	ChannelHash string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPostChatroomsChannelHashReadParams() beforehand.
func (o *PostChatroomsChannelHashReadParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rChannelHash, rhkChannelHash, _ := route.Params.GetOK("channel_hash")
	if err := o.bindChannelHash(rChannelHash, rhkChannelHash, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindChannelHash binds and validates parameter ChannelHash from path.
func (o *PostChatroomsChannelHashReadParams) bindChannelHash(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.ChannelHash = raw

	return nil
}