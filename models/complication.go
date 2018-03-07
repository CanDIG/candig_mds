// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Complication complication
// swagger:model Complication
type Complication struct {

	// date
	Date *string `json:"date,omitempty" db:"date"`

	// late complication of therapy developed
	LateComplicationOfTherapyDeveloped *string `json:"late_complication_of_therapy_developed,omitempty" db:"late_complication_of_therapy_developed"`

	// late toxicity detail
	LateToxicityDetail *string `json:"late_toxicity_detail,omitempty" db:"late_toxicity_detail"`

	// patient id
	PatientID *string `json:"patient_id,omitempty" db:"patient_id"`

	// suspected treatment induced neoplasm developed
	SuspectedTreatmentInducedNeoplasmDeveloped *string `json:"suspected_treatment_induced_neoplasm_developed,omitempty" db:"suspected_treatment_induced_neoplasm_developed"`

	// treatment induced neoplasm details
	TreatmentInducedNeoplasmDetails *string `json:"treatment_induced_neoplasm_details,omitempty" db:"treatment_induced_neoplasm_details"`
}

// Validate validates this complication
func (m *Complication) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *Complication) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Complication) UnmarshalBinary(b []byte) error {
	var res Complication
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
