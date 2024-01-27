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

// UnbindHostReader is a Reader for the UnbindHost structure.
type UnbindHostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UnbindHostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUnbindHostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUnbindHostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUnbindHostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUnbindHostForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUnbindHostNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewUnbindHostMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUnbindHostConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUnbindHostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 501:
		result := NewUnbindHostNotImplemented()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewUnbindHostServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUnbindHostOK creates a UnbindHostOK with default headers values
func NewUnbindHostOK() *UnbindHostOK {
	return &UnbindHostOK{}
}

/*
UnbindHostOK describes a response with status code 200, with default header values.

Success.
*/
type UnbindHostOK struct {
	Payload *models.Host
}

// IsSuccess returns true when this unbind host o k response has a 2xx status code
func (o *UnbindHostOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this unbind host o k response has a 3xx status code
func (o *UnbindHostOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this unbind host o k response has a 4xx status code
func (o *UnbindHostOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this unbind host o k response has a 5xx status code
func (o *UnbindHostOK) IsServerError() bool {
	return false
}

// IsCode returns true when this unbind host o k response a status code equal to that given
func (o *UnbindHostOK) IsCode(code int) bool {
	return code == 200
}

func (o *UnbindHostOK) Error() string {
	return fmt.Sprintf("[POST /v2/infra-envs/{infra_env_id}/hosts/{host_id}/actions/unbind][%d] unbindHostOK  %+v", 200, o.Payload)
}

func (o *UnbindHostOK) String() string {
	return fmt.Sprintf("[POST /v2/infra-envs/{infra_env_id}/hosts/{host_id}/actions/unbind][%d] unbindHostOK  %+v", 200, o.Payload)
}

func (o *UnbindHostOK) GetPayload() *models.Host {
	return o.Payload
}

func (o *UnbindHostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Host)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUnbindHostBadRequest creates a UnbindHostBadRequest with default headers values
func NewUnbindHostBadRequest() *UnbindHostBadRequest {
	return &UnbindHostBadRequest{}
}

/*
UnbindHostBadRequest describes a response with status code 400, with default header values.

Error.
*/
type UnbindHostBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this unbind host bad request response has a 2xx status code
func (o *UnbindHostBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this unbind host bad request response has a 3xx status code
func (o *UnbindHostBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this unbind host bad request response has a 4xx status code
func (o *UnbindHostBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this unbind host bad request response has a 5xx status code
func (o *UnbindHostBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this unbind host bad request response a status code equal to that given
func (o *UnbindHostBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *UnbindHostBadRequest) Error() string {
	return fmt.Sprintf("[POST /v2/infra-envs/{infra_env_id}/hosts/{host_id}/actions/unbind][%d] unbindHostBadRequest  %+v", 400, o.Payload)
}

func (o *UnbindHostBadRequest) String() string {
	return fmt.Sprintf("[POST /v2/infra-envs/{infra_env_id}/hosts/{host_id}/actions/unbind][%d] unbindHostBadRequest  %+v", 400, o.Payload)
}

func (o *UnbindHostBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *UnbindHostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUnbindHostUnauthorized creates a UnbindHostUnauthorized with default headers values
func NewUnbindHostUnauthorized() *UnbindHostUnauthorized {
	return &UnbindHostUnauthorized{}
}

/*
UnbindHostUnauthorized describes a response with status code 401, with default header values.

Unauthorized.
*/
type UnbindHostUnauthorized struct {
	Payload *models.InfraError
}

// IsSuccess returns true when this unbind host unauthorized response has a 2xx status code
func (o *UnbindHostUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this unbind host unauthorized response has a 3xx status code
func (o *UnbindHostUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this unbind host unauthorized response has a 4xx status code
func (o *UnbindHostUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this unbind host unauthorized response has a 5xx status code
func (o *UnbindHostUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this unbind host unauthorized response a status code equal to that given
func (o *UnbindHostUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *UnbindHostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v2/infra-envs/{infra_env_id}/hosts/{host_id}/actions/unbind][%d] unbindHostUnauthorized  %+v", 401, o.Payload)
}

func (o *UnbindHostUnauthorized) String() string {
	return fmt.Sprintf("[POST /v2/infra-envs/{infra_env_id}/hosts/{host_id}/actions/unbind][%d] unbindHostUnauthorized  %+v", 401, o.Payload)
}

func (o *UnbindHostUnauthorized) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *UnbindHostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUnbindHostForbidden creates a UnbindHostForbidden with default headers values
func NewUnbindHostForbidden() *UnbindHostForbidden {
	return &UnbindHostForbidden{}
}

/*
UnbindHostForbidden describes a response with status code 403, with default header values.

Forbidden.
*/
type UnbindHostForbidden struct {
	Payload *models.InfraError
}

// IsSuccess returns true when this unbind host forbidden response has a 2xx status code
func (o *UnbindHostForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this unbind host forbidden response has a 3xx status code
func (o *UnbindHostForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this unbind host forbidden response has a 4xx status code
func (o *UnbindHostForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this unbind host forbidden response has a 5xx status code
func (o *UnbindHostForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this unbind host forbidden response a status code equal to that given
func (o *UnbindHostForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UnbindHostForbidden) Error() string {
	return fmt.Sprintf("[POST /v2/infra-envs/{infra_env_id}/hosts/{host_id}/actions/unbind][%d] unbindHostForbidden  %+v", 403, o.Payload)
}

func (o *UnbindHostForbidden) String() string {
	return fmt.Sprintf("[POST /v2/infra-envs/{infra_env_id}/hosts/{host_id}/actions/unbind][%d] unbindHostForbidden  %+v", 403, o.Payload)
}

func (o *UnbindHostForbidden) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *UnbindHostForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUnbindHostNotFound creates a UnbindHostNotFound with default headers values
func NewUnbindHostNotFound() *UnbindHostNotFound {
	return &UnbindHostNotFound{}
}

/*
UnbindHostNotFound describes a response with status code 404, with default header values.

Error.
*/
type UnbindHostNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this unbind host not found response has a 2xx status code
func (o *UnbindHostNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this unbind host not found response has a 3xx status code
func (o *UnbindHostNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this unbind host not found response has a 4xx status code
func (o *UnbindHostNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this unbind host not found response has a 5xx status code
func (o *UnbindHostNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this unbind host not found response a status code equal to that given
func (o *UnbindHostNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *UnbindHostNotFound) Error() string {
	return fmt.Sprintf("[POST /v2/infra-envs/{infra_env_id}/hosts/{host_id}/actions/unbind][%d] unbindHostNotFound  %+v", 404, o.Payload)
}

func (o *UnbindHostNotFound) String() string {
	return fmt.Sprintf("[POST /v2/infra-envs/{infra_env_id}/hosts/{host_id}/actions/unbind][%d] unbindHostNotFound  %+v", 404, o.Payload)
}

func (o *UnbindHostNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UnbindHostNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUnbindHostMethodNotAllowed creates a UnbindHostMethodNotAllowed with default headers values
func NewUnbindHostMethodNotAllowed() *UnbindHostMethodNotAllowed {
	return &UnbindHostMethodNotAllowed{}
}

/*
UnbindHostMethodNotAllowed describes a response with status code 405, with default header values.

Method Not Allowed.
*/
type UnbindHostMethodNotAllowed struct {
	Payload *models.Error
}

// IsSuccess returns true when this unbind host method not allowed response has a 2xx status code
func (o *UnbindHostMethodNotAllowed) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this unbind host method not allowed response has a 3xx status code
func (o *UnbindHostMethodNotAllowed) IsRedirect() bool {
	return false
}

// IsClientError returns true when this unbind host method not allowed response has a 4xx status code
func (o *UnbindHostMethodNotAllowed) IsClientError() bool {
	return true
}

// IsServerError returns true when this unbind host method not allowed response has a 5xx status code
func (o *UnbindHostMethodNotAllowed) IsServerError() bool {
	return false
}

// IsCode returns true when this unbind host method not allowed response a status code equal to that given
func (o *UnbindHostMethodNotAllowed) IsCode(code int) bool {
	return code == 405
}

func (o *UnbindHostMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /v2/infra-envs/{infra_env_id}/hosts/{host_id}/actions/unbind][%d] unbindHostMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *UnbindHostMethodNotAllowed) String() string {
	return fmt.Sprintf("[POST /v2/infra-envs/{infra_env_id}/hosts/{host_id}/actions/unbind][%d] unbindHostMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *UnbindHostMethodNotAllowed) GetPayload() *models.Error {
	return o.Payload
}

func (o *UnbindHostMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUnbindHostConflict creates a UnbindHostConflict with default headers values
func NewUnbindHostConflict() *UnbindHostConflict {
	return &UnbindHostConflict{}
}

/*
UnbindHostConflict describes a response with status code 409, with default header values.

Conflict.
*/
type UnbindHostConflict struct {
	Payload *models.Error
}

// IsSuccess returns true when this unbind host conflict response has a 2xx status code
func (o *UnbindHostConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this unbind host conflict response has a 3xx status code
func (o *UnbindHostConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this unbind host conflict response has a 4xx status code
func (o *UnbindHostConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this unbind host conflict response has a 5xx status code
func (o *UnbindHostConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this unbind host conflict response a status code equal to that given
func (o *UnbindHostConflict) IsCode(code int) bool {
	return code == 409
}

func (o *UnbindHostConflict) Error() string {
	return fmt.Sprintf("[POST /v2/infra-envs/{infra_env_id}/hosts/{host_id}/actions/unbind][%d] unbindHostConflict  %+v", 409, o.Payload)
}

func (o *UnbindHostConflict) String() string {
	return fmt.Sprintf("[POST /v2/infra-envs/{infra_env_id}/hosts/{host_id}/actions/unbind][%d] unbindHostConflict  %+v", 409, o.Payload)
}

func (o *UnbindHostConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *UnbindHostConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUnbindHostInternalServerError creates a UnbindHostInternalServerError with default headers values
func NewUnbindHostInternalServerError() *UnbindHostInternalServerError {
	return &UnbindHostInternalServerError{}
}

/*
UnbindHostInternalServerError describes a response with status code 500, with default header values.

Error.
*/
type UnbindHostInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this unbind host internal server error response has a 2xx status code
func (o *UnbindHostInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this unbind host internal server error response has a 3xx status code
func (o *UnbindHostInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this unbind host internal server error response has a 4xx status code
func (o *UnbindHostInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this unbind host internal server error response has a 5xx status code
func (o *UnbindHostInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this unbind host internal server error response a status code equal to that given
func (o *UnbindHostInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *UnbindHostInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v2/infra-envs/{infra_env_id}/hosts/{host_id}/actions/unbind][%d] unbindHostInternalServerError  %+v", 500, o.Payload)
}

func (o *UnbindHostInternalServerError) String() string {
	return fmt.Sprintf("[POST /v2/infra-envs/{infra_env_id}/hosts/{host_id}/actions/unbind][%d] unbindHostInternalServerError  %+v", 500, o.Payload)
}

func (o *UnbindHostInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *UnbindHostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUnbindHostNotImplemented creates a UnbindHostNotImplemented with default headers values
func NewUnbindHostNotImplemented() *UnbindHostNotImplemented {
	return &UnbindHostNotImplemented{}
}

/*
UnbindHostNotImplemented describes a response with status code 501, with default header values.

Not implemented.
*/
type UnbindHostNotImplemented struct {
	Payload *models.Error
}

// IsSuccess returns true when this unbind host not implemented response has a 2xx status code
func (o *UnbindHostNotImplemented) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this unbind host not implemented response has a 3xx status code
func (o *UnbindHostNotImplemented) IsRedirect() bool {
	return false
}

// IsClientError returns true when this unbind host not implemented response has a 4xx status code
func (o *UnbindHostNotImplemented) IsClientError() bool {
	return false
}

// IsServerError returns true when this unbind host not implemented response has a 5xx status code
func (o *UnbindHostNotImplemented) IsServerError() bool {
	return true
}

// IsCode returns true when this unbind host not implemented response a status code equal to that given
func (o *UnbindHostNotImplemented) IsCode(code int) bool {
	return code == 501
}

func (o *UnbindHostNotImplemented) Error() string {
	return fmt.Sprintf("[POST /v2/infra-envs/{infra_env_id}/hosts/{host_id}/actions/unbind][%d] unbindHostNotImplemented  %+v", 501, o.Payload)
}

func (o *UnbindHostNotImplemented) String() string {
	return fmt.Sprintf("[POST /v2/infra-envs/{infra_env_id}/hosts/{host_id}/actions/unbind][%d] unbindHostNotImplemented  %+v", 501, o.Payload)
}

func (o *UnbindHostNotImplemented) GetPayload() *models.Error {
	return o.Payload
}

func (o *UnbindHostNotImplemented) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUnbindHostServiceUnavailable creates a UnbindHostServiceUnavailable with default headers values
func NewUnbindHostServiceUnavailable() *UnbindHostServiceUnavailable {
	return &UnbindHostServiceUnavailable{}
}

/*
UnbindHostServiceUnavailable describes a response with status code 503, with default header values.

Unavailable.
*/
type UnbindHostServiceUnavailable struct {
	Payload *models.Error
}

// IsSuccess returns true when this unbind host service unavailable response has a 2xx status code
func (o *UnbindHostServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this unbind host service unavailable response has a 3xx status code
func (o *UnbindHostServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this unbind host service unavailable response has a 4xx status code
func (o *UnbindHostServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this unbind host service unavailable response has a 5xx status code
func (o *UnbindHostServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this unbind host service unavailable response a status code equal to that given
func (o *UnbindHostServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *UnbindHostServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /v2/infra-envs/{infra_env_id}/hosts/{host_id}/actions/unbind][%d] unbindHostServiceUnavailable  %+v", 503, o.Payload)
}

func (o *UnbindHostServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /v2/infra-envs/{infra_env_id}/hosts/{host_id}/actions/unbind][%d] unbindHostServiceUnavailable  %+v", 503, o.Payload)
}

func (o *UnbindHostServiceUnavailable) GetPayload() *models.Error {
	return o.Payload
}

func (o *UnbindHostServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
