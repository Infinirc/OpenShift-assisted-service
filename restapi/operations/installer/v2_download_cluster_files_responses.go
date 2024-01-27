// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openshift/assisted-service/models"
)

// V2DownloadClusterFilesOKCode is the HTTP code returned for type V2DownloadClusterFilesOK
const V2DownloadClusterFilesOKCode int = 200

/*
V2DownloadClusterFilesOK Success.

swagger:response v2DownloadClusterFilesOK
*/
type V2DownloadClusterFilesOK struct {

	/*
	  In: Body
	*/
	Payload io.ReadCloser `json:"body,omitempty"`
}

// NewV2DownloadClusterFilesOK creates V2DownloadClusterFilesOK with default headers values
func NewV2DownloadClusterFilesOK() *V2DownloadClusterFilesOK {

	return &V2DownloadClusterFilesOK{}
}

// WithPayload adds the payload to the v2 download cluster files o k response
func (o *V2DownloadClusterFilesOK) WithPayload(payload io.ReadCloser) *V2DownloadClusterFilesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 download cluster files o k response
func (o *V2DownloadClusterFilesOK) SetPayload(payload io.ReadCloser) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2DownloadClusterFilesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// V2DownloadClusterFilesUnauthorizedCode is the HTTP code returned for type V2DownloadClusterFilesUnauthorized
const V2DownloadClusterFilesUnauthorizedCode int = 401

/*
V2DownloadClusterFilesUnauthorized Unauthorized.

swagger:response v2DownloadClusterFilesUnauthorized
*/
type V2DownloadClusterFilesUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewV2DownloadClusterFilesUnauthorized creates V2DownloadClusterFilesUnauthorized with default headers values
func NewV2DownloadClusterFilesUnauthorized() *V2DownloadClusterFilesUnauthorized {

	return &V2DownloadClusterFilesUnauthorized{}
}

// WithPayload adds the payload to the v2 download cluster files unauthorized response
func (o *V2DownloadClusterFilesUnauthorized) WithPayload(payload *models.InfraError) *V2DownloadClusterFilesUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 download cluster files unauthorized response
func (o *V2DownloadClusterFilesUnauthorized) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2DownloadClusterFilesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2DownloadClusterFilesForbiddenCode is the HTTP code returned for type V2DownloadClusterFilesForbidden
const V2DownloadClusterFilesForbiddenCode int = 403

/*
V2DownloadClusterFilesForbidden Forbidden.

swagger:response v2DownloadClusterFilesForbidden
*/
type V2DownloadClusterFilesForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewV2DownloadClusterFilesForbidden creates V2DownloadClusterFilesForbidden with default headers values
func NewV2DownloadClusterFilesForbidden() *V2DownloadClusterFilesForbidden {

	return &V2DownloadClusterFilesForbidden{}
}

// WithPayload adds the payload to the v2 download cluster files forbidden response
func (o *V2DownloadClusterFilesForbidden) WithPayload(payload *models.InfraError) *V2DownloadClusterFilesForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 download cluster files forbidden response
func (o *V2DownloadClusterFilesForbidden) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2DownloadClusterFilesForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2DownloadClusterFilesNotFoundCode is the HTTP code returned for type V2DownloadClusterFilesNotFound
const V2DownloadClusterFilesNotFoundCode int = 404

/*
V2DownloadClusterFilesNotFound Error.

swagger:response v2DownloadClusterFilesNotFound
*/
type V2DownloadClusterFilesNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewV2DownloadClusterFilesNotFound creates V2DownloadClusterFilesNotFound with default headers values
func NewV2DownloadClusterFilesNotFound() *V2DownloadClusterFilesNotFound {

	return &V2DownloadClusterFilesNotFound{}
}

