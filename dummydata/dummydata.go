package dummydata

import (
	"math/rand"
	"strconv"

	models "github.com/CanDIG/candig_mds/models"
	repos "github.com/CanDIG/candig_mds/repos"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var r *rand.Rand
var id1 int
var id2 int
var id3 int
var id4 int
var id5 int

//MakeData makes dummy data and puts it into the db
func MakeData(numberPatients, numberSamples int) {
	r = rand.New(rand.NewSource(99))
	createPatients(numberPatients)
	createSamples(numberSamples)
	createExperiments(numberSamples)
	createResults(numberSamples)
	createResultDetails(numberSamples)
}

func makeRandomString() string {
	num := rand.Intn(50)
	value := randSeq(num)
	return value
}

func makeRandomDate() string {
	date := strconv.Itoa(rand.Intn(118)+1900) + "-" + strconv.Itoa(rand.Intn(11)+1) + "-" + strconv.Itoa(rand.Intn(27)+1)
	return date
}

func makeRandomFloat() float32 {
	num := rand.Float32()
	num += float32(rand.Intn(600))
	return num
}

func makeRandomBool() bool {
	num := rand.Intn(1)
	if num == 1 {
		return true
	}
	return false
}

func createPatients(number int) {
	for i := 0; i < number; i++ {
		id1++
		tempPatient := makePatient()
		repos.InsertPatient(&tempPatient)
	}
}
func createResults(number int) {
	for i := 0; i < number; i++ {
		id4++
		tempResult := makeResult()
		repos.InsertResult(&tempResult)
	}
}
func createResultDetails(number int) {
	for i := 0; i < number; i++ {
		id5++
		tempResultDetail := makeResultDetail()
		repos.InsertResultDetail(&tempResultDetail)
	}
}

func createSamples(number int) {
	for i := 0; i < number; i++ {
		id2++
		tempSample := makeSample()
		repos.InsertSample(&tempSample)
	}
}
func createExperiments(number int) {
	for i := 0; i < number; i++ {
		id3++
		tempExperiment := makeExperiment()
		repos.InsertExperiment(&tempExperiment)
	}
}

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func makePatient() models.Patient {
	patient := models.Patient{}
	ClinicalHistory := makeRandomString()
	patient.ClinicalHistory = &ClinicalHistory
	DateReceived := makeRandomDate()
	patient.DateReceived = &DateReceived
	DateReported := makeRandomDate()
	patient.DateReported = &DateReported
	Dob := makeRandomDate()
	patient.Dob = &Dob
	FirstName := makeRandomString()
	patient.FirstName = &FirstName
	Gender := makeRandomString()
	patient.Gender = &Gender
	Initials := makeRandomString()
	patient.Initials = &Initials
	LastName := makeRandomString()
	patient.LastName = &LastName
	Mrn := makeRandomString()
	patient.Mrn = &Mrn
	OnHcn := makeRandomString()
	patient.OnHcn = &OnHcn
	PatientID := makeRandomString()
	patient.PatientID = &PatientID
	PatientType := makeRandomString()
	patient.PatientType = &PatientType
	ReferringPhysican := makeRandomString()
	patient.ReferringPhysican = &ReferringPhysican
	SampleID := strconv.Itoa(id1)
	patient.SampleID = &SampleID
	SeNum := makeRandomString()
	patient.SeNum = &SeNum
	SurgicalDate := makeRandomDate()
	patient.SurgicalDate = &SurgicalDate

	return patient
}

func makeSample() models.Sample {
	sample := models.Sample{}
	SampleID := strconv.Itoa(id2)
	sample.SampleID = &SampleID
	Facility := makeRandomString()
	sample.Facility = &Facility
	TestRequested := makeRandomString()
	sample.TestRequested = &TestRequested
	SeNum := makeRandomString()
	sample.SeNum = &SeNum
	DateCollected := makeRandomDate()
	sample.DateCollected = &DateCollected
	DateReceived := makeRandomDate()
	sample.DateReceived = &DateReceived
	SampleType := makeRandomString()
	sample.SampleType = &SampleType
	MaterialReceived := makeRandomString()
	sample.MaterialReceived = &MaterialReceived
	MaterialReceivedNum := makeRandomString()
	sample.MaterialReceivedNum = &MaterialReceivedNum
	MaterialReceivedOther := makeRandomString()
	sample.MaterialReceivedOther = &MaterialReceivedOther
	VolumeOfBloodMarrow := makeRandomFloat()
	sample.VolumeOfBloodMarrow = &VolumeOfBloodMarrow
	SurgicalNum := makeRandomString()
	sample.SurgicalNum = &SurgicalNum
	TumorSite := makeRandomString()
	sample.TumorSite = &TumorSite
	HistoricalDiagnosis := makeRandomString()
	sample.HistoricalDiagnosis = &HistoricalDiagnosis
	TumorPercntOfTotal := makeRandomFloat()
	sample.TumorPercntOfTotal = &TumorPercntOfTotal
	TumorPercntOfCircled := makeRandomFloat()
	sample.TumorPercntOfCircled = &TumorPercntOfCircled
	ReviewedBy := makeRandomString()
	sample.ReviewedBy = &ReviewedBy
	HESlideLocation := makeRandomString()
	sample.HESlideLocation = &HESlideLocation
	NonUhnID := makeRandomString()
	sample.NonUhnID = &NonUhnID
	NameOfRequestor := makeRandomString()
	sample.NameOfRequestor = &NameOfRequestor
	DnaConcentration := makeRandomFloat()
	sample.DnaConcentration = &DnaConcentration
	DnaVolume := makeRandomFloat()
	sample.DnaVolume = &DnaVolume
	DnaLocation := makeRandomString()
	sample.DnaLocation = &DnaLocation
	RnaConcentration := makeRandomFloat()
	sample.RnaConcentration = &RnaConcentration
	RnaVolume := makeRandomFloat()
	sample.RnaVolume = &RnaVolume
	RnaLocation := makeRandomString()
	sample.RnaLocation = &RnaLocation
	WbcLocation := makeRandomString()
	sample.WbcLocation = &WbcLocation
	PlasmaLocation := makeRandomString()
	sample.PlasmaLocation = &PlasmaLocation
	CfPlasmaLocation := makeRandomString()
	sample.CfPlasmaLocation = &CfPlasmaLocation
	PbBmLocation := makeRandomString()
	sample.PbBmLocation = &PbBmLocation
	RnaLysateLocation := makeRandomString()
	sample.RnaLysateLocation = &RnaLysateLocation
	SampleSize := makeRandomString()
	sample.SampleSize = &SampleSize
	StudyID := makeRandomString()
	sample.StudyID = &StudyID
	SampleName := makeRandomString()
	sample.SampleName = &SampleName
	DateSubmitted := makeRandomDate()
	sample.DateSubmitted = &DateSubmitted
	ContainerType := makeRandomString()
	sample.ContainerType = &ContainerType
	ContainerID := makeRandomString()
	sample.ContainerID = &ContainerID
	ContainerWell := makeRandomString()
	sample.ContainerWell = &ContainerWell
	CopathNum := makeRandomString()
	sample.CopathNum = &CopathNum
	OtherIdentifier := makeRandomString()
	sample.OtherIdentifier = &OtherIdentifier
	HasSampleFiles := makeRandomBool()
	sample.HasSampleFiles = &HasSampleFiles
	DnaSampleBarcode := makeRandomString()
	sample.DnaSampleBarcode = &DnaSampleBarcode
	DnaExtractionDate := makeRandomDate()
	sample.DnaExtractionDate = &DnaExtractionDate
	DnaQuality := makeRandomString()
	sample.DnaQuality = &DnaQuality
	FfpeQcDate := makeRandomDate()
	sample.FfpeQcDate = &FfpeQcDate
	DeltaCtValue := makeRandomFloat()
	sample.DeltaCtValue = &DeltaCtValue
	Comments := makeRandomString()
	sample.Comments = &Comments
	RnasePDate := makeRandomDate()
	sample.RnasePDate = &RnasePDate
	DnaQualityByRnaseP := makeRandomFloat()
	sample.DnaQualityByRnaseP = &DnaQualityByRnaseP
	RnaQuality := makeRandomFloat()
	sample.RnaQuality = &RnaQuality
	RnaExtractionDate := makeRandomDate()
	sample.RnaExtractionDate = &RnaExtractionDate

	return sample
}

func makeExperiment() models.Experiment {
	experiment := models.Experiment{}
	ChipCartridgeBarcode := makeRandomString()
	experiment.ChipCartridgeBarcode = &ChipCartridgeBarcode
	CompleteDate := makeRandomDate()
	experiment.CompleteDate = &CompleteDate
	ExperimentID := makeRandomString()
	experiment.ExperimentID = &ExperimentID
	HasProjectFiles := makeRandomBool()
	experiment.HasProjectFiles = &HasProjectFiles
	OpenedDate := makeRandomDate()
	experiment.OpenedDate = &OpenedDate
	PanelAssayScreened := makeRandomString()
	experiment.PanelAssayScreened = &PanelAssayScreened
	Pcr := makeRandomString()
	experiment.Pcr = &Pcr
	Priority := makeRandomString()
	experiment.Priority = &Priority
	ProcedureOrderDatetime := makeRandomDate()
	experiment.ProcedureOrderDatetime = &ProcedureOrderDatetime
	ProjectID := makeRandomString()
	experiment.ProjectID = &ProjectID
	ProjectName := makeRandomString()
	experiment.ProjectName = &ProjectName
	SampleID := strconv.Itoa(id3)
	experiment.SampleID = &SampleID
	StudyID := makeRandomString()
	experiment.StudyID = &StudyID
	TestDate := makeRandomDate()
	experiment.TestDate = &TestDate

	return experiment
}

func makeResult() models.Result {
	result := models.Result{}
	FailedRegions := makeRandomString()
	result.FailedRegions = &FailedRegions
	MeanDepthOfCoveage := makeRandomFloat()
	result.MeanDepthOfCoveage = &MeanDepthOfCoveage
	MlpaPcr := makeRandomString()
	result.MlpaPcr = &MlpaPcr
	Mutation := makeRandomString()
	result.Mutation = &Mutation
	OverallHotspotsThreshold := makeRandomFloat()
	result.OverallHotspotsThreshold = &OverallHotspotsThreshold
	OverallQualityThreshold := makeRandomFloat()
	result.OverallQualityThreshold = &OverallQualityThreshold
	ResultsID := makeRandomString()
	result.ResultsID = &ResultsID
	SampleID := strconv.Itoa(id4)
	result.SampleID = &SampleID
	UID := makeRandomString()
	result.UID = &UID
	VerificationPcr := makeRandomString()
	result.VerificationPcr = &VerificationPcr

	return result
}

func makeResultDetail() models.Resultdetails {
	resultdetail := models.Resultdetails{}
	CNomenclature := makeRandomString()
	resultdetail.CNomenclature = &CNomenclature
	Coverage := int64(rand.Intn(1000))
	resultdetail.Coverage = &Coverage
	Exon := int64(rand.Intn(1000))
	resultdetail.Exon = &Exon
	Gene := makeRandomString()
	resultdetail.Gene = &Gene
	Pcr := makeRandomString()
	resultdetail.Pcr = &Pcr
	PNomenclature := makeRandomString()
	resultdetail.PNomenclature = &PNomenclature
	QualityScore := makeRandomFloat()
	resultdetail.QualityScore = &QualityScore
	Result := makeRandomString()
	resultdetail.Result = &Result
	ResultsDetailsID := makeRandomString()
	resultdetail.ResultsDetailsID = &ResultsDetailsID
	ResultsID := makeRandomString()
	resultdetail.ResultsID = &ResultsID
	RiskScore := makeRandomFloat()
	resultdetail.RiskScore = &RiskScore
	SampleID := strconv.Itoa(id5)
	resultdetail.SampleID = &SampleID
	UID := makeRandomString()
	resultdetail.UID = &UID
	VAF := makeRandomFloat()
	resultdetail.VAF = &VAF

	return resultdetail
}
