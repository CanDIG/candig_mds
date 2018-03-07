// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	errors "github.com/go-openapi/errors"
	loads "github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	security "github.com/go-openapi/runtime/security"
	spec "github.com/go-openapi/spec"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewCandigMetadataAPI creates a new CandigMetadata instance
func NewCandigMetadataAPI(spec *loads.Document) *CandigMetadataAPI {
	return &CandigMetadataAPI{
		handlers:              make(map[string]map[string]http.Handler),
		formats:               strfmt.Default,
		defaultConsumes:       "application/json",
		defaultProduces:       "application/json",
		ServerShutdown:        func() {},
		spec:                  spec,
		ServeError:            errors.ServeError,
		BasicAuthenticator:    security.BasicAuth,
		APIKeyAuthenticator:   security.APIKeyAuth,
		BearerAuthenticator:   security.BearerAuth,
		JSONConsumer:          runtime.JSONConsumer(),
		MultipartformConsumer: runtime.DiscardConsumer,
		JSONProducer:          runtime.JSONProducer(),
		PostUploadHandler: PostUploadHandlerFunc(func(params PostUploadParams) middleware.Responder {
			return middleware.NotImplemented("operation PostUpload has not yet been implemented")
		}),
		AddComplicationHandler: AddComplicationHandlerFunc(func(params AddComplicationParams) middleware.Responder {
			return middleware.NotImplemented("operation AddComplication has not yet been implemented")
		}),
		AddConsentHandler: AddConsentHandlerFunc(func(params AddConsentParams) middleware.Responder {
			return middleware.NotImplemented("operation AddConsent has not yet been implemented")
		}),
		AddDiagnosisHandler: AddDiagnosisHandlerFunc(func(params AddDiagnosisParams) middleware.Responder {
			return middleware.NotImplemented("operation AddDiagnosis has not yet been implemented")
		}),
		AddEnrollmentHandler: AddEnrollmentHandlerFunc(func(params AddEnrollmentParams) middleware.Responder {
			return middleware.NotImplemented("operation AddEnrollment has not yet been implemented")
		}),
		AddOutcomeHandler: AddOutcomeHandlerFunc(func(params AddOutcomeParams) middleware.Responder {
			return middleware.NotImplemented("operation AddOutcome has not yet been implemented")
		}),
		AddPatientHandler: AddPatientHandlerFunc(func(params AddPatientParams) middleware.Responder {
			return middleware.NotImplemented("operation AddPatient has not yet been implemented")
		}),
		AddSampleHandler: AddSampleHandlerFunc(func(params AddSampleParams) middleware.Responder {
			return middleware.NotImplemented("operation AddSample has not yet been implemented")
		}),
		AddTreatmentHandler: AddTreatmentHandlerFunc(func(params AddTreatmentParams) middleware.Responder {
			return middleware.NotImplemented("operation AddTreatment has not yet been implemented")
		}),
		AddTumourboardHandler: AddTumourboardHandlerFunc(func(params AddTumourboardParams) middleware.Responder {
			return middleware.NotImplemented("operation AddTumourboard has not yet been implemented")
		}),
		GetSamplesByQueryHandler: GetSamplesByQueryHandlerFunc(func(params GetSamplesByQueryParams) middleware.Responder {
			return middleware.NotImplemented("operation GetSamplesByQuery has not yet been implemented")
		}),
		LogoutHandler: LogoutHandlerFunc(func(params LogoutParams) middleware.Responder {
			return middleware.NotImplemented("operation Logout has not yet been implemented")
		}),
	}
}

