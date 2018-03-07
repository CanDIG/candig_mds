package repos

import (
	"log"

	"github.com/CanDIG/candig_mds/database"
	"github.com/CanDIG/candig_mds/models"
)

//GetAllDiagnosis returns an array of all objects in a collection
func GetAllDiagnosis(collection string) []*models.Diagnosis {
	c := database.SetCollection(collection)
	list := make([]*models.Diagnosis, 0)
	err := c.Find(nil).All(&list)
	if err != nil {
		log.Fatal(err)
	}
	return list
}

//InsertDiagnosis inserts a patient to the patient collection
func InsertDiagnosis(diagnosis models.Diagnosis) bool {
	return database.Insert("diagnosis", diagnosis)
}
