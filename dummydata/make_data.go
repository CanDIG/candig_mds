package dummydata

import (
	"math/rand"

	"github.com/CanDIG/candig_mds/models"
)

func makeSample() models.Sample {
	sample := models.Sample{}
	AnatomicSiteTheSampleObtainedFrom := makeRandomString()
	sample.AnatomicSiteTheSampleObtainedFrom = &AnatomicSiteTheSampleObtainedFrom
	AssociatedBiobank := makeRandomString()
	sample.AssociatedBiobank = &AssociatedBiobank
	CancerSubtype := makeRandomString()
	sample.CancerSubtype = &CancerSubtype
	CancerType := makeRandomString()
	sample.CancerType = &CancerType
	CollectionDate := makeRandomDate()
	sample.CollectionDate = &CollectionDate
	CollectionHospital := makeRandomString()
	sample.CollectionHospital = &CollectionHospital
	DiagnosisID := makeRandomString()
	sample.DiagnosisID = &DiagnosisID
	EstimatedTumourContent := makeRandomString()
	sample.EstimatedTumourContent = &EstimatedTumourContent
	IfNotExplainAnyDeviation := makeRandomString()
	sample.IfNotExplainAnyDeviation = &IfNotExplainAnyDeviation
	LocalBiobankID := makeRandomString()
	sample.LocalBiobankID = &LocalBiobankID
	MorphologicalCode := makeRandomString()
	sample.MorphologicalCode = &MorphologicalCode
	OtherBiobank := makeRandomString()
	sample.OtherBiobank = &OtherBiobank
	PathologyReportID := makeRandomString()
	sample.PathologyReportID = &PathologyReportID
	QualityControlPerformed := makeRandomString()
	sample.QualityControlPerformed = &QualityControlPerformed
	Quantity := makeRandomString()
	sample.Quantity = &Quantity
	ReceivedDate := makeRandomDate()
	sample.ReceivedDate = &ReceivedDate
	SampleID := string(sampleID)
	sample.SampleID = &SampleID
	SampleType := makeRandomString()
	sample.SampleType = &SampleType
	ShippingDate := makeRandomDate()
	sample.ShippingDate = &ShippingDate
	SopFollowed := makeRandomString()
	sample.SopFollowed = &SopFollowed
	TissueDiseaseState := makeRandomString()
	sample.TissueDiseaseState = &TissueDiseaseState
	TopologicalCode := makeRandomString()
	sample.TopologicalCode = &TopologicalCode
	Units := makeRandomString()
	sample.Units = &Units

	return sample
}

func makePatient() models.Patient {
	patient := models.Patient{}
	AutopsyTissueForResearch := makeRandomString()
	patient.AutopsyTissueForResearch = &AutopsyTissueForResearch
	CauseOfDeath := makeRandomString()
	patient.CauseOfDeath = &CauseOfDeath
	DateOfBirth := makeRandomDate()
	patient.DateOfBirth = &DateOfBirth
	DateOfDeath := makeRandomDate()
	patient.DateOfDeath = &DateOfDeath
	DetailsOfPredispositionSyndrome := makeRandomString()
	patient.DetailsOfPredispositionSyndrome = &DetailsOfPredispositionSyndrome
	Ethnicity := makeRandomString()
	patient.Ethnicity = &Ethnicity
	FamilyHistoryAndRiskFactors := makeRandomString()
	patient.FamilyHistoryAndRiskFactors = &FamilyHistoryAndRiskFactors
	FamilyHistoryOfPredispositionSyndrome := makeRandomString()
	patient.FamilyHistoryOfPredispositionSyndrome = &FamilyHistoryOfPredispositionSyndrome
	Gender := makeRandomString()
	patient.Gender = &Gender
	GeneticCancerSyndrome := makeRandomString()
	patient.GeneticCancerSyndrome = &GeneticCancerSyndrome
	OccupationalOrEnvironmentalExposure := makeRandomString()
	patient.OccupationalOrEnvironmentalExposure = &OccupationalOrEnvironmentalExposure
	OtherGeneticConditionOrSignificantComorbidity := makeRandomString()
	patient.OtherGeneticConditionOrSignificantComorbidity = &OtherGeneticConditionOrSignificantComorbidity
	OtherIds := makeRandomString()
	patient.OtherIds = &OtherIds
	PatientID := makeRandomString()
	patient.PatientID = &PatientID
	PriorMalignancy := makeRandomString()
	patient.PriorMalignancy = &PriorMalignancy
	ProvinceOfResidence := makeRandomString()
	patient.ProvinceOfResidence = &ProvinceOfResidence
	Race := makeRandomString()
	patient.Race = &Race
	SampleID := string(sampleID)
	patient.SampleID = &SampleID

	return patient
}

