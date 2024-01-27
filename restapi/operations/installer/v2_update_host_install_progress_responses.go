// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openshift/assisted-service/models"
)

// V2UpdateHostInstallProgressOKCode is the HTTP code returned for type V2UpdateHostInstallProgressOK
const V2UpdateHostInstallProgressOKCode int = 200

/*
V2UpdateHostInstallProgressOK Update install progress.

swagger:response v2UpdateHostInstallProgressOK
*/
type V2UpdateHostInstallProgressOK struct {
}

// NewV2UpdateHostInstallProgressOK creates V2UpdateHostInstallProgressOK with default headers values
func NewV2UpdateHostInstallProgressOK() *V2UpdateHostInstallProgressOK {

	return &V2UpdateHostInstallProgressOK{}
}

// WriteResponse to the client
func (o *V2UpdateHostInstallProgressOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// V2UpdateHostInstallProgressUnauthorizedCode is the HTTP code returned for type V2UpdateHostInstallProgressUnauthorized
const V2UpdateHostInstallProgressUnauthorizedCode int = 401

/*
V2UpdateHostInstallProgressUnauthorized Unauthorized.

swagger:response v2UpdateHostInstallProgressUnauthorized
*/
type V2UpdateHostInstallProgressUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewV2UpdateHostInstallProgressUnauthorized creates V2UpdateHostInstallProgressUnauthorized with default headers values
func NewV2UpdateHostInstallProgressUnauthorized() *V2UpdateHostInstallProgressUnauthorized {

	return &V2UpdateHostInstallProgressUnauthorized{}
}

// WithPayload adds the payload to the v2 update host install progress unauthorized response
func (o *V2UpdateHostInstallProgressUnauthorized) WithPayload(payload *models.InfraError) *V2UpdateHostInstallProgressUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 update host install progress unauthorized response
func (o *V2UpdateHostInstallProgressUnauthorized) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2UpdateHostInstallProgressUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2UpdateHostInstallProgressForbiddenCode is the HTTP code returned for type V2UpdateHostInstallProgressForbidden
const V2UpdateHostInstallProgressForbiddenCode int = 403

/*
V2UpdateHostInstallProgressForbidden Forbidden.

swagger:response v2UpdateHostInstallProgressForbidden
*/
type V2UpdateHostInstallProgressForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewV2UpdateHostInstallProgressForbidden creates V2UpdateHostInstallProgressForbidden with default headers values
func NewV2UpdateHostInstallProgressForbidden() *V2UpdateHostInstallProgressForbidden {

	return &V2UpdateHostInstallProgressForbidden{}
}

// WithPayload adds the payload to the v2 update host install progress forbidden response
func (o *V2UpdateHostInstallProgressForbidden) WithPayload(payload *models.InfraError) *V2UpdateHostInstallProgressForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 update host install progress forbidden response
func (o *V2UpdateHostInstallProgressForbidden) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2UpdateHostInstallProgressForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2UpdateHostInstallProgressNotFoundCode is the HTTP code returned for type V2UpdateHostInstallProgressNotFound
const V2UpdateHostInstallProgressNotFoundCode int = 404

/*
V2UpdateHostInstallProgressNotFound Error.

swagger:response v2UpdateHostInstallProgressNotFound
*/
type V2UpdateHostInstallProgressNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewV2UpdateHostInstallProgressNotFound creates V2UpdateHostInstallProgressNotFound with default headers values
func NewV2UpdateHostInstallProgressNotFound() *V2UpdateHostInstallProgressNotFound {

	return &V2UpdateHostInstallProgressNotFound{}
}

// WithPayload adds the payload to the v2 update host install progress not found response
func (o *V2UpdateHostInstallProgressNotFound) WithPayload(payload *models.Error) *V2UpdateHostInstallProgressNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 update host install progress not found response
func (o *V2UpdateHostInstallProgressNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2UpdateHostInstallProgressNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2UpdateHostInstallProgressMethodNotAllowedCode is the HTTP code returned for type V2UpdateHostInstallProgressMethodNotAllowed
const V2UpdateHostInstallProgressMethodNotAllowedCode int = 405

/*
V2UpdateHostInstallProgressMethodNotAllowed Method Not Allowed.

swagger:response v2UpdateHostInstallProgressMethodNotAllowed
*/
type V2UpdateHostInstallProgressMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewV2UpdateHostInstallProgressMethodNotAllowed creates V2UpdateHostInstallProgressMethodNotAllowed with default headers values
func NewV2UpdateHostInstallProgressMethodNotAllowed() *V2UpdateHostInstallProgressMethodNotAllowed {

	return &V2UpdateHostInstallProgressMethodNotAllowed{}
}

// WithPayload adds the payload to the v2 update host install progress method not allowed response
func (o *V2UpdateHostInstallProgressMethodNotAllowed) WithPayload(payload *models.Error) *V2UpdateHostInstallProgressMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 update host install progress method not allowed response
func (o *V2UpdateHostInstallProgressMethodNotAllowed) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2UpdateHostInstallProgressMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2UpdateHostInstallProgressInternalServerErrorCode is the HTTP code returned for type V2UpdateHostInstallProgressInternalServerError
const V2UpdateHostInstallProgressInternalServerErrorCode int = 500

/*
V2UpdateHostInstallProgressInternalServerError Error.

swagger:response v2UpdateHostInstallProgressInternalServerError
*/
type V2UpdateHostInstallProgressInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewV2UpdateHostInstallProgressInternalServerError creates V2UpdateHostInstallProgressInternalServerError with default headers values
func NewV2UpdateHostInstallProgressInternalServerError() *V2UpdateHostInstallProgressInternalServerError {

	return &V2UpdateHostInstallProgressInternalServerError{}
}

// WithPayload adds the payload to the v2 update host install progress internal server error response
func (o *V2UpdateHostInstallProgressInternalServerError) WithPayload(payload *models.Error) *V2UpdateHostInstallProgressInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 update host install progress internal server error response
func (o *V2UpdateHostInstallProgressInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2UpdateHostInstallProgressInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2UpdateHostInstallProgressServiceUnavailableCode is the HTTP code returned for type V2UpdateHostInstallProgressServiceUnavailable
const V2UpdateHostInstallProgressServiceUnavailableCode int = 503

/*
V2UpdateHostInstallProgressServiceUnavailable Unavailable.

swagger:response v2UpdateHostInstallProgressServiceUnavailable
*/
type V2UpdateHostInstallProgressServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewV2UpdateHostInstallProgressServiceUnavailable creates V2UpdateHostInstallProgressServiceUnavailable with default headers values
func NewV2UpdateHostInstallProgressServiceUnavailable() *V2UpdateHostInstallProgressServiceUnavailable {

	return &V2UpdateHostInstallProgressServiceUnavailable{}
}

// WithPayload adds the payload to the v2 update host install progress service unavailable response
func (o *V2UpdateHostInstallProgressServiceUnavailable) WithPayload(payload *models.Error) *V2UpdateHostInstallProgressServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 update host install progress service unavailable response
func (o *V2UpdateHostInstallProgressServiceUnavailable) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2UpdateHostInstallProgressServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
