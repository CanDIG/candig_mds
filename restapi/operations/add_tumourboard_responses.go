// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// AddTumourboardCreatedCode is the HTTP code returned for type AddTumourboardCreated
const AddTumourboardCreatedCode int = 201

/*AddTumourboardCreated item created

swagger:response addTumourboardCreated
*/
type AddTumourboardCreated struct {
}

// NewAddTumourboardCreated creates AddTumourboardCreated with default headers values
func NewAddTumourboardCreated() *AddTumourboardCreated {
	return &AddTumourboardCreated{}
}

// WriteResponse to the client
func (o *AddTumourboardCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(201)
}

// AddTumourboardBadRequestCode is the HTTP code returned for type AddTumourboardBadRequest
const AddTumourboardBadRequestCode int = 400

/*AddTumourboardBadRequest invalid input, object invalid

swagger:response addTumourboardBadRequest
*/
type AddTumourboardBadRequest struct {
}

// NewAddTumourboardBadRequest creates AddTumourboardBadRequest with default headers values
func NewAddTumourboardBadRequest() *AddTumourboardBadRequest {
	return &AddTumourboardBadRequest{}
}

// WriteResponse to the client
func (o *AddTumourboardBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// AddTumourboardConflictCode is the HTTP code returned for type AddTumourboardConflict
const AddTumourboardConflictCode int = 409

/*AddTumourboardConflict an existing item already exists

swagger:response addTumourboardConflict
*/
type AddTumourboardConflict struct {
}

// NewAddTumourboardConflict creates AddTumourboardConflict with default headers values
func NewAddTumourboardConflict() *AddTumourboardConflict {
	return &AddTumourboardConflict{}
}

// WriteResponse to the client
func (o *AddTumourboardConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(409)
}
