// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// AddSampleHandlerFunc turns a function with the right signature into a add sample handler
type AddSampleHandlerFunc func(AddSampleParams) middleware.Responder

// Handle executing the request and returning a response
func (fn AddSampleHandlerFunc) Handle(params AddSampleParams) middleware.Responder {
	return fn(params)
}

// AddSampleHandler interface for that can handle valid add sample params
type AddSampleHandler interface {
	Handle(AddSampleParams) middleware.Responder
}

// NewAddSample creates a new http.Handler for the add sample operation
func NewAddSample(ctx *middleware.Context, handler AddSampleHandler) *AddSample {
	return &AddSample{Context: ctx, Handler: handler}
}

/*AddSample swagger:route POST /sample addSample

adds a sample

Adds a sample to the system

*/
type AddSample struct {
	Context *middleware.Context
	Handler AddSampleHandler
}

func (o *AddSample) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAddSampleParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
