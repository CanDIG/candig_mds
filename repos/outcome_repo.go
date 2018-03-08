package repos

import (
	"log"

	"github.com/CanDIG/candig_mds/database"
	"github.com/CanDIG/candig_mds/models"
	"gopkg.in/mgo.v2/bson"
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

//CheckIfOutcomeExists returns an object in a collection
func CheckIfOutcomeExists(hash string) bool {
	c := database.SetCollection("sample")
	samples := make([]*models.Sample, 0)
	err := c.Find(bson.M{"outcomes.hash": hash}).All(&samples)
	if err != nil {
		log.Fatal(err)
	}
	if len(samples) != 0 {
		return true
	}
	return false
}

//InsertOutcome inserts a outcome to the outcome collection
func InsertOutcome(outcome models.Outcome) bool {
	return database.Insert("outcome", outcome)
}
