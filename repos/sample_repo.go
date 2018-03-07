package repos

import (
	"log"

	"github.com/CanDIG/candig_mds/database"
	"github.com/CanDIG/candig_mds/models"
	"gopkg.in/mgo.v2/bson"
)

//GetAllSamples returns an array of all objects in a collection
func GetAllSamples(collection string) []*models.Sample {
	c := database.SetCollection(collection)
	list := make([]*models.Sample, 0)
	err := c.Find(nil).All(&list)
	if err != nil {
		log.Fatal(err)
	}
	return list
}

//CheckIfSampleExists returns an object in a collection
func CheckIfSampleExists(sampleID string) bool {
	c := database.SetCollection("sample")
	samples := make([]*models.Sample, 0)
	err := c.Find(bson.M{"sampleid": sampleID}).All(&samples)
	if err != nil {
		log.Fatal(err)
	}
	if len(samples) != 0 {
		return true
	}
	return false
}

//GetSampleByID gets a sample by its id
func GetSampleByID(sampleID string) *models.Sample {
	c := database.SetCollection("sample")
	var sample *models.Sample
	err := c.Find(bson.M{"sampleid": sampleID}).One(&sample)
	if err != nil {
		log.Printf("%v", err)
		return nil
	}
	return sample
}

//InsertSample inserts a patient to the patient collection
func InsertSample(sample models.Sample) bool {
	return database.Insert("sample", sample)
}

//UpdateSample updates a sample
func UpdateSample(sample models.Sample) bool {
	return database.Update("sample", bson.M{"sampleid": sample.SampleID}, sample)
}
