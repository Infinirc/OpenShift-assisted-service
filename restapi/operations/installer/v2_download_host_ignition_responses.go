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

// V2DownloadHostIgnitionOKCode is the HTTP code returned for type V2DownloadHostIgnitionOK
const V2DownloadHostIgnitionOKCode int = 200

/*
V2DownloadHostIgnitionOK Success.

swagger:response v2DownloadHostIgnitionOK
*/
type V2DownloadHostIgnitionOK struct {

	/*
	  In: Body
	*/
	Payload io.ReadCloser `json:"body,omitempty"`
}

// NewV2DownloadHostIgnitionOK creates V2DownloadHostIgnitionOK with default headers values
func NewV2DownloadHostIgnitionOK() *V2DownloadHostIgnitionOK {

	return &V2DownloadHostIgnitionOK{}
}

// WithPayload adds the payload to the v2 download host ignition o k response
func (o *V2DownloadHostIgnitionOK) WithPayload(payload io.ReadCloser) *V2DownloadHostIgnitionOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 download host ignition o k response
func (o *V2DownloadHostIgnitionOK) SetPayload(payload io.ReadCloser) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2DownloadHostIgnitionOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// V2DownloadHostIgnitionUnauthorizedCode is the HTTP code returned for type V2DownloadHostIgnitionUnauthorized
const V2DownloadHostIgnitionUnauthorizedCode int = 401

/*
V2DownloadHostIgnitionUnauthorized Unauthorized.

swagger:response v2DownloadHostIgnitionUnauthorized
*/
type V2DownloadHostIgnitionUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewV2DownloadHostIgnitionUnauthorized creates V2DownloadHostIgnitionUnauthorized with default headers values
func NewV2DownloadHostIgnitionUnauthorized() *V2DownloadHostIgnitionUnauthorized {

	return &V2DownloadHostIgnitionUnauthorized{}
}

// WithPayload adds the payload to the v2 download host ignition unauthorized response
func (o *V2DownloadHostIgnitionUnauthorized) WithPayload(payload *models.InfraError) *V2DownloadHostIgnitionUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 download host ignition unauthorized response
func (o *V2DownloadHostIgnitionUnauthorized) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2DownloadHostIgnitionUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2DownloadHostIgnitionForbiddenCode is the HTTP code returned for type V2DownloadHostIgnitionForbidden
const V2DownloadHostIgnitionForbiddenCode int = 403

/*
V2DownloadHostIgnitionForbidden Forbidden.

swagger:response v2DownloadHostIgnitionForbidden
*/
type V2DownloadHostIgnitionForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewV2DownloadHostIgnitionForbidden creates V2DownloadHostIgnitionForbidden with default headers values
func NewV2DownloadHostIgnitionForbidden() *V2DownloadHostIgnitionForbidden {

	return &V2DownloadHostIgnitionForbidden{}
}

// WithPayload adds the payload to the v2 download host ignition forbidden response
func (o *V2DownloadHostIgnitionForbidden) WithPayload(payload *models.InfraError) *V2DownloadHostIgnitionForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 download host ignition forbidden response
func (o *V2DownloadHostIgnitionForbidden) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2DownloadHostIgnitionForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2DownloadHostIgnitionNotFoundCode is the HTTP code returned for type V2DownloadHostIgnitionNotFound
const V2DownloadHostIgnitionNotFoundCode int = 404

/*
V2DownloadHostIgnitionNotFound Error.

swagger:response v2DownloadHostIgnitionNotFound
*/
type V2DownloadHostIgnitionNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewV2DownloadHostIgnitionNotFound creates V2DownloadHostIgnitionNotFound with default headers values
func NewV2DownloadHostIgnitionNotFound() *V2DownloadHostIgnitionNotFound {

	return &V2DownloadHostIgnitionNotFound{}
}

