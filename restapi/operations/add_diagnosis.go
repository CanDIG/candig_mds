// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// AddDiagnosisHandlerFunc turns a function with the right signature into a add diagnosis handler
type AddDiagnosisHandlerFunc func(AddDiagnosisParams) middleware.Responder

// Handle executing the request and returning a response
func (fn AddDiagnosisHandlerFunc) Handle(params AddDiagnosisParams) middleware.Responder {
	return fn(params)
}

// AddDiagnosisHandler interface for that can handle valid add diagnosis params
type AddDiagnosisHandler interface {
	Handle(AddDiagnosisParams) middleware.Responder
}

// NewAddDiagnosis creates a new http.Handler for the add diagnosis operation
func NewAddDiagnosis(ctx *middleware.Context, handler AddDiagnosisHandler) *AddDiagnosis {
	return &AddDiagnosis{Context: ctx, Handler: handler}
}

/*AddDiagnosis swagger:route POST /diagnosis addDiagnosis

adds a diagnosis

Adds a diagnosis to the system

*/
type AddDiagnosis struct {
	Context *middleware.Context
	Handler AddDiagnosisHandler
}

func (o *AddDiagnosis) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAddDiagnosisParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
