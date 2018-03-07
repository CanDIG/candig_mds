package repos

import (
	"log"

	"github.com/CanDIG/candig_mds/database"
	"github.com/CanDIG/candig_mds/models"
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

//InsertConsent inserts a patient to the patient collection
func InsertConsent(consent models.Consent) bool {
	return database.Insert("consent", consent)
}