func makeComplication() models.Complication {
	complication := models.Complication{}
	Date := makeRandomDate()
	complication.Date = &Date
	LateComplicationOfTherapyDeveloped := makeRandomString()
	complication.LateComplicationOfTherapyDeveloped = &LateComplicationOfTherapyDeveloped
	LateToxicityDetail := makeRandomString()
	complication.LateToxicityDetail = &LateToxicityDetail
	SampleID := string(sampleID)
	complication.SampleID = &SampleID
	SuspectedTreatmentInducedNeoplasmDeveloped := makeRandomString()
	complication.SuspectedTreatmentInducedNeoplasmDeveloped = &SuspectedTreatmentInducedNeoplasmDeveloped
	TreatmentInducedNeoplasmDetails := makeRandomString()
	complication.TreatmentInducedNeoplasmDetails = &TreatmentInducedNeoplasmDetails

	return complication
}

func makeConsent() models.Consent {
	consent := models.Consent{}
	AssentFormVersion := makeRandomString()
	consent.AssentFormVersion = &AssentFormVersion
	ConsentDate := makeRandomDate()
	consent.ConsentDate = &ConsentDate
	ConsentFormComplete := makeRandomString()
	consent.ConsentFormComplete = &ConsentFormComplete
	ConsentID := makeRandomString()
	consent.ConsentID = &ConsentID
	ConsentingCoordinatorName := makeRandomString()
	consent.ConsentingCoordinatorName = &ConsentingCoordinatorName
	ConsentVersion := makeRandomString()
	consent.ConsentVersion = &ConsentVersion
	DateOfAssent := makeRandomDate()
	consent.DateOfAssent = &DateOfAssent
	DateOfConsentWithdrawal := makeRandomDate()
	consent.DateOfConsentWithdrawal = &DateOfConsentWithdrawal
	HasConsentBeenWithdrawn := makeRandomString()
	consent.HasConsentBeenWithdrawn = &HasConsentBeenWithdrawn
	IfAssentNotObtainedWhyNot := makeRandomString()
	consent.IfAssentNotObtainedWhyNot = &IfAssentNotObtainedWhyNot
	NameOfOtherBiobank := makeRandomString()
	consent.NameOfOtherBiobank = &NameOfOtherBiobank
	PatientConsentedTo := makeRandomString()
	consent.PatientConsentedTo = &PatientConsentedTo
	PreviouslyConsented := makeRandomString()
	consent.PreviouslyConsented = &PreviouslyConsented
	ReasonForConsentWithdrawal := makeRandomString()
	consent.ReasonForConsentWithdrawal = &ReasonForConsentWithdrawal
	ReasonForRejection := makeRandomString()
	consent.ReasonForRejection = &ReasonForRejection
	ReconsentDate := makeRandomString()
	consent.ReconsentDate = &ReconsentDate
	ReconsentVersion := makeRandomString()
	consent.ReconsentVersion = &ReconsentVersion
	SampleID := string(sampleID)
	consent.SampleID = &SampleID
	TypeOfConsentWithdrawal := makeRandomString()
	consent.TypeOfConsentWithdrawal = &TypeOfConsentWithdrawal
	WasAssentObtained := makeRandomString()
	consent.WasAssentObtained = &WasAssentObtained

	return consent
}

