package repos

import (
	"log"

	"github.com/CanDIG/candig_mds/database"
	"github.com/CanDIG/candig_mds/models"
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

//InsertEnrollment inserts a patient to the patient collection
func InsertEnrollment(enrollment models.Enrollment) bool {
	return database.Insert("enrollment", enrollment)
}
