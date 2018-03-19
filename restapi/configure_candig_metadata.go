package restapi

import (
	"crypto/tls"
	"encoding/json"
	localerrors "errors"
	"fmt"
	"hash/fnv"
	"io"
	"net/http"
	"os"
	"strconv"

	config "github.com/Bio-core/jtree/conf"
	keycloak "github.com/Bio-core/keycloakgo"
	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	graceful "github.com/tylerb/graceful"

	"github.com/CanDIG/candig_mds/database"
	"github.com/CanDIG/candig_mds/dummydata"
	"github.com/CanDIG/candig_mds/models"
	"github.com/CanDIG/candig_mds/repos"
	"github.com/CanDIG/candig_mds/restapi/operations"
)

var c config.Conf

func addPatient(patient *models.Patient) error {
	return repos.AddPatient(patient)
}

func addSample(sample *models.Sample) error {
	if sample == nil {
		return errors.New(500, "item must be present")
	}
	if repos.CheckIfSampleExists(*sample.SampleID) {
		return errors.New(500, "Already Exists")
	}
	sample = AddHashes(sample)
	repos.InsertSample(*sample)
	return nil
}

func addEnrollment(enrollment *models.Enrollment) error {
	return repos.AddEnrollment(enrollment)
}

func addConsent(consent *models.Consent) error {
	return repos.AddConsent(consent)
}

func addDiagnosis(diagnosis *models.Diagnosis) error {
	return repos.AddDiagnosis(diagnosis)
}

func addTreatment(treatment *models.Treatment) error {
	return repos.AddTreatment(treatment)
}

func addOutcome(outcome *models.Outcome) error {
	return repos.AddOutcome(outcome)
}

func addComplication(complication *models.Complication) error {
	return repos.AddComplication(complication)
}

func addTumourboard(tumourboard *models.Tumourboard) error {
	return repos.AddTumourboard(tumourboard)
}

func getSamplesByQuery(query *models.Query) []*models.Sample {
	list := repos.GetSamplesByQuery(query)
	return list
}

func logout() bool {
	return true
}

