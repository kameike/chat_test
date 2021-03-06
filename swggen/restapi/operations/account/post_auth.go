// Code generated by go-swagger; DO NOT EDIT.

package account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	swag "github.com/go-openapi/swag"
)

// PostAuthHandlerFunc turns a function with the right signature into a post auth handler
type PostAuthHandlerFunc func(PostAuthParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostAuthHandlerFunc) Handle(params PostAuthParams) middleware.Responder {
	return fn(params)
}

// PostAuthHandler interface for that can handle valid post auth params
type PostAuthHandler interface {
	Handle(PostAuthParams) middleware.Responder
}

// NewPostAuth creates a new http.Handler for the post auth operation
func NewPostAuth(ctx *middleware.Context, handler PostAuthHandler) *PostAuth {
	return &PostAuth{Context: ctx, Handler: handler}
}

/*PostAuth swagger:route POST /auth account postAuth

ログイン

サインアップもしくはアクセストークンの更新を行います

*/
type PostAuth struct {
	Context *middleware.Context
	Handler PostAuthHandler
}

func (o *PostAuth) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostAuthParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PostAuthBody post auth body
// swagger:model PostAuthBody
type PostAuthBody struct {

	// account hash
	AccountHash string `json:"accountHash,omitempty"`

	// auth token
	AuthToken string `json:"authToken,omitempty"`
}

// Validate validates this post auth body
func (o *PostAuthBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostAuthBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostAuthBody) UnmarshalBinary(b []byte) error {
	var res PostAuthBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
