package repos

import (
	"log"

	"github.com/CanDIG/candig_mds/database"
	"github.com/CanDIG/candig_mds/models"
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

//InsertTreatment inserts a treatment to the treatment collection
func InsertTreatment(treatment models.Treatment) bool {
	return database.Insert("treatment", treatment)
}