func upload(file operations.PostUploadParams) error {
	if _, err := os.Stat("./uploads/" + file.Filename); !os.IsNotExist(err) {
		return localerrors.New("File already exists")
	}
	f, err := os.OpenFile("./uploads/"+file.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return err
	}
	_, err = io.Copy(f, file.Upfile)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

var databaseFlags = struct {
	Host       string `long:"databaseHost" description:"Database Host" required:"false"`
	Name       string `long:"databaseName" description:"Database Name" required:"false"`
	SelectUser string `long:"dbUsernameSelect" description:"Database Username for Select" required:"false"`
	SelectPass string `long:"dbPasswordSelect" description:"Database Password for Select" required:"false"`
	UpdateUser string `long:"dbUsernameUpdate" description:"Database Username for Update" required:"false"`
	UpdatePass string `long:"dbPasswordUpdate" description:"Database Password for Update" required:"false"`
}{}
var keycloakFlags = struct {
	Active bool   `short:"s" description:"Use Security Bool" required:"false"`
	Host   string `long:"keycloakHost" description:"Keycloak Host" required:"false"`
}{}
var dataGenFlags = struct {
	Generate int `short:"g" description:"generate data" required:"false"`
}{}

func configureFlags(api *operations.CandigMetadataAPI) {
	api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{
		swag.CommandLineOptionsGroup{
			ShortDescription: "Database Flags",
			LongDescription:  "",
			Options:          &databaseFlags,
		},
		swag.CommandLineOptionsGroup{
			ShortDescription: "Keycloak Flags",
			LongDescription:  "",
			Options:          &keycloakFlags,
		},
		swag.CommandLineOptionsGroup{
			ShortDescription: "Data Generation Flags",
			LongDescription:  "",
			Options:          &dataGenFlags,
		},
	}
}

func configureAPI(api *operations.CandigMetadataAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	c.GetConf()
	setupOptions()

	ServerName := c.App.Host + ":" + strconv.Itoa(c.App.Port)
	KeycloakserverName := c.Keycloak.Host

	if keycloakFlags.Active {
		keycloak.Init(KeycloakserverName, "http://"+ServerName, "/Jtree/metadata/0.1.0/columns", "/Jtree/metadata/0.1.0/logout")
	}

	database.Init("candig", "mongodb://localhost:27017/")

	if dataGenFlags.Generate != 0 {
		dummydata.GenerateData(dataGenFlags.Generate)
	}

	api.JSONConsumer = runtime.JSONConsumer()

	api.MultipartformConsumer = runtime.DiscardConsumer

	api.JSONProducer = runtime.JSONProducer()

	api.PostUploadHandler = operations.PostUploadHandlerFunc(func(params operations.PostUploadParams) middleware.Responder {
		if err := upload(params); err != nil {
			return operations.NewPostUploadConflict()
		}
		return operations.NewPostUploadOK().WithPayload(true)
	})
	api.AddComplicationHandler = operations.AddComplicationHandlerFunc(func(params operations.AddComplicationParams) middleware.Responder {
		if err := addComplication(params.Complication); err != nil {
			return operations.NewAddComplicationBadRequest()
		}
		return operations.NewAddComplicationCreated()
	})
	api.AddConsentHandler = operations.AddConsentHandlerFunc(func(params operations.AddConsentParams) middleware.Responder {
		if err := addConsent(params.Consent); err != nil {
			return operations.NewAddConsentBadRequest()
		}
		return operations.NewAddConsentCreated()
	})
	api.AddDiagnosisHandler = operations.AddDiagnosisHandlerFunc(func(params operations.AddDiagnosisParams) middleware.Responder {
		if err := addDiagnosis(params.Diagnosis); err != nil {
			return operations.NewAddDiagnosisBadRequest()
		}
		return operations.NewAddDiagnosisCreated()
	})
	api.AddEnrollmentHandler = operations.AddEnrollmentHandlerFunc(func(params operations.AddEnrollmentParams) middleware.Responder {
		if err := addEnrollment(params.Enrollment); err != nil {
			return operations.NewAddEnrollmentBadRequest()
		}
		return operations.NewAddEnrollmentCreated()
	})
	api.AddOutcomeHandler = operations.AddOutcomeHandlerFunc(func(params operations.AddOutcomeParams) middleware.Responder {
		if err := addOutcome(params.Outcome); err != nil {
			return operations.NewAddOutcomeBadRequest()
		}
		return operations.NewAddOutcomeCreated()
	})
	api.AddPatientHandler = operations.AddPatientHandlerFunc(func(params operations.AddPatientParams) middleware.Responder {
		if err := addPatient(params.Patient); err != nil {
			return operations.NewAddPatientBadRequest()
		}
		return operations.NewAddPatientCreated()
	})
	api.AddSampleHandler = operations.AddSampleHandlerFunc(func(params operations.AddSampleParams) middleware.Responder {
		if err := addSample(params.Sample); err != nil {
			return operations.NewAddSampleBadRequest()
		}
		return operations.NewAddSampleCreated()
	})
	api.AddTreatmentHandler = operations.AddTreatmentHandlerFunc(func(params operations.AddTreatmentParams) middleware.Responder {
		if err := addTreatment(params.Treatment); err != nil {
			return operations.NewAddTreatmentBadRequest()
		}
		return operations.NewAddTreatmentCreated()
	})
	api.AddTumourboardHandler = operations.AddTumourboardHandlerFunc(func(params operations.AddTumourboardParams) middleware.Responder {
		if err := addTumourboard(params.Tumourboard); err != nil {
			return operations.NewAddTumourboardBadRequest()
		}
		return operations.NewAddTreatmentCreated()
	})
	api.GetSamplesByQueryHandler = operations.GetSamplesByQueryHandlerFunc(func(params operations.GetSamplesByQueryParams) middleware.Responder {
		return operations.NewGetSamplesByQueryOK().WithPayload(getSamplesByQuery(params.Query))
	})
	api.LogoutHandler = operations.LogoutHandlerFunc(func(params operations.LogoutParams) middleware.Responder {
		return operations.NewLogoutOK().WithPayload(logout())
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *graceful.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}

func setupOptions() {
	if databaseFlags.Host != "" {
		c.Database.Host = databaseFlags.Host
	}
	if databaseFlags.Name != "" {
		c.Database.Name = databaseFlags.Name
	}
	if databaseFlags.SelectUser != "" {
		c.Database.Selectuser = databaseFlags.SelectUser
	}
	if databaseFlags.SelectPass != "" {
		c.Database.Selectpass = databaseFlags.SelectPass
	}
	if databaseFlags.UpdateUser != "" {
		c.Database.Updateuser = databaseFlags.UpdateUser
	}
	if databaseFlags.UpdatePass != "" {
		c.Database.Updatepass = databaseFlags.UpdatePass
	}
	if keycloakFlags.Host != "" {
		c.Keycloak.Host = keycloakFlags.Host
	}
}

func hash(s string) string {
	h := fnv.New32a()
	h.Write([]byte(s))
	return fmt.Sprint(h.Sum32())
}

//AddHashes for all sub fields
func AddHashes(sample *models.Sample) *models.Sample {
	for _, complication := range sample.Complications {
		jdata, _ := json.Marshal(complication)
		jstring := hash(string(jdata))
		complication.Hash = &jstring
	}
	for _, consent := range sample.Consents {
		jdata, _ := json.Marshal(consent)
		jstring := hash(string(jdata))
		consent.Hash = &jstring
	}
	for _, diagnosis := range sample.Diagnosis {
		jdata, _ := json.Marshal(diagnosis)
		jstring := hash(string(jdata))
		diagnosis.Hash = &jstring
	}
	for _, outcome := range sample.Outcomes {
		jdata, _ := json.Marshal(outcome)
		jstring := hash(string(jdata))
		outcome.Hash = &jstring
	}
	for _, treatment := range sample.Treatments {
		jdata, _ := json.Marshal(treatment)
		jstring := hash(string(jdata))
		treatment.Hash = &jstring
	}
	for _, tumourboard := range sample.Tumourboards {
		jdata, _ := json.Marshal(tumourboard)
		jstring := hash(string(jdata))
		tumourboard.Hash = &jstring
	}
	for _, enrollment := range sample.Enrollments {
		jdata, _ := json.Marshal(enrollment)
		jstring := hash(string(jdata))
		enrollment.Hash = &jstring
	}
	return sample
}
