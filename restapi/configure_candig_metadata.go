package restapi

import (
	"crypto/tls"
	"net/http"
	"strconv"

	config "github.com/Bio-core/jtree/conf"
	keycloak "github.com/Bio-core/keycloakgo"
	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	graceful "github.com/tylerb/graceful"

	"github.com/CanDIG/candig_mds/restapi/operations"
)

var c config.Conf

// This file is safe to edit. Once it exists it will not be overwritten

//go:generate swagger generate server --target .. --name  --spec ../swagger.yaml

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

	api.JSONConsumer = runtime.JSONConsumer()

	api.MultipartformConsumer = runtime.DiscardConsumer

	api.JSONProducer = runtime.JSONProducer()

	api.PostUploadHandler = operations.PostUploadHandlerFunc(func(params operations.PostUploadParams) middleware.Responder {
		return middleware.NotImplemented("operation .PostUpload has not yet been implemented")
	})
	api.AddComplicationHandler = operations.AddComplicationHandlerFunc(func(params operations.AddComplicationParams) middleware.Responder {
		return middleware.NotImplemented("operation .AddComplication has not yet been implemented")
	})
	api.AddConsentHandler = operations.AddConsentHandlerFunc(func(params operations.AddConsentParams) middleware.Responder {
		return middleware.NotImplemented("operation .AddConsent has not yet been implemented")
	})
	api.AddDiagnosisHandler = operations.AddDiagnosisHandlerFunc(func(params operations.AddDiagnosisParams) middleware.Responder {
		return middleware.NotImplemented("operation .AddDiagnosis has not yet been implemented")
	})
	api.AddEnrollmentHandler = operations.AddEnrollmentHandlerFunc(func(params operations.AddEnrollmentParams) middleware.Responder {
		return middleware.NotImplemented("operation .AddEnrollment has not yet been implemented")
	})
	api.AddOutcomeHandler = operations.AddOutcomeHandlerFunc(func(params operations.AddOutcomeParams) middleware.Responder {
		return middleware.NotImplemented("operation .AddOutcome has not yet been implemented")
	})
	api.AddPatientHandler = operations.AddPatientHandlerFunc(func(params operations.AddPatientParams) middleware.Responder {
		return middleware.NotImplemented("operation .AddPatient has not yet been implemented")
	})
	api.AddSampleHandler = operations.AddSampleHandlerFunc(func(params operations.AddSampleParams) middleware.Responder {
		return middleware.NotImplemented("operation .AddSample has not yet been implemented")
	})
	api.AddTreatmentHandler = operations.AddTreatmentHandlerFunc(func(params operations.AddTreatmentParams) middleware.Responder {
		return middleware.NotImplemented("operation .AddTreatment has not yet been implemented")
	})
	api.AddTumourboardHandler = operations.AddTumourboardHandlerFunc(func(params operations.AddTumourboardParams) middleware.Responder {
		return middleware.NotImplemented("operation .AddTumourboard has not yet been implemented")
	})
	api.GetSamplesByQueryHandler = operations.GetSamplesByQueryHandlerFunc(func(params operations.GetSamplesByQueryParams) middleware.Responder {
		return middleware.NotImplemented("operation .GetSamplesByQuery has not yet been implemented")
	})
	api.LogoutHandler = operations.LogoutHandlerFunc(func(params operations.LogoutParams) middleware.Responder {
		return middleware.NotImplemented("operation .Logout has not yet been implemented")
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
