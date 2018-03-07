package repos

import (
	"log"

	"github.com/CanDIG/candig_mds/database"
	"github.com/CanDIG/candig_mds/models"
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

//InsertTumourboard inserts a tumourboard to the tumourboard collection
func InsertTumourboard(tumourboard models.Tumourboard) bool {
	return database.Insert("tumourboard", tumourboard)
}
