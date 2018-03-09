package repos

import (
	"encoding/json"
	"log"

	"github.com/CanDIG/candig_mds/database"
	"github.com/CanDIG/candig_mds/models"
	errors "github.com/go-openapi/errors"
	"gopkg.in/mgo.v2/bson"
)

//GetAllConsents returns an array of all objects in a collection
func GetAllConsents(collection string) []*models.Consent {
	c := database.SetCollection(collection)
	list := make([]*models.Consent, 0)
	err := c.Find(nil).All(&list)
	if err != nil {
		log.Fatal(err)
	}
	return list
}

//CheckIfConsentExists returns an object in a collection
func CheckIfConsentExists(hash string) bool {
	c := database.SetCollection("sample")
	samples := make([]*models.Sample, 0)
	err := c.Find(bson.M{"consents.hash": hash}).All(&samples)
	if err != nil {
		log.Fatal(err)
	}
	if len(samples) != 0 {
		return true
	}
	return false
}

//InsertConsent inserts a patient to the patient collection
func InsertConsent(consent models.Consent) bool {
	return database.Insert("consent", consent)
}

//AddConsent for sample
func AddConsent(consent *models.Consent) error {
	if consent == nil {
		return errors.New(500, "item must be present")
	}
	jdata, _ := json.Marshal(consent)
	jstring := database.Hash(string(jdata))
	consent.Hash = &jstring
	if CheckIfConsentExists(*consent.Hash) {
		return errors.New(500, "Already Exists")
	}
	sample := GetSampleByID(*consent.SampleID)
	if sample == nil {
		return errors.New(500, "Sample does not exist")
	}
	sample.Consents = append(sample.Consents, consent)
	UpdateSample(*sample)
	return nil
}