func makeDiagnosis() models.Diagnosis {
	diagnosis := models.Diagnosis{}
	AdditionalMolecularTesting := makeRandomString()
	diagnosis.AdditionalMolecularTesting = &AdditionalMolecularTesting
	AgeAtDiagnosis := makeRandomString()
	diagnosis.AgeAtDiagnosis = &AgeAtDiagnosis
	BiomarkerQuantification := makeRandomString()
	diagnosis.BiomarkerQuantification = &BiomarkerQuantification
	CancerSite := makeRandomString()
	diagnosis.CancerSite = &CancerSite
	CancerType := makeRandomString()
	diagnosis.CancerType = &CancerType
	Classification := makeRandomString()
	diagnosis.Classification = &Classification
	DiagnosisDate := makeRandomDate()
	diagnosis.DiagnosisDate = &DiagnosisDate
	DiagnosisID := makeRandomString()
	diagnosis.DiagnosisID = &DiagnosisID
	GradingSystemUsed := makeRandomString()
	diagnosis.GradingSystemUsed = &GradingSystemUsed
	Histology := makeRandomString()
	diagnosis.Histology = &Histology
	MethodOfDefinitiveDiagnosis := makeRandomString()
	diagnosis.MethodOfDefinitiveDiagnosis = &MethodOfDefinitiveDiagnosis
	PrognosticBiomarkers := makeRandomString()
	diagnosis.PrognosticBiomarkers = &PrognosticBiomarkers
	SampleID := string(sampleID)
	diagnosis.SampleID = &SampleID
	SampleSite := makeRandomString()
	diagnosis.SampleSite = &SampleSite
	SampleType := makeRandomString()
	diagnosis.SampleType = &SampleType
	SitesOfMetastases := makeRandomString()
	diagnosis.SitesOfMetastases = &SitesOfMetastases
	SpecificTumourStageAtDiagnosis := makeRandomString()
	diagnosis.SpecificTumourStageAtDiagnosis = &SpecificTumourStageAtDiagnosis
	StagingSystem := makeRandomString()
	diagnosis.StagingSystem = &StagingSystem
	TumourGrade := makeRandomString()
	diagnosis.TumourGrade = &TumourGrade
	VersionOrEditionOfTheStagingSystem := makeRandomString()
	diagnosis.VersionOrEditionOfTheStagingSystem = &VersionOrEditionOfTheStagingSystem

	return diagnosis
}

func makeEnrollment() models.Enrollment {
	enrollment := models.Enrollment{}
	AgeAtEnrollment := string(rand.Intn(600))
	enrollment.AgeAtEnrollment = &AgeAtEnrollment
	CrossEnrollment := makeRandomString()
	enrollment.CrossEnrollment = &CrossEnrollment
	EligibilityAtEnrollment := makeRandomString()
	enrollment.EligibilityAtEnrollment = &EligibilityAtEnrollment
	EnrollmentApprovalDate := makeRandomDate()
	enrollment.EnrollmentApprovalDate = &EnrollmentApprovalDate
	EnrollmentInstitution := makeRandomString()
	enrollment.EnrollmentInstitution = &EnrollmentInstitution
	OtherPersonalizedMedicineStudyID := makeRandomString()
	enrollment.OtherPersonalizedMedicineStudyID = &OtherPersonalizedMedicineStudyID
	OtherPersonalizedMedicineStudyName := makeRandomString()
	enrollment.OtherPersonalizedMedicineStudyName = &OtherPersonalizedMedicineStudyName
	PrimaryOncologistContact := makeRandomString()
	enrollment.PrimaryOncologistContact = &PrimaryOncologistContact
	PrimaryOncologistName := makeRandomString()
	enrollment.PrimaryOncologistName = &PrimaryOncologistName
	ReferringPhysicianContact := makeRandomString()
	enrollment.ReferringPhysicianContact = &ReferringPhysicianContact
	ReferringPhysicianName := makeRandomString()
	enrollment.ReferringPhysicianName = &ReferringPhysicianName
	SampleID := string(sampleID)
	enrollment.SampleID = &SampleID
	StatusAtEnrollment := makeRandomString()
	enrollment.StatusAtEnrollment = &StatusAtEnrollment
	SummaryOfIDRequest := makeRandomString()
	enrollment.SummaryOfIDRequest = &SummaryOfIDRequest
	TreatingCentreName := makeRandomString()
	enrollment.TreatingCentreName = &TreatingCentreName
	TreatingCentreProvince := makeRandomString()
	enrollment.TreatingCentreProvince = &TreatingCentreProvince

	return enrollment
}

