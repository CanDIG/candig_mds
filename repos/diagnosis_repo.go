package repos

import (
	"log"

	"github.com/CanDIG/candig_mds/database"
	"github.com/CanDIG/candig_mds/models"
	"gopkg.in/mgo.v2/bson"
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

//CheckIfDiagnosisExists returns an object in a collection
func CheckIfDiagnosisExists(hash string) bool {
	c := database.SetCollection("sample")
	samples := make([]*models.Sample, 0)
	err := c.Find(bson.M{"diagnosis.hash": hash}).All(&samples)
	if err != nil {
		log.Fatal(err)
	}
	if len(samples) != 0 {
		return true
	}
	return false
}

//InsertDiagnosis inserts a patient to the patient collection
func InsertDiagnosis(diagnosis models.Diagnosis) bool {
	return database.Insert("diagnosis", diagnosis)
}
