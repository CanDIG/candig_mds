// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeleteExperimentHandlerFunc turns a function with the right signature into a delete experiment handler
type DeleteExperimentHandlerFunc func(DeleteExperimentParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteExperimentHandlerFunc) Handle(params DeleteExperimentParams) middleware.Responder {
	return fn(params)
}

// DeleteExperimentHandler interface for that can handle valid delete experiment params
type DeleteExperimentHandler interface {
	Handle(DeleteExperimentParams) middleware.Responder
}

// NewDeleteExperiment creates a new http.Handler for the delete experiment operation
func NewDeleteExperiment(ctx *middleware.Context, handler DeleteExperimentHandler) *DeleteExperiment {
	return &DeleteExperiment{Context: ctx, Handler: handler}
}

/*DeleteExperiment swagger:route DELETE /experiment/{id} deleteExperiment

deletes an experiment

Deletes an experiment to the system

*/
type DeleteExperiment struct {
	Context *middleware.Context
	Handler DeleteExperimentHandler
}

func (o *DeleteExperiment) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteExperimentParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}