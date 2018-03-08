package repos

import (
	"log"

	"github.com/CanDIG/candig_mds/database"
	"github.com/CanDIG/candig_mds/models"
	"gopkg.in/mgo.v2/bson"
)

//GetAllEnrollments returns an array of all objects in a collection
func GetAllEnrollments(collection string) []*models.Enrollment {
	c := database.SetCollection(collection)
	list := make([]*models.Enrollment, 0)
	err := c.Find(nil).All(&list)
	if err != nil {
		log.Fatal(err)
	}
	return list
}

//CheckIfEnrollmentExists returns an object in a collection
func CheckIfEnrollmentExists(hash string) bool {
	c := database.SetCollection("sample")
	samples := make([]*models.Sample, 0)
	err := c.Find(bson.M{"enrollments.hash": hash}).All(&samples)
	if err != nil {
		log.Fatal(err)
	}
	if len(samples) != 0 {
		return true
	}
	return false
}

//InsertEnrollment inserts a patient to the patient collection
func InsertEnrollment(enrollment models.Enrollment) bool {
	return database.Insert("enrollment", enrollment)
}
