// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// AddComplicationCreatedCode is the HTTP code returned for type AddComplicationCreated
const AddComplicationCreatedCode int = 201

/*AddComplicationCreated item created

swagger:response addComplicationCreated
*/
type AddComplicationCreated struct {
}

// NewAddComplicationCreated creates AddComplicationCreated with default headers values
func NewAddComplicationCreated() *AddComplicationCreated {
	return &AddComplicationCreated{}
}

// WriteResponse to the client
func (o *AddComplicationCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(201)
}

// AddComplicationBadRequestCode is the HTTP code returned for type AddComplicationBadRequest
const AddComplicationBadRequestCode int = 400

/*AddComplicationBadRequest invalid input, object invalid

swagger:response addComplicationBadRequest
*/
type AddComplicationBadRequest struct {
}

// NewAddComplicationBadRequest creates AddComplicationBadRequest with default headers values
func NewAddComplicationBadRequest() *AddComplicationBadRequest {
	return &AddComplicationBadRequest{}
}

// WriteResponse to the client
func (o *AddComplicationBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// AddComplicationConflictCode is the HTTP code returned for type AddComplicationConflict
const AddComplicationConflictCode int = 409

/*AddComplicationConflict an existing item already exists

swagger:response addComplicationConflict
*/
type AddComplicationConflict struct {
}

// NewAddComplicationConflict creates AddComplicationConflict with default headers values
func NewAddComplicationConflict() *AddComplicationConflict {
	return &AddComplicationConflict{}
}

// WriteResponse to the client
func (o *AddComplicationConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(409)
}
