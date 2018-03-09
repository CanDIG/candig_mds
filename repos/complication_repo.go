package repos

import (
	"encoding/json"
	"log"

	"github.com/CanDIG/candig_mds/database"
	"github.com/CanDIG/candig_mds/models"
	errors "github.com/go-openapi/errors"
	"gopkg.in/mgo.v2/bson"
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

//CheckIfComplicationExists returns an object in a collection
func CheckIfComplicationExists(hash string) bool {
	c := database.SetCollection("sample")
	samples := make([]*models.Sample, 0)
	err := c.Find(bson.M{"complications.hash": hash}).All(&samples)
	if err != nil {
		log.Fatal(err)
	}
	if len(samples) != 0 {
		return true
	}
	return false
}

//InsertComplication inserts a patient to the patient collection
func InsertComplication(complication models.Complication) bool {
	return database.Insert("complication", complication)
}

//AddComplication 4 comp
func AddComplication(complication *models.Complication) error {
	if complication == nil {
		return errors.New(500, "item must be present")
	}
	jdata, _ := json.Marshal(complication)
	jstring := database.Hash(string(jdata))
	complication.Hash = &jstring
	if CheckIfComplicationExists(*complication.Hash) {
		return errors.New(500, "Already Exists")
	}
	sample := GetSampleByID(*complication.SampleID)
	if sample == nil {
		return errors.New(500, "Sample does not exist")
	}
	sample.Complications = append(sample.Complications, complication)
	UpdateSample(*sample)
	return nil
}
