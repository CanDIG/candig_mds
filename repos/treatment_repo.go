package repos

import (
	"encoding/json"
	"log"

	"github.com/CanDIG/candig_mds/database"
	"github.com/CanDIG/candig_mds/models"
	errors "github.com/go-openapi/errors"
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
	err := c.Find(bson.M{"treatments.hash": hash}).All(&samples)
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

//AddTreatment for treatments
func AddTreatment(treatment *models.Treatment) error {
	if treatment == nil {
		return errors.New(500, "item must be present")
	}
	jdata, _ := json.Marshal(treatment)
	jstring := database.Hash(string(jdata))
	treatment.Hash = &jstring
	if CheckIfTreatmentExists(*treatment.Hash) {
		return errors.New(500, "Already Exists")
	}
	sample := GetSampleByID(*treatment.SampleID)
	if sample == nil {
		return errors.New(500, "Sample does not exist")
	}
	sample.Treatments = append(sample.Treatments, treatment)
	UpdateSample(*sample)
	return nil
}
