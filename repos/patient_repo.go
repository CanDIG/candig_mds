package repos

import (
	"log"

	"github.com/CanDIG/candig_mds/database"
	"github.com/CanDIG/candig_mds/models"
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

//InsertPatient inserts a patient to the patient collection
func InsertPatient(patient models.Patient) bool {
	return database.Insert("patient", patient)
}
