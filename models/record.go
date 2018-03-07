package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/swag"
)

// Record record
// swagger:model Record
type Record struct {

	// complication
	Complication *Complication

	// consent
	Consent *Consent

	// diagnosis
	Diagnosis *Diagnosis

	// enrollment
	Enrollment *Enrollment

	// outcome
	Outcome *Outcome

	// patient
	Patient *Patient

	// sample
	Sample *Sample

	// treatment
	Treatment *Treatment

	// tumourboard
	Tumourboard *Tumourboard
}

// MarshalBinary interface implementation
func (m *Record) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Record) UnmarshalBinary(b []byte) error {
	var res Record
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
