package repos

import (
	"log"

	"github.com/CanDIG/candig_mds/database"
	"github.com/CanDIG/candig_mds/models"
	errors "github.com/go-openapi/errors"
	"gopkg.in/mgo.v2/bson"
)

//GetAllPatients returns an array of all objects in a collection
func GetAllPatients(collection string) []*models.Patient {
	c := database.SetCollection(collection)
	list := make([]*models.Patient, 0)
	err := c.Find(nil).All(&list)
	if err != nil {
		log.Fatal(err)
	}
	return list
}

//CheckIfPatientExists returns an object in a collection
func CheckIfPatientExists(patientID string) bool {
	c := database.SetCollection("sample")
	samples := make([]*models.Sample, 0)
	err := c.Find(bson.M{"patients.patientid": patientID}).All(&samples)
	if err != nil {
		log.Fatal(err)
	}
	if len(samples) != 0 {
		return true
	}
	return false
}

//InsertPatient inserts a patient to the patient collection
func InsertPatient(patient models.Patient) bool {
	return database.Insert("patient", patient)
}

//AddPatient is for patients
func AddPatient(patient *models.Patient) error {
	if patient == nil {
		return errors.New(500, "item must be present")
	}
	if CheckIfPatientExists(*patient.PatientID) {
		return errors.New(500, "Already Exists")
	}
	sample := GetSampleByID(*patient.SampleID)
	if sample == nil {
		return errors.New(500, "Sample does not exist")
	}
	sample.Patient = patient
	UpdateSample(*sample)
	return nil
}
