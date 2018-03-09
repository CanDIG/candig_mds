package dummydata

import (
	"math/rand"

	"github.com/CanDIG/candig_mds/repos"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var r *rand.Rand
var sampleID int

//GenerateData creates n fake records
func GenerateData(number int) {
	generateSamples(number)
	generatePatients(number)
	generateComplications(number)
	generateConsents(number)
	generateDiagnosis(number)
	generateEnrollments(number)
	generateOutcomes(number)
	generateTreatments(number)
	generateTumourboards(number)
}

func generateSamples(number int) {
	sampleID = 0
	for i := 0; i < number; i++ {
		sampleID++
		tempSample := makeSample()
		repos.InsertSample(tempSample)
	}
}

func generatePatients(number int) {
	sampleID = 0
	for i := 0; i < number; i++ {
		sampleID++
		tempPatient := makePatient()
		repos.InsertPatient(tempPatient)
	}
}

func generateComplications(number int) {
	sampleID = 0
	for i := 0; i < number; i++ {
		sampleID++
		tempComplication := makeComplication()
		repos.AddComplication(&tempComplication)
	}
}

func generateConsents(number int) {
	sampleID = 0
	for i := 0; i < number; i++ {
		sampleID++
		tempConsent := makeConsent()
		repos.AddConsent(&tempConsent)
	}
}

func generateDiagnosis(number int) {
	sampleID = 0
	for i := 0; i < number; i++ {
		sampleID++
		tempDiagnosis := makeDiagnosis()
		repos.AddDiagnosis(&tempDiagnosis)
	}
}

func generateEnrollments(number int) {
	sampleID = 0
	for i := 0; i < number; i++ {
		sampleID++
		tempEnrollent := makeEnrollment()
		repos.AddEnrollment(&tempEnrollent)
	}
}

func generateOutcomes(number int) {
	sampleID = 0
	for i := 0; i < number; i++ {
		sampleID++
		tempOutcome := makeOutcome()
		repos.AddOutcome(&tempOutcome)
	}
}

func generateTreatments(number int) {
	sampleID = 0
	for i := 0; i < number; i++ {
		sampleID++
		tempTreatment := makeTreatment()
		repos.AddTreatment(&tempTreatment)
	}
}

func generateTumourboards(number int) {
	sampleID = 0
	for i := 0; i < number; i++ {
		sampleID++
		tempTumourboard := makeTumourboard()
		repos.AddTumourboard(&tempTumourboard)
	}
}
