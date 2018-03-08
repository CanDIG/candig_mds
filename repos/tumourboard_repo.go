package repos

import (
	"log"

	"github.com/CanDIG/candig_mds/database"
	"github.com/CanDIG/candig_mds/models"
	"gopkg.in/mgo.v2/bson"
)

//GetAllTumourboards returns an array of all objects in a collection
func GetAllTumourboards(collection string) []*models.Tumourboard {
	c := database.SetCollection(collection)
	list := make([]*models.Tumourboard, 0)
	err := c.Find(nil).All(&list)
	if err != nil {
		log.Fatal(err)
	}
	return list
}

//CheckIfTumourboardExists returns an object in a collection
func CheckIfTumourboardExists(hash string) bool {
	c := database.SetCollection("sample")
	samples := make([]*models.Sample, 0)
	err := c.Find(bson.M{"tumourboards.hash": hash}).All(&samples)
	if err != nil {
		log.Fatal(err)
	}
	if len(samples) != 0 {
		return true
	}
	return false
}

//InsertTumourboard inserts a tumourboard to the tumourboard collection
func InsertTumourboard(tumourboard models.Tumourboard) bool {
	return database.Insert("tumourboard", tumourboard)
}
