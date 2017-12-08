// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// AddBiosampleCreatedCode is the HTTP code returned for type AddBiosampleCreated
const AddBiosampleCreatedCode int = 201

/*AddBiosampleCreated item created

swagger:response addBiosampleCreated
*/
type AddBiosampleCreated struct {
}

// NewAddBiosampleCreated creates AddBiosampleCreated with default headers values
func NewAddBiosampleCreated() *AddBiosampleCreated {
	return &AddBiosampleCreated{}
}

// WriteResponse to the client
func (o *AddBiosampleCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(201)
}

// AddBiosampleBadRequestCode is the HTTP code returned for type AddBiosampleBadRequest
const AddBiosampleBadRequestCode int = 400

/*AddBiosampleBadRequest invalid input, object invalid

swagger:response addBiosampleBadRequest
*/
type AddBiosampleBadRequest struct {
}

// NewAddBiosampleBadRequest creates AddBiosampleBadRequest with default headers values
func NewAddBiosampleBadRequest() *AddBiosampleBadRequest {
	return &AddBiosampleBadRequest{}
}

// WriteResponse to the client
func (o *AddBiosampleBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// AddBiosampleConflictCode is the HTTP code returned for type AddBiosampleConflict
const AddBiosampleConflictCode int = 409

/*AddBiosampleConflict an existing item already exists

swagger:response addBiosampleConflict
*/
type AddBiosampleConflict struct {
}

// NewAddBiosampleConflict creates AddBiosampleConflict with default headers values
func NewAddBiosampleConflict() *AddBiosampleConflict {
	return &AddBiosampleConflict{}
}

// WriteResponse to the client
func (o *AddBiosampleConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(409)
}
