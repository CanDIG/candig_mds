// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Treatment treatment
// swagger:model Treatment
type Treatment struct {

	// chemotherapy details
	ChemotherapyDetails *string `json:"chemotherapy_details,omitempty" db:"chemotherapy_details"`

	// course number
	CourseNumber *string `json:"course_number,omitempty" db:"course_number"`

	// date of recurrence or progression after this treatment
	DateOfRecurrenceOrProgressionAfterThisTreatment *string `json:"date_of_recurrence_or_progression_after_this_treatment,omitempty" db:"date_of_recurrence_or_progression_after_this_treatment"`

	// drug id numbers
	DrugIDNumbers *string `json:"drug_id_numbers,omitempty" db:"drug_id_numbers"`

	// drug list or agent
	DrugListOrAgent *string `json:"drug_list_or_agent,omitempty" db:"drug_list_or_agent"`

	// hematopoietic cell transplant
	HematopoieticCellTransplant *string `json:"hematopoietic_cell_transplant,omitempty" db:"hematopoietic_cell_transplant"`

	// immunotherapy details
	ImmunotherapyDetails *string `json:"immunotherapy_details,omitempty" db:"immunotherapy_details"`

	// sample id
	SampleID *string `json:"sample_id,omitempty" db:"sample_id"`

	// protocol number or code
	ProtocolNumberOrCode *string `json:"protocol_number_or_code,omitempty" db:"protocol_number_or_code"`

	// radiotherapy details
	RadiotherapyDetails *string `json:"radiotherapy_details,omitempty" db:"radiotherapy_details"`

	// reason for ending the treatment
	ReasonForEndingTheTreatment *string `json:"reason_for_ending_the_treatment,omitempty" db:"reason_for_ending_the_treatment"`

	// response criteria used
	ResponseCriteriaUsed *string `json:"response_criteria_used,omitempty" db:"response_criteria_used"`

	// response to treatment
	ResponseToTreatment *string `json:"response_to_treatment,omitempty" db:"response_to_treatment"`

	// start date
	StartDate *string `json:"start_date,omitempty" db:"start_date"`

	// stop date
	StopDate *string `json:"stop_date,omitempty" db:"stop_date"`

	// surgery details
	SurgeryDetails *string `json:"surgery_details,omitempty" db:"surgery_details"`

	// systematic therapy agent name
	SystematicTherapyAgentName *string `json:"systematic_therapy_agent_name,omitempty" db:"systematic_therapy_agent_name"`

	// therapeutic modality
	TherapeuticModality *string `json:"therapeutic_modality,omitempty" db:"therapeutic_modality"`

	// treatment intent
	TreatmentIntent *string `json:"treatment_intent,omitempty" db:"treatment_intent"`

	// treatment plan type
	TreatmentPlanType *string `json:"treatment_plan_type,omitempty" db:"treatment_plan_type"`

	// unexpected or unusual toxicity during treatment
	UnexpectedOrUnusualToxicityDuringTreatment *string `json:"unexpected_or_unusual_toxicity_during_treatment,omitempty" db:"unexpected_or_unusual_toxicity_during_treatment"`

	// hash induced neoplasm details
	Hash *string `json:"hash,omitempty" db:"hash"`
}

// Validate validates this treatment
func (m *Treatment) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *Treatment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Treatment) UnmarshalBinary(b []byte) error {
	var res Treatment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
