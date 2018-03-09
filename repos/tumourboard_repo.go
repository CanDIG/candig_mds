package repos

import (
	"encoding/json"
	"log"

	"github.com/CanDIG/candig_mds/database"
	"github.com/CanDIG/candig_mds/models"
	errors "github.com/go-openapi/errors"
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

//AddTumourboard 4 tum
func AddTumourboard(tumourboard *models.Tumourboard) error {
	if tumourboard == nil {
		return errors.New(500, "item must be present")
	}
	jdata, _ := json.Marshal(tumourboard)
	jstring := database.Hash(string(jdata))
	tumourboard.Hash = &jstring
	if CheckIfTumourboardExists(*tumourboard.Hash) {
		return errors.New(500, "Already Exists")
	}
	sample := GetSampleByID(*tumourboard.SampleID)
	if sample == nil {
		return errors.New(500, "Sample does not exist")
	}
	sample.Tumourboards = append(sample.Tumourboards, tumourboard)
	UpdateSample(*sample)
	return nil
}
