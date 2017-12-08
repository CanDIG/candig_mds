// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetIndividualParams creates a new GetIndividualParams object
// with the default values initialized.
func NewGetIndividualParams() GetIndividualParams {
	var ()
	return GetIndividualParams{}
}

// GetIndividualParams contains all the bound params for the get individual operation
// typically these are obtained from a http.Request
//
// swagger:parameters getIndividual
type GetIndividualParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: path
	*/
	IndividualID strfmt.UUID
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *GetIndividualParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	rIndividualID, rhkIndividualID, _ := route.Params.GetOK("individualId")
	if err := o.bindIndividualID(rIndividualID, rhkIndividualID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetIndividualParams) bindIndividualID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	value, err := formats.Parse("uuid", raw)
	if err != nil {
		return errors.InvalidType("individualId", "path", "strfmt.UUID", raw)
	}
	o.IndividualID = *(value.(*strfmt.UUID))

	return nil
}
