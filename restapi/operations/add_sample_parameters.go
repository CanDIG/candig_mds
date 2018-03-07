// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/CanDIG/candig_mds/models"
)

// NewAddSampleParams creates a new AddSampleParams object
// with the default values initialized.
func NewAddSampleParams() AddSampleParams {
	var ()
	return AddSampleParams{}
}

// AddSampleParams contains all the bound params for the add sample operation
// typically these are obtained from a http.Request
//
// swagger:parameters addSample
type AddSampleParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Sample
	  In: body
	*/
	Sample *models.Sample
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *AddSampleParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.Sample
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("sample", "body", "", err))
		} else {
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Sample = &body
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
