package repos

import (
	"log"

	"github.com/CanDIG/candig_mds/database"
	"github.com/CanDIG/candig_mds/models"
)

//GetAllComplications returns an array of all objects in a collection
func GetAllComplications(collection string) []*models.Complication {
	c := database.SetCollection(collection)
	list := make([]*models.Complication, 0)
	err := c.Find(nil).All(&list)
	if err != nil {
		log.Fatal(err)
	}
	return list
}

//InsertComplication inserts a patient to the patient collection
func InsertComplication(complication models.Complication) bool {
	return database.Insert("complication", complication)
}
