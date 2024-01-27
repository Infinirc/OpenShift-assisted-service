// Code generated by go-swagger; DO NOT EDIT.

package operators

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openshift/assisted-service/models"
)

// V2ListOfClusterOperatorsReader is a Reader for the V2ListOfClusterOperators structure.
type V2ListOfClusterOperatorsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V2ListOfClusterOperatorsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV2ListOfClusterOperatorsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewV2ListOfClusterOperatorsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewV2ListOfClusterOperatorsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewV2ListOfClusterOperatorsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewV2ListOfClusterOperatorsMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewV2ListOfClusterOperatorsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV2ListOfClusterOperatorsOK creates a V2ListOfClusterOperatorsOK with default headers values
func NewV2ListOfClusterOperatorsOK() *V2ListOfClusterOperatorsOK {
	return &V2ListOfClusterOperatorsOK{}
}

/*
V2ListOfClusterOperatorsOK describes a response with status code 200, with default header values.

Success.
*/
type V2ListOfClusterOperatorsOK struct {
	Payload models.MonitoredOperatorsList
}

// IsSuccess returns true when this v2 list of cluster operators o k response has a 2xx status code
func (o *V2ListOfClusterOperatorsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this v2 list of cluster operators o k response has a 3xx status code
func (o *V2ListOfClusterOperatorsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v2 list of cluster operators o k response has a 4xx status code
func (o *V2ListOfClusterOperatorsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this v2 list of cluster operators o k response has a 5xx status code
func (o *V2ListOfClusterOperatorsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this v2 list of cluster operators o k response a status code equal to that given
func (o *V2ListOfClusterOperatorsOK) IsCode(code int) bool {
	return code == 200
}

func (o *V2ListOfClusterOperatorsOK) Error() string {
	return fmt.Sprintf("[GET /v2/clusters/{cluster_id}/monitored-operators][%d] v2ListOfClusterOperatorsOK  %+v", 200, o.Payload)
}

func (o *V2ListOfClusterOperatorsOK) String() string {
	return fmt.Sprintf("[GET /v2/clusters/{cluster_id}/monitored-operators][%d] v2ListOfClusterOperatorsOK  %+v", 200, o.Payload)
}

func (o *V2ListOfClusterOperatorsOK) GetPayload() models.MonitoredOperatorsList {
	return o.Payload
}

func (o *V2ListOfClusterOperatorsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2ListOfClusterOperatorsUnauthorized creates a V2ListOfClusterOperatorsUnauthorized with default headers values
func NewV2ListOfClusterOperatorsUnauthorized() *V2ListOfClusterOperatorsUnauthorized {
	return &V2ListOfClusterOperatorsUnauthorized{}
}

/*
V2ListOfClusterOperatorsUnauthorized describes a response with status code 401, with default header values.

Unauthorized.
*/
type V2ListOfClusterOperatorsUnauthorized struct {
	Payload *models.InfraError
}

// IsSuccess returns true when this v2 list of cluster operators unauthorized response has a 2xx status code
func (o *V2ListOfClusterOperatorsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v2 list of cluster operators unauthorized response has a 3xx status code
func (o *V2ListOfClusterOperatorsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v2 list of cluster operators unauthorized response has a 4xx status code
func (o *V2ListOfClusterOperatorsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this v2 list of cluster operators unauthorized response has a 5xx status code
func (o *V2ListOfClusterOperatorsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this v2 list of cluster operators unauthorized response a status code equal to that given
func (o *V2ListOfClusterOperatorsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *V2ListOfClusterOperatorsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v2/clusters/{cluster_id}/monitored-operators][%d] v2ListOfClusterOperatorsUnauthorized  %+v", 401, o.Payload)
}

func (o *V2ListOfClusterOperatorsUnauthorized) String() string {
	return fmt.Sprintf("[GET /v2/clusters/{cluster_id}/monitored-operators][%d] v2ListOfClusterOperatorsUnauthorized  %+v", 401, o.Payload)
}

func (o *V2ListOfClusterOperatorsUnauthorized) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *V2ListOfClusterOperatorsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2ListOfClusterOperatorsForbidden creates a V2ListOfClusterOperatorsForbidden with default headers values
func NewV2ListOfClusterOperatorsForbidden() *V2ListOfClusterOperatorsForbidden {
	return &V2ListOfClusterOperatorsForbidden{}
}

/*
V2ListOfClusterOperatorsForbidden describes a response with status code 403, with default header values.

Forbidden.
*/
type V2ListOfClusterOperatorsForbidden struct {
	Payload *models.InfraError
}

// IsSuccess returns true when this v2 list of cluster operators forbidden response has a 2xx status code
func (o *V2ListOfClusterOperatorsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v2 list of cluster operators forbidden response has a 3xx status code
func (o *V2ListOfClusterOperatorsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v2 list of cluster operators forbidden response has a 4xx status code
func (o *V2ListOfClusterOperatorsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this v2 list of cluster operators forbidden response has a 5xx status code
func (o *V2ListOfClusterOperatorsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this v2 list of cluster operators forbidden response a status code equal to that given
func (o *V2ListOfClusterOperatorsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *V2ListOfClusterOperatorsForbidden) Error() string {
	return fmt.Sprintf("[GET /v2/clusters/{cluster_id}/monitored-operators][%d] v2ListOfClusterOperatorsForbidden  %+v", 403, o.Payload)
}

func (o *V2ListOfClusterOperatorsForbidden) String() string {
	return fmt.Sprintf("[GET /v2/clusters/{cluster_id}/monitored-operators][%d] v2ListOfClusterOperatorsForbidden  %+v", 403, o.Payload)
}

func (o *V2ListOfClusterOperatorsForbidden) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *V2ListOfClusterOperatorsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2ListOfClusterOperatorsNotFound creates a V2ListOfClusterOperatorsNotFound with default headers values
func NewV2ListOfClusterOperatorsNotFound() *V2ListOfClusterOperatorsNotFound {
	return &V2ListOfClusterOperatorsNotFound{}
}

/*
V2ListOfClusterOperatorsNotFound describes a response with status code 404, with default header values.

Error.
*/
type V2ListOfClusterOperatorsNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this v2 list of cluster operators not found response has a 2xx status code
func (o *V2ListOfClusterOperatorsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v2 list of cluster operators not found response has a 3xx status code
func (o *V2ListOfClusterOperatorsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v2 list of cluster operators not found response has a 4xx status code
func (o *V2ListOfClusterOperatorsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this v2 list of cluster operators not found response has a 5xx status code
func (o *V2ListOfClusterOperatorsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this v2 list of cluster operators not found response a status code equal to that given
func (o *V2ListOfClusterOperatorsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *V2ListOfClusterOperatorsNotFound) Error() string {
	return fmt.Sprintf("[GET /v2/clusters/{cluster_id}/monitored-operators][%d] v2ListOfClusterOperatorsNotFound  %+v", 404, o.Payload)
}

func (o *V2ListOfClusterOperatorsNotFound) String() string {
	return fmt.Sprintf("[GET /v2/clusters/{cluster_id}/monitored-operators][%d] v2ListOfClusterOperatorsNotFound  %+v", 404, o.Payload)
}

func (o *V2ListOfClusterOperatorsNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2ListOfClusterOperatorsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2ListOfClusterOperatorsMethodNotAllowed creates a V2ListOfClusterOperatorsMethodNotAllowed with default headers values
func NewV2ListOfClusterOperatorsMethodNotAllowed() *V2ListOfClusterOperatorsMethodNotAllowed {
	return &V2ListOfClusterOperatorsMethodNotAllowed{}
}

/*
V2ListOfClusterOperatorsMethodNotAllowed describes a response with status code 405, with default header values.

Method Not Allowed.
*/
type V2ListOfClusterOperatorsMethodNotAllowed struct {
	Payload *models.Error
}

// IsSuccess returns true when this v2 list of cluster operators method not allowed response has a 2xx status code
func (o *V2ListOfClusterOperatorsMethodNotAllowed) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v2 list of cluster operators method not allowed response has a 3xx status code
func (o *V2ListOfClusterOperatorsMethodNotAllowed) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v2 list of cluster operators method not allowed response has a 4xx status code
func (o *V2ListOfClusterOperatorsMethodNotAllowed) IsClientError() bool {
	return true
}

// IsServerError returns true when this v2 list of cluster operators method not allowed response has a 5xx status code
func (o *V2ListOfClusterOperatorsMethodNotAllowed) IsServerError() bool {
	return false
}

// IsCode returns true when this v2 list of cluster operators method not allowed response a status code equal to that given
func (o *V2ListOfClusterOperatorsMethodNotAllowed) IsCode(code int) bool {
	return code == 405
}

func (o *V2ListOfClusterOperatorsMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /v2/clusters/{cluster_id}/monitored-operators][%d] v2ListOfClusterOperatorsMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *V2ListOfClusterOperatorsMethodNotAllowed) String() string {
	return fmt.Sprintf("[GET /v2/clusters/{cluster_id}/monitored-operators][%d] v2ListOfClusterOperatorsMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *V2ListOfClusterOperatorsMethodNotAllowed) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2ListOfClusterOperatorsMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2ListOfClusterOperatorsInternalServerError creates a V2ListOfClusterOperatorsInternalServerError with default headers values
func NewV2ListOfClusterOperatorsInternalServerError() *V2ListOfClusterOperatorsInternalServerError {
	return &V2ListOfClusterOperatorsInternalServerError{}
}

/*
V2ListOfClusterOperatorsInternalServerError describes a response with status code 500, with default header values.

Error.
*/
type V2ListOfClusterOperatorsInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this v2 list of cluster operators internal server error response has a 2xx status code
func (o *V2ListOfClusterOperatorsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v2 list of cluster operators internal server error response has a 3xx status code
func (o *V2ListOfClusterOperatorsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v2 list of cluster operators internal server error response has a 4xx status code
func (o *V2ListOfClusterOperatorsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this v2 list of cluster operators internal server error response has a 5xx status code
func (o *V2ListOfClusterOperatorsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this v2 list of cluster operators internal server error response a status code equal to that given
func (o *V2ListOfClusterOperatorsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *V2ListOfClusterOperatorsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v2/clusters/{cluster_id}/monitored-operators][%d] v2ListOfClusterOperatorsInternalServerError  %+v", 500, o.Payload)
}

func (o *V2ListOfClusterOperatorsInternalServerError) String() string {
	return fmt.Sprintf("[GET /v2/clusters/{cluster_id}/monitored-operators][%d] v2ListOfClusterOperatorsInternalServerError  %+v", 500, o.Payload)
}

func (o *V2ListOfClusterOperatorsInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2ListOfClusterOperatorsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