// WithPayload adds the payload to the v2 download host ignition not found response
func (o *V2DownloadHostIgnitionNotFound) WithPayload(payload *models.Error) *V2DownloadHostIgnitionNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 download host ignition not found response
func (o *V2DownloadHostIgnitionNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2DownloadHostIgnitionNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2DownloadHostIgnitionMethodNotAllowedCode is the HTTP code returned for type V2DownloadHostIgnitionMethodNotAllowed
const V2DownloadHostIgnitionMethodNotAllowedCode int = 405

/*
V2DownloadHostIgnitionMethodNotAllowed Method Not Allowed.

swagger:response v2DownloadHostIgnitionMethodNotAllowed
*/
type V2DownloadHostIgnitionMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewV2DownloadHostIgnitionMethodNotAllowed creates V2DownloadHostIgnitionMethodNotAllowed with default headers values
func NewV2DownloadHostIgnitionMethodNotAllowed() *V2DownloadHostIgnitionMethodNotAllowed {

	return &V2DownloadHostIgnitionMethodNotAllowed{}
}

// WithPayload adds the payload to the v2 download host ignition method not allowed response
func (o *V2DownloadHostIgnitionMethodNotAllowed) WithPayload(payload *models.Error) *V2DownloadHostIgnitionMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 download host ignition method not allowed response
func (o *V2DownloadHostIgnitionMethodNotAllowed) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2DownloadHostIgnitionMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2DownloadHostIgnitionConflictCode is the HTTP code returned for type V2DownloadHostIgnitionConflict
const V2DownloadHostIgnitionConflictCode int = 409

/*
V2DownloadHostIgnitionConflict Error.

swagger:response v2DownloadHostIgnitionConflict
*/
type V2DownloadHostIgnitionConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewV2DownloadHostIgnitionConflict creates V2DownloadHostIgnitionConflict with default headers values
func NewV2DownloadHostIgnitionConflict() *V2DownloadHostIgnitionConflict {

	return &V2DownloadHostIgnitionConflict{}
}

// WithPayload adds the payload to the v2 download host ignition conflict response
func (o *V2DownloadHostIgnitionConflict) WithPayload(payload *models.Error) *V2DownloadHostIgnitionConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 download host ignition conflict response
func (o *V2DownloadHostIgnitionConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2DownloadHostIgnitionConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2DownloadHostIgnitionInternalServerErrorCode is the HTTP code returned for type V2DownloadHostIgnitionInternalServerError
const V2DownloadHostIgnitionInternalServerErrorCode int = 500

/*
V2DownloadHostIgnitionInternalServerError Error.

swagger:response v2DownloadHostIgnitionInternalServerError
*/
type V2DownloadHostIgnitionInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewV2DownloadHostIgnitionInternalServerError creates V2DownloadHostIgnitionInternalServerError with default headers values
func NewV2DownloadHostIgnitionInternalServerError() *V2DownloadHostIgnitionInternalServerError {

	return &V2DownloadHostIgnitionInternalServerError{}
}

// WithPayload adds the payload to the v2 download host ignition internal server error response
func (o *V2DownloadHostIgnitionInternalServerError) WithPayload(payload *models.Error) *V2DownloadHostIgnitionInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 download host ignition internal server error response
func (o *V2DownloadHostIgnitionInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2DownloadHostIgnitionInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2DownloadHostIgnitionServiceUnavailableCode is the HTTP code returned for type V2DownloadHostIgnitionServiceUnavailable
const V2DownloadHostIgnitionServiceUnavailableCode int = 503

/*
V2DownloadHostIgnitionServiceUnavailable Unavailable.

swagger:response v2DownloadHostIgnitionServiceUnavailable
*/
type V2DownloadHostIgnitionServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewV2DownloadHostIgnitionServiceUnavailable creates V2DownloadHostIgnitionServiceUnavailable with default headers values
func NewV2DownloadHostIgnitionServiceUnavailable() *V2DownloadHostIgnitionServiceUnavailable {

	return &V2DownloadHostIgnitionServiceUnavailable{}
}

// WithPayload adds the payload to the v2 download host ignition service unavailable response
func (o *V2DownloadHostIgnitionServiceUnavailable) WithPayload(payload *models.Error) *V2DownloadHostIgnitionServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 download host ignition service unavailable response
func (o *V2DownloadHostIgnitionServiceUnavailable) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2DownloadHostIgnitionServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
