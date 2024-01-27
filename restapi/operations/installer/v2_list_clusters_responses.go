// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openshift/assisted-service/models"
)

// V2ListClustersOKCode is the HTTP code returned for type V2ListClustersOK
const V2ListClustersOKCode int = 200

/*
V2ListClustersOK Success.

swagger:response v2ListClustersOK
*/
type V2ListClustersOK struct {

	/*
	  In: Body
	*/
	Payload models.ClusterList `json:"body,omitempty"`
}

// NewV2ListClustersOK creates V2ListClustersOK with default headers values
func NewV2ListClustersOK() *V2ListClustersOK {

	return &V2ListClustersOK{}
}

// WithPayload adds the payload to the v2 list clusters o k response
func (o *V2ListClustersOK) WithPayload(payload models.ClusterList) *V2ListClustersOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 list clusters o k response
func (o *V2ListClustersOK) SetPayload(payload models.ClusterList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2ListClustersOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.ClusterList{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// V2ListClustersUnauthorizedCode is the HTTP code returned for type V2ListClustersUnauthorized
const V2ListClustersUnauthorizedCode int = 401

/*
V2ListClustersUnauthorized Unauthorized.

swagger:response v2ListClustersUnauthorized
*/
type V2ListClustersUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewV2ListClustersUnauthorized creates V2ListClustersUnauthorized with default headers values
func NewV2ListClustersUnauthorized() *V2ListClustersUnauthorized {

	return &V2ListClustersUnauthorized{}
}

// WithPayload adds the payload to the v2 list clusters unauthorized response
func (o *V2ListClustersUnauthorized) WithPayload(payload *models.InfraError) *V2ListClustersUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 list clusters unauthorized response
func (o *V2ListClustersUnauthorized) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2ListClustersUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2ListClustersForbiddenCode is the HTTP code returned for type V2ListClustersForbidden
const V2ListClustersForbiddenCode int = 403

/*
V2ListClustersForbidden Forbidden.

swagger:response v2ListClustersForbidden
*/
type V2ListClustersForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewV2ListClustersForbidden creates V2ListClustersForbidden with default headers values
func NewV2ListClustersForbidden() *V2ListClustersForbidden {

	return &V2ListClustersForbidden{}
}

// WithPayload adds the payload to the v2 list clusters forbidden response
func (o *V2ListClustersForbidden) WithPayload(payload *models.InfraError) *V2ListClustersForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 list clusters forbidden response
func (o *V2ListClustersForbidden) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2ListClustersForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2ListClustersMethodNotAllowedCode is the HTTP code returned for type V2ListClustersMethodNotAllowed
const V2ListClustersMethodNotAllowedCode int = 405

/*
V2ListClustersMethodNotAllowed Method Not Allowed.

swagger:response v2ListClustersMethodNotAllowed
*/
type V2ListClustersMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewV2ListClustersMethodNotAllowed creates V2ListClustersMethodNotAllowed with default headers values
func NewV2ListClustersMethodNotAllowed() *V2ListClustersMethodNotAllowed {

	return &V2ListClustersMethodNotAllowed{}
}

// WithPayload adds the payload to the v2 list clusters method not allowed response
func (o *V2ListClustersMethodNotAllowed) WithPayload(payload *models.Error) *V2ListClustersMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 list clusters method not allowed response
func (o *V2ListClustersMethodNotAllowed) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2ListClustersMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2ListClustersInternalServerErrorCode is the HTTP code returned for type V2ListClustersInternalServerError
const V2ListClustersInternalServerErrorCode int = 500

/*
V2ListClustersInternalServerError Error.

swagger:response v2ListClustersInternalServerError
*/
type V2ListClustersInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewV2ListClustersInternalServerError creates V2ListClustersInternalServerError with default headers values
func NewV2ListClustersInternalServerError() *V2ListClustersInternalServerError {

	return &V2ListClustersInternalServerError{}
}

// WithPayload adds the payload to the v2 list clusters internal server error response
func (o *V2ListClustersInternalServerError) WithPayload(payload *models.Error) *V2ListClustersInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 list clusters internal server error response
func (o *V2ListClustersInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2ListClustersInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2ListClustersServiceUnavailableCode is the HTTP code returned for type V2ListClustersServiceUnavailable
const V2ListClustersServiceUnavailableCode int = 503

/*
V2ListClustersServiceUnavailable Unavailable.

swagger:response v2ListClustersServiceUnavailable
*/
type V2ListClustersServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewV2ListClustersServiceUnavailable creates V2ListClustersServiceUnavailable with default headers values
func NewV2ListClustersServiceUnavailable() *V2ListClustersServiceUnavailable {

	return &V2ListClustersServiceUnavailable{}
}

// WithPayload adds the payload to the v2 list clusters service unavailable response
func (o *V2ListClustersServiceUnavailable) WithPayload(payload *models.Error) *V2ListClustersServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 list clusters service unavailable response
func (o *V2ListClustersServiceUnavailable) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2ListClustersServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