/*CandigMetadataAPI Metadata API */
type CandigMetadataAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for a "application/json" mime type
	JSONConsumer runtime.Consumer
	// MultipartformConsumer registers a consumer for a "multipart/form-data" mime type
	MultipartformConsumer runtime.Consumer

	// JSONProducer registers a producer for a "application/json" mime type
	JSONProducer runtime.Producer

	// PostUploadHandler sets the operation handler for the post upload operation
	PostUploadHandler PostUploadHandler
	// AddComplicationHandler sets the operation handler for the add complication operation
	AddComplicationHandler AddComplicationHandler
	// AddConsentHandler sets the operation handler for the add consent operation
	AddConsentHandler AddConsentHandler
	// AddDiagnosisHandler sets the operation handler for the add diagnosis operation
	AddDiagnosisHandler AddDiagnosisHandler
	// AddEnrollmentHandler sets the operation handler for the add enrollment operation
	AddEnrollmentHandler AddEnrollmentHandler
	// AddOutcomeHandler sets the operation handler for the add outcome operation
	AddOutcomeHandler AddOutcomeHandler
	// AddPatientHandler sets the operation handler for the add patient operation
	AddPatientHandler AddPatientHandler
	// AddSampleHandler sets the operation handler for the add sample operation
	AddSampleHandler AddSampleHandler
	// AddTreatmentHandler sets the operation handler for the add treatment operation
	AddTreatmentHandler AddTreatmentHandler
	// AddTumourboardHandler sets the operation handler for the add tumourboard operation
	AddTumourboardHandler AddTumourboardHandler
	// GetSamplesByQueryHandler sets the operation handler for the get samples by query operation
	GetSamplesByQueryHandler GetSamplesByQueryHandler
	// LogoutHandler sets the operation handler for the logout operation
	LogoutHandler LogoutHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *CandigMetadataAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *CandigMetadataAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *CandigMetadataAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *CandigMetadataAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *CandigMetadataAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *CandigMetadataAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *CandigMetadataAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the CandigMetadataAPI
func (o *CandigMetadataAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.MultipartformConsumer == nil {
		unregistered = append(unregistered, "MultipartformConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.PostUploadHandler == nil {
		unregistered = append(unregistered, "PostUploadHandler")
	}

	if o.AddComplicationHandler == nil {
		unregistered = append(unregistered, "AddComplicationHandler")
	}

	if o.AddConsentHandler == nil {
		unregistered = append(unregistered, "AddConsentHandler")
	}

	if o.AddDiagnosisHandler == nil {
		unregistered = append(unregistered, "AddDiagnosisHandler")
	}

	if o.AddEnrollmentHandler == nil {
		unregistered = append(unregistered, "AddEnrollmentHandler")
	}

	if o.AddOutcomeHandler == nil {
		unregistered = append(unregistered, "AddOutcomeHandler")
	}

	if o.AddPatientHandler == nil {
		unregistered = append(unregistered, "AddPatientHandler")
	}

	if o.AddSampleHandler == nil {
		unregistered = append(unregistered, "AddSampleHandler")
	}

	if o.AddTreatmentHandler == nil {
		unregistered = append(unregistered, "AddTreatmentHandler")
	}

	if o.AddTumourboardHandler == nil {
		unregistered = append(unregistered, "AddTumourboardHandler")
	}

	if o.GetSamplesByQueryHandler == nil {
		unregistered = append(unregistered, "GetSamplesByQueryHandler")
	}

	if o.LogoutHandler == nil {
		unregistered = append(unregistered, "LogoutHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *CandigMetadataAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *CandigMetadataAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	return nil

}

// Authorizer returns the registered authorizer
func (o *CandigMetadataAPI) Authorizer() runtime.Authorizer {

	return nil

}

// ConsumersFor gets the consumers for the specified media types
func (o *CandigMetadataAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {

	result := make(map[string]runtime.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONConsumer

		case "multipart/form-data":
			result["multipart/form-data"] = o.MultipartformConsumer

		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *CandigMetadataAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {

	result := make(map[string]runtime.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONProducer

		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *CandigMetadataAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the candig metadata API
func (o *CandigMetadataAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *CandigMetadataAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/upload"] = NewPostUpload(o.context, o.PostUploadHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/complication"] = NewAddComplication(o.context, o.AddComplicationHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/consent"] = NewAddConsent(o.context, o.AddConsentHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/diagnosis"] = NewAddDiagnosis(o.context, o.AddDiagnosisHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/enrollment"] = NewAddEnrollment(o.context, o.AddEnrollmentHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/outcome"] = NewAddOutcome(o.context, o.AddOutcomeHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/patient"] = NewAddPatient(o.context, o.AddPatientHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/sample"] = NewAddSample(o.context, o.AddSampleHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/treatment"] = NewAddTreatment(o.context, o.AddTreatmentHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/tumourboard"] = NewAddTumourboard(o.context, o.AddTumourboardHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/query"] = NewGetSamplesByQuery(o.context, o.GetSamplesByQueryHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/logout"] = NewLogout(o.context, o.LogoutHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *CandigMetadataAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middelware as you see fit
func (o *CandigMetadataAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}