func makeOutcome() models.Outcome {
	outcome := models.Outcome{}
	DateOfAssessment := makeRandomDate()
	outcome.DateOfAssessment = &DateOfAssessment
	DiseaseResponseOrStatus := makeRandomString()
	outcome.DiseaseResponseOrStatus = &DiseaseResponseOrStatus
	Height := makeRandomString()
	outcome.Height = &Height
	HeightUnits := makeRandomString()
	outcome.HeightUnits = &HeightUnits
	MethodOfResponseEvaluation := makeRandomString()
	outcome.MethodOfResponseEvaluation = &MethodOfResponseEvaluation
	MinimalResidualDiseaseAssessment := makeRandomString()
	outcome.MinimalResidualDiseaseAssessment = &MinimalResidualDiseaseAssessment
	OtherResponseClassification := makeRandomString()
	outcome.OtherResponseClassification = &OtherResponseClassification
	PerformanceStatus := makeRandomString()
	outcome.PerformanceStatus = &PerformanceStatus
	PhysicalExamID := makeRandomString()
	outcome.PhysicalExamID = &PhysicalExamID
	ResponseCriteriaUsed := makeRandomString()
	outcome.ResponseCriteriaUsed = &ResponseCriteriaUsed
	SampleID := string(sampleID)
	outcome.SampleID = &SampleID
	SitesOfAnyProgressionOrRecurrence := makeRandomString()
	outcome.SitesOfAnyProgressionOrRecurrence = &SitesOfAnyProgressionOrRecurrence
	SummaryStage := makeRandomString()
	outcome.SummaryStage = &SummaryStage
	VitalStatus := makeRandomString()
	outcome.VitalStatus = &VitalStatus
	Weight := makeRandomString()
	outcome.Weight = &Weight
	WeightUnits := makeRandomString()
	outcome.WeightUnits = &WeightUnits

	return outcome
}

func makeTreatment() models.Treatment {
	treatment := models.Treatment{}
	ChemotherapyDetails := makeRandomString()
	treatment.ChemotherapyDetails = &ChemotherapyDetails
	CourseNumber := makeRandomString()
	treatment.CourseNumber = &CourseNumber
	DateOfRecurrenceOrProgressionAfterThisTreatment := makeRandomDate()
	treatment.DateOfRecurrenceOrProgressionAfterThisTreatment = &DateOfRecurrenceOrProgressionAfterThisTreatment
	DrugIDNumbers := makeRandomString()
	treatment.DrugIDNumbers = &DrugIDNumbers
	DrugListOrAgent := makeRandomString()
	treatment.DrugListOrAgent = &DrugListOrAgent
	HematopoieticCellTransplant := makeRandomString()
	treatment.HematopoieticCellTransplant = &HematopoieticCellTransplant
	ImmunotherapyDetails := makeRandomString()
	treatment.ImmunotherapyDetails = &ImmunotherapyDetails
	ProtocolNumberOrCode := makeRandomString()
	treatment.ProtocolNumberOrCode = &ProtocolNumberOrCode
	RadiotherapyDetails := makeRandomString()
	treatment.RadiotherapyDetails = &RadiotherapyDetails
	ReasonForEndingTheTreatment := makeRandomString()
	treatment.ReasonForEndingTheTreatment = &ReasonForEndingTheTreatment
	ResponseCriteriaUsed := makeRandomString()
	treatment.ResponseCriteriaUsed = &ResponseCriteriaUsed
	ResponseToTreatment := makeRandomString()
	treatment.ResponseToTreatment = &ResponseToTreatment
	SampleID := string(sampleID)
	treatment.SampleID = &SampleID
	StartDate := makeRandomDate()
	treatment.StartDate = &StartDate
	StopDate := makeRandomDate()
	treatment.StopDate = &StopDate
	SurgeryDetails := makeRandomString()
	treatment.SurgeryDetails = &SurgeryDetails
	SystematicTherapyAgentName := makeRandomString()
	treatment.SystematicTherapyAgentName = &SystematicTherapyAgentName
	TherapeuticModality := makeRandomString()
	treatment.TherapeuticModality = &TherapeuticModality
	TreatmentIntent := makeRandomString()
	treatment.TreatmentIntent = &TreatmentIntent
	TreatmentPlanType := makeRandomString()
	treatment.TreatmentPlanType = &TreatmentPlanType
	UnexpectedOrUnusualToxicityDuringTreatment := makeRandomString()
	treatment.UnexpectedOrUnusualToxicityDuringTreatment = &UnexpectedOrUnusualToxicityDuringTreatment

	return treatment
}

