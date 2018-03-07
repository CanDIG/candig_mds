package repos

import (
	"log"

	"github.com/CanDIG/candig_mds/database"
	"github.com/CanDIG/candig_mds/models"
)

//GetAllOutcomes returns an array of all objects in a collection
func GetAllOutcomes(collection string) []*models.Outcome {
	c := database.SetCollection(collection)
	list := make([]*models.Outcome, 0)
	err := c.Find(nil).All(&list)
	if err != nil {
		log.Fatal(err)
	}
	return list
}

//InsertOutcome inserts a outcome to the outcome collection
func InsertOutcome(outcome models.Outcome) bool {
	return database.Insert("outcome", outcome)
}
