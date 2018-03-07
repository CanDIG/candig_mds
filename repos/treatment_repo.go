package repos

import (
	"log"

	"github.com/CanDIG/candig_mds/database"
	"github.com/CanDIG/candig_mds/models"
	"gopkg.in/mgo.v2/bson"
)

//GetAllTreatments returns an array of all objects in a collection
func GetAllTreatments(collection string) []*models.Treatment {
	c := database.SetCollection(collection)
	list := make([]*models.Treatment, 0)
	err := c.Find(nil).All(&list)
	if err != nil {
		log.Fatal(err)
	}
	return list
}

//CheckIfTreatmentExists returns an object in a collection
func CheckIfTreatmentExists(hash string) bool {
	c := database.SetCollection("sample")
	samples := make([]*models.Sample, 0)
	err := c.Find(bson.M{"treatment.hash": hash}).All(&samples)
	if err != nil {
		log.Fatal(err)
	}
	if len(samples) != 0 {
		return true
	}
	return false
}

//InsertTreatment inserts a treatment to the treatment collection
func InsertTreatment(treatment models.Treatment) bool {
	return database.Insert("treatment", treatment)
}
