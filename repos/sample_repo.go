package repos

import (
	"log"

	"github.com/CanDIG/candig_mds/database"
	"github.com/CanDIG/candig_mds/models"
)

//GetAllSamples returns an array of all objects in a collection
func GetAllSamples(collection string) []*models.Sample {
	c := database.SetCollection(collection)
	list := make([]*models.Sample, 0)
	err := c.Find(nil).All(&list)
	if err != nil {
		log.Fatal(err)
	}
	return list
}

//InsertSample inserts a patient to the patient collection
func InsertSample(sample models.Sample) bool {
	return database.Insert("sample", sample)
}
