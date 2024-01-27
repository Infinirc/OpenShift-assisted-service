// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openshift/assisted-service/models"
)

// V2UploadLogsReader is a Reader for the V2UploadLogs structure.
type V2UploadLogsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V2UploadLogsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV2UploadLogsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewV2UploadLogsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewV2UploadLogsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewV2UploadLogsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewV2UploadLogsMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewV2UploadLogsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewV2UploadLogsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewV2UploadLogsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV2UploadLogsNoContent creates a V2UploadLogsNoContent with default headers values
func NewV2UploadLogsNoContent() *V2UploadLogsNoContent {
	return &V2UploadLogsNoContent{}
}

/*
V2UploadLogsNoContent describes a response with status code 204, with default header values.

Success.
*/
type V2UploadLogsNoContent struct {
}

// IsSuccess returns true when this v2 upload logs no content response has a 2xx status code
func (o *V2UploadLogsNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this v2 upload logs no content response has a 3xx status code
func (o *V2UploadLogsNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v2 upload logs no content response has a 4xx status code
func (o *V2UploadLogsNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this v2 upload logs no content response has a 5xx status code
func (o *V2UploadLogsNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this v2 upload logs no content response a status code equal to that given
func (o *V2UploadLogsNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *V2UploadLogsNoContent) Error() string {
	return fmt.Sprintf("[POST /v2/clusters/{cluster_id}/logs][%d] v2UploadLogsNoContent ", 204)
}

func (o *V2UploadLogsNoContent) String() string {
	return fmt.Sprintf("[POST /v2/clusters/{cluster_id}/logs][%d] v2UploadLogsNoContent ", 204)
}

func (o *V2UploadLogsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewV2UploadLogsUnauthorized creates a V2UploadLogsUnauthorized with default headers values
func NewV2UploadLogsUnauthorized() *V2UploadLogsUnauthorized {
	return &V2UploadLogsUnauthorized{}
}

/*
V2UploadLogsUnauthorized describes a response with status code 401, with default header values.

Unauthorized.
*/
type V2UploadLogsUnauthorized struct {
	Payload *models.InfraError
}

// IsSuccess returns true when this v2 upload logs unauthorized response has a 2xx status code
func (o *V2UploadLogsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v2 upload logs unauthorized response has a 3xx status code
func (o *V2UploadLogsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v2 upload logs unauthorized response has a 4xx status code
func (o *V2UploadLogsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this v2 upload logs unauthorized response has a 5xx status code
func (o *V2UploadLogsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this v2 upload logs unauthorized response a status code equal to that given
func (o *V2UploadLogsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *V2UploadLogsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v2/clusters/{cluster_id}/logs][%d] v2UploadLogsUnauthorized  %+v", 401, o.Payload)
}

func (o *V2UploadLogsUnauthorized) String() string {
	return fmt.Sprintf("[POST /v2/clusters/{cluster_id}/logs][%d] v2UploadLogsUnauthorized  %+v", 401, o.Payload)
}

func (o *V2UploadLogsUnauthorized) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *V2UploadLogsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2UploadLogsForbidden creates a V2UploadLogsForbidden with default headers values
func NewV2UploadLogsForbidden() *V2UploadLogsForbidden {
	return &V2UploadLogsForbidden{}
}

/*
V2UploadLogsForbidden describes a response with status code 403, with default header values.

Forbidden.
*/
type V2UploadLogsForbidden struct {
	Payload *models.InfraError
}

// IsSuccess returns true when this v2 upload logs forbidden response has a 2xx status code
func (o *V2UploadLogsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v2 upload logs forbidden response has a 3xx status code
func (o *V2UploadLogsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v2 upload logs forbidden response has a 4xx status code
func (o *V2UploadLogsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this v2 upload logs forbidden response has a 5xx status code
func (o *V2UploadLogsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this v2 upload logs forbidden response a status code equal to that given
func (o *V2UploadLogsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *V2UploadLogsForbidden) Error() string {
	return fmt.Sprintf("[POST /v2/clusters/{cluster_id}/logs][%d] v2UploadLogsForbidden  %+v", 403, o.Payload)
}

func (o *V2UploadLogsForbidden) String() string {
	return fmt.Sprintf("[POST /v2/clusters/{cluster_id}/logs][%d] v2UploadLogsForbidden  %+v", 403, o.Payload)
}

func (o *V2UploadLogsForbidden) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *V2UploadLogsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2UploadLogsNotFound creates a V2UploadLogsNotFound with default headers values
func NewV2UploadLogsNotFound() *V2UploadLogsNotFound {
	return &V2UploadLogsNotFound{}
}

/*
V2UploadLogsNotFound describes a response with status code 404, with default header values.

Error.
*/
type V2UploadLogsNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this v2 upload logs not found response has a 2xx status code
func (o *V2UploadLogsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v2 upload logs not found response has a 3xx status code
func (o *V2UploadLogsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v2 upload logs not found response has a 4xx status code
func (o *V2UploadLogsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this v2 upload logs not found response has a 5xx status code
func (o *V2UploadLogsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this v2 upload logs not found response a status code equal to that given
func (o *V2UploadLogsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *V2UploadLogsNotFound) Error() string {
	return fmt.Sprintf("[POST /v2/clusters/{cluster_id}/logs][%d] v2UploadLogsNotFound  %+v", 404, o.Payload)
}

func (o *V2UploadLogsNotFound) String() string {
	return fmt.Sprintf("[POST /v2/clusters/{cluster_id}/logs][%d] v2UploadLogsNotFound  %+v", 404, o.Payload)
}

func (o *V2UploadLogsNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2UploadLogsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2UploadLogsMethodNotAllowed creates a V2UploadLogsMethodNotAllowed with default headers values
func NewV2UploadLogsMethodNotAllowed() *V2UploadLogsMethodNotAllowed {
	return &V2UploadLogsMethodNotAllowed{}
}

/*
V2UploadLogsMethodNotAllowed describes a response with status code 405, with default header values.

Method Not Allowed.
*/
type V2UploadLogsMethodNotAllowed struct {
	Payload *models.Error
}

// IsSuccess returns true when this v2 upload logs method not allowed response has a 2xx status code
func (o *V2UploadLogsMethodNotAllowed) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v2 upload logs method not allowed response has a 3xx status code
func (o *V2UploadLogsMethodNotAllowed) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v2 upload logs method not allowed response has a 4xx status code
func (o *V2UploadLogsMethodNotAllowed) IsClientError() bool {
	return true
}

// IsServerError returns true when this v2 upload logs method not allowed response has a 5xx status code
func (o *V2UploadLogsMethodNotAllowed) IsServerError() bool {
	return false
}

// IsCode returns true when this v2 upload logs method not allowed response a status code equal to that given
func (o *V2UploadLogsMethodNotAllowed) IsCode(code int) bool {
	return code == 405
}

func (o *V2UploadLogsMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /v2/clusters/{cluster_id}/logs][%d] v2UploadLogsMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *V2UploadLogsMethodNotAllowed) String() string {
	return fmt.Sprintf("[POST /v2/clusters/{cluster_id}/logs][%d] v2UploadLogsMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *V2UploadLogsMethodNotAllowed) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2UploadLogsMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2UploadLogsConflict creates a V2UploadLogsConflict with default headers values
func NewV2UploadLogsConflict() *V2UploadLogsConflict {
	return &V2UploadLogsConflict{}
}

/*
V2UploadLogsConflict describes a response with status code 409, with default header values.

Error.
*/
type V2UploadLogsConflict struct {
	Payload *models.Error
}

// IsSuccess returns true when this v2 upload logs conflict response has a 2xx status code
func (o *V2UploadLogsConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v2 upload logs conflict response has a 3xx status code
func (o *V2UploadLogsConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v2 upload logs conflict response has a 4xx status code
func (o *V2UploadLogsConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this v2 upload logs conflict response has a 5xx status code
func (o *V2UploadLogsConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this v2 upload logs conflict response a status code equal to that given
func (o *V2UploadLogsConflict) IsCode(code int) bool {
	return code == 409
}

func (o *V2UploadLogsConflict) Error() string {
	return fmt.Sprintf("[POST /v2/clusters/{cluster_id}/logs][%d] v2UploadLogsConflict  %+v", 409, o.Payload)
}

func (o *V2UploadLogsConflict) String() string {
	return fmt.Sprintf("[POST /v2/clusters/{cluster_id}/logs][%d] v2UploadLogsConflict  %+v", 409, o.Payload)
}

func (o *V2UploadLogsConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2UploadLogsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2UploadLogsInternalServerError creates a V2UploadLogsInternalServerError with default headers values
func NewV2UploadLogsInternalServerError() *V2UploadLogsInternalServerError {
	return &V2UploadLogsInternalServerError{}
}

/*
V2UploadLogsInternalServerError describes a response with status code 500, with default header values.

Error.
*/
type V2UploadLogsInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this v2 upload logs internal server error response has a 2xx status code
func (o *V2UploadLogsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v2 upload logs internal server error response has a 3xx status code
func (o *V2UploadLogsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v2 upload logs internal server error response has a 4xx status code
func (o *V2UploadLogsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this v2 upload logs internal server error response has a 5xx status code
func (o *V2UploadLogsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this v2 upload logs internal server error response a status code equal to that given
func (o *V2UploadLogsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *V2UploadLogsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v2/clusters/{cluster_id}/logs][%d] v2UploadLogsInternalServerError  %+v", 500, o.Payload)
}

func (o *V2UploadLogsInternalServerError) String() string {
	return fmt.Sprintf("[POST /v2/clusters/{cluster_id}/logs][%d] v2UploadLogsInternalServerError  %+v", 500, o.Payload)
}

func (o *V2UploadLogsInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2UploadLogsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2UploadLogsServiceUnavailable creates a V2UploadLogsServiceUnavailable with default headers values
func NewV2UploadLogsServiceUnavailable() *V2UploadLogsServiceUnavailable {
	return &V2UploadLogsServiceUnavailable{}
}

/*
V2UploadLogsServiceUnavailable describes a response with status code 503, with default header values.

Unavailable.
*/
type V2UploadLogsServiceUnavailable struct {
	Payload *models.Error
}

// IsSuccess returns true when this v2 upload logs service unavailable response has a 2xx status code
func (o *V2UploadLogsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v2 upload logs service unavailable response has a 3xx status code
func (o *V2UploadLogsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v2 upload logs service unavailable response has a 4xx status code
func (o *V2UploadLogsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this v2 upload logs service unavailable response has a 5xx status code
func (o *V2UploadLogsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this v2 upload logs service unavailable response a status code equal to that given
func (o *V2UploadLogsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *V2UploadLogsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /v2/clusters/{cluster_id}/logs][%d] v2UploadLogsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *V2UploadLogsServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /v2/clusters/{cluster_id}/logs][%d] v2UploadLogsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *V2UploadLogsServiceUnavailable) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2UploadLogsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
