package restapi

import (
	"crypto/tls"
	localerrors "errors"
	"fmt"
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
	"github.com/CanDIG/candig_mds/models"
	"github.com/CanDIG/candig_mds/repos"
	"github.com/CanDIG/candig_mds/restapi/operations"
)

var c config.Conf

func addPatient(patient *models.Patient) error {
	if patient == nil {
		return errors.New(500, "item must be present")
	}
	repos.InsertPatient(*patient)
	return nil
}

func addSample(sample *models.Sample) error {
	if sample == nil {
		return errors.New(500, "item must be present")
	}
	repos.InsertSample(*sample)
	return nil
}

func addEnrollment(enrollment *models.Enrollment) error {
	if enrollment == nil {
		return errors.New(500, "item must be present")
	}
	repos.InsertEnrollment(*enrollment)
	return nil
}

func addConsent(consent *models.Consent) error {
	if consent == nil {
		return errors.New(500, "item must be present")
	}
	repos.InsertConsent(*consent)
	return nil
}

func addDiagnosis(diagnosis *models.Diagnosis) error {
	if diagnosis == nil {
		return errors.New(500, "item must be present")
	}
	repos.InsertDiagnosis(*diagnosis)
	return nil
}

func addTreatment(treatment *models.Treatment) error {
	if treatment == nil {
		return errors.New(500, "item must be present")
	}
	repos.InsertTreatment(*treatment)
	return nil
}

func addOutcome(outcome *models.Outcome) error {
	if outcome == nil {
		return errors.New(500, "item must be present")
	}
	repos.InsertOutcome(*outcome)
	return nil
}

func addComplication(complication *models.Complication) error {
	if complication == nil {
		return errors.New(500, "item must be present")
	}
	repos.InsertComplication(*complication)
	return nil
}

func addTumourboard(tumourboard *models.Tumourboard) error {
	if tumourboard == nil {
		return errors.New(500, "item must be present")
	}
	repos.InsertTumourboard(*tumourboard)
	return nil
}

func getSamplesByQuery(query *models.Query) []*models.Sample {
	list := repos.GetAllSamples("sample")
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

	database.DatabaseInit("candig", "mongodb://localhost:27017/")

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