// WithPayload adds the payload to the v2 download cluster files not found response
func (o *V2DownloadClusterFilesNotFound) WithPayload(payload *models.Error) *V2DownloadClusterFilesNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 download cluster files not found response
func (o *V2DownloadClusterFilesNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2DownloadClusterFilesNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2DownloadClusterFilesMethodNotAllowedCode is the HTTP code returned for type V2DownloadClusterFilesMethodNotAllowed
const V2DownloadClusterFilesMethodNotAllowedCode int = 405

/*
V2DownloadClusterFilesMethodNotAllowed Method Not Allowed.

swagger:response v2DownloadClusterFilesMethodNotAllowed
*/
type V2DownloadClusterFilesMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewV2DownloadClusterFilesMethodNotAllowed creates V2DownloadClusterFilesMethodNotAllowed with default headers values
func NewV2DownloadClusterFilesMethodNotAllowed() *V2DownloadClusterFilesMethodNotAllowed {

	return &V2DownloadClusterFilesMethodNotAllowed{}
}

// WithPayload adds the payload to the v2 download cluster files method not allowed response
func (o *V2DownloadClusterFilesMethodNotAllowed) WithPayload(payload *models.Error) *V2DownloadClusterFilesMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 download cluster files method not allowed response
func (o *V2DownloadClusterFilesMethodNotAllowed) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2DownloadClusterFilesMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2DownloadClusterFilesConflictCode is the HTTP code returned for type V2DownloadClusterFilesConflict
const V2DownloadClusterFilesConflictCode int = 409

/*
V2DownloadClusterFilesConflict Error.

swagger:response v2DownloadClusterFilesConflict
*/
type V2DownloadClusterFilesConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewV2DownloadClusterFilesConflict creates V2DownloadClusterFilesConflict with default headers values
func NewV2DownloadClusterFilesConflict() *V2DownloadClusterFilesConflict {

	return &V2DownloadClusterFilesConflict{}
}

// WithPayload adds the payload to the v2 download cluster files conflict response
func (o *V2DownloadClusterFilesConflict) WithPayload(payload *models.Error) *V2DownloadClusterFilesConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 download cluster files conflict response
func (o *V2DownloadClusterFilesConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2DownloadClusterFilesConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2DownloadClusterFilesInternalServerErrorCode is the HTTP code returned for type V2DownloadClusterFilesInternalServerError
const V2DownloadClusterFilesInternalServerErrorCode int = 500

/*
V2DownloadClusterFilesInternalServerError Error.

swagger:response v2DownloadClusterFilesInternalServerError
*/
type V2DownloadClusterFilesInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewV2DownloadClusterFilesInternalServerError creates V2DownloadClusterFilesInternalServerError with default headers values
func NewV2DownloadClusterFilesInternalServerError() *V2DownloadClusterFilesInternalServerError {

	return &V2DownloadClusterFilesInternalServerError{}
}

// WithPayload adds the payload to the v2 download cluster files internal server error response
func (o *V2DownloadClusterFilesInternalServerError) WithPayload(payload *models.Error) *V2DownloadClusterFilesInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 download cluster files internal server error response
func (o *V2DownloadClusterFilesInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2DownloadClusterFilesInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2DownloadClusterFilesServiceUnavailableCode is the HTTP code returned for type V2DownloadClusterFilesServiceUnavailable
const V2DownloadClusterFilesServiceUnavailableCode int = 503

/*
V2DownloadClusterFilesServiceUnavailable Unavailable.

swagger:response v2DownloadClusterFilesServiceUnavailable
*/
type V2DownloadClusterFilesServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewV2DownloadClusterFilesServiceUnavailable creates V2DownloadClusterFilesServiceUnavailable with default headers values
func NewV2DownloadClusterFilesServiceUnavailable() *V2DownloadClusterFilesServiceUnavailable {

	return &V2DownloadClusterFilesServiceUnavailable{}
}

// WithPayload adds the payload to the v2 download cluster files service unavailable response
func (o *V2DownloadClusterFilesServiceUnavailable) WithPayload(payload *models.Error) *V2DownloadClusterFilesServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 download cluster files service unavailable response
func (o *V2DownloadClusterFilesServiceUnavailable) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2DownloadClusterFilesServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
