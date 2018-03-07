// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Consent consent
// swagger:model Consent
type Consent struct {

	// assent form version
	AssentFormVersion *string `json:"assent_form_version,omitempty" db:"assent_form_version"`

	// consent date
	ConsentDate *string `json:"consent_date,omitempty" db:"consent_date"`

	// consent form complete
	ConsentFormComplete *string `json:"consent_form_complete,omitempty" db:"consent_form_complete"`

	// consent id
	ConsentID *string `json:"consent_id,omitempty" db:"consent_id"`

	// consent version
	ConsentVersion *string `json:"consent_version,omitempty" db:"consent_version"`

	// consenting coordinator name
	ConsentingCoordinatorName *string `json:"consenting_coordinator_name,omitempty" db:"consenting_coordinator_name"`

	// date of assent
	DateOfAssent *string `json:"date_of_assent,omitempty" db:"date_of_assent"`

	// date of consent withdrawal
	DateOfConsentWithdrawal *string `json:"date_of_consent_withdrawal,omitempty" db:"date_of_consent_withdrawal"`

	// has consent been withdrawn
	HasConsentBeenWithdrawn *string `json:"has_consent_been_withdrawn,omitempty" db:"has_consent_been_withdrawn"`

	// if assent not obtained why not
	IfAssentNotObtainedWhyNot *string `json:"if_assent_not_obtained_why_not,omitempty" db:"if_assent_not_obtained_why_not"`

	// name of other biobank
	NameOfOtherBiobank *string `json:"name_of_other_biobank,omitempty" db:"name_of_other_biobank"`

	// patient consented to
	PatientConsentedTo *string `json:"patient_consented_to,omitempty" db:"patient_consented_to"`

	// sample id
	SampleID *string `json:"sample_id,omitempty" db:"sample_id"`

	// previously consented
	PreviouslyConsented *string `json:"previously_consented,omitempty" db:"previously_consented"`

	// reason for consent withdrawal
	ReasonForConsentWithdrawal *string `json:"reason_for_consent_withdrawal,omitempty" db:"reason_for_consent_withdrawal"`

	// reason for rejection
	ReasonForRejection *string `json:"reason_for_rejection,omitempty" db:"reason_for_rejection"`

	// reconsent date
	ReconsentDate *string `json:"reconsent_date,omitempty" db:"reconsent_date"`

	// reconsent version
	ReconsentVersion *string `json:"reconsent_version,omitempty" db:"reconsent_version"`

	// type of consent withdrawal
	TypeOfConsentWithdrawal *string `json:"type_of_consent_withdrawal,omitempty" db:"type_of_consent_withdrawal"`

	// was assent obtained
	WasAssentObtained *string `json:"was_assent_obtained,omitempty" db:"was_assent_obtained"`

	// hash induced neoplasm details
	Hash *string `json:"hash,omitempty" db:"hash"`
}

// Validate validates this consent
func (m *Consent) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *Consent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Consent) UnmarshalBinary(b []byte) error {
	var res Consent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