func makeTumourboard() models.Tumourboard {
	tumourboard := models.Tumourboard{}
	ActionableTargetFound := makeRandomString()
	tumourboard.ActionableTargetFound = &ActionableTargetFound
	AgentOrDrugClass := makeRandomString()
	tumourboard.AgentOrDrugClass = &AgentOrDrugClass
	AnalysesDiscussed := makeRandomString()
	tumourboard.AnalysesDiscussed = &AnalysesDiscussed
	ClassificationOfVariants := makeRandomString()
	tumourboard.ClassificationOfVariants = &ClassificationOfVariants
	ClinicalValidationProgress := makeRandomString()
	tumourboard.ClinicalValidationProgress = &ClinicalValidationProgress
	CnvsDiscussed := makeRandomString()
	tumourboard.CnvsDiscussed = &CnvsDiscussed
	DateOfMolecularTumourBoard := makeRandomDate()
	tumourboard.DateOfMolecularTumourBoard = &DateOfMolecularTumourBoard
	DetailsOfTreatmentPlanImpact := makeRandomString()
	tumourboard.DetailsOfTreatmentPlanImpact = &DetailsOfTreatmentPlanImpact
	DidTreatmentPlanChangeBasedOnProfilingResult := makeRandomString()
	tumourboard.DidTreatmentPlanChangeBasedOnProfilingResult = &DidTreatmentPlanChangeBasedOnProfilingResult
	DiseaseExpressionComparator := makeRandomString()
	tumourboard.DiseaseExpressionComparator = &DiseaseExpressionComparator
	GermlineDnaSampleID := makeRandomString()
	tumourboard.GermlineDnaSampleID = &GermlineDnaSampleID
	GermlineSnvDiscussed := makeRandomString()
	tumourboard.GermlineSnvDiscussed = &GermlineSnvDiscussed
	HasAGermlineVariantBeenIdentifiedByProfilingThatMayPredisposeToCancer := makeRandomString()
	tumourboard.HasAGermlineVariantBeenIdentifiedByProfilingThatMayPredisposeToCancer = &HasAGermlineVariantBeenIdentifiedByProfilingThatMayPredisposeToCancer
	HowTreatmentHasAlteredBasedOnProfiling := makeRandomString()
	tumourboard.HowTreatmentHasAlteredBasedOnProfiling = &HowTreatmentHasAlteredBasedOnProfiling
	LevelOfEvidenceForExpressionTargetAgentMatch := makeRandomString()
	tumourboard.LevelOfEvidenceForExpressionTargetAgentMatch = &LevelOfEvidenceForExpressionTargetAgentMatch
	MolecularTumourBoardRecommendation := makeRandomString()
	tumourboard.MolecularTumourBoardRecommendation = &MolecularTumourBoardRecommendation
	NormalExpressionComparator := makeRandomString()
	tumourboard.NormalExpressionComparator = &NormalExpressionComparator
	PatientHasBeenReferredToAHereditaryCancerProgramBasedOnThisMolecularProfiling := makeRandomString()
	tumourboard.PatientHasBeenReferredToAHereditaryCancerProgramBasedOnThisMolecularProfiling = &PatientHasBeenReferredToAHereditaryCancerProgramBasedOnThisMolecularProfiling
	PatientOrFamilyInformedOfGermlineVariant := makeRandomString()
	tumourboard.PatientOrFamilyInformedOfGermlineVariant = &PatientOrFamilyInformedOfGermlineVariant
	ReasonTreatmentPlanDidNotChangeBasedOnProfiling := makeRandomString()
	tumourboard.ReasonTreatmentPlanDidNotChangeBasedOnProfiling = &ReasonTreatmentPlanDidNotChangeBasedOnProfiling
	SampleID := string(sampleID)
	tumourboard.SampleID = &SampleID
	SomaticSampleType := makeRandomString()
	tumourboard.SomaticSampleType = &SomaticSampleType
	SomaticSnvDiscussed := makeRandomString()
	tumourboard.SomaticSnvDiscussed = &SomaticSnvDiscussed
	StructuralVariantDiscussed := makeRandomString()
	tumourboard.StructuralVariantDiscussed = &StructuralVariantDiscussed
	SummaryReport := makeRandomString()
	tumourboard.SummaryReport = &SummaryReport
	TumourDnaSampleID := makeRandomString()
	tumourboard.TumourDnaSampleID = &TumourDnaSampleID
	TumourRnaSampleID := makeRandomString()
	tumourboard.TumourRnaSampleID = &TumourRnaSampleID
	TypeOfSampleAnalyzed := makeRandomString()
	tumourboard.TypeOfSampleAnalyzed = &TypeOfSampleAnalyzed
	TypeOfTumourSampleAnalyzed := makeRandomString()
	tumourboard.TypeOfTumourSampleAnalyzed = &TypeOfTumourSampleAnalyzed
	TypeOfValidation := makeRandomString()
	tumourboard.TypeOfValidation = &TypeOfValidation

	return tumourboard
}
