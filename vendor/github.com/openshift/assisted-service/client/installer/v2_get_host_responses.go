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

// V2GetHostReader is a Reader for the V2GetHost structure.
type V2GetHostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V2GetHostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV2GetHostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewV2GetHostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewV2GetHostForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewV2GetHostNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewV2GetHostMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewV2GetHostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 501:
		result := NewV2GetHostNotImplemented()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV2GetHostOK creates a V2GetHostOK with default headers values
func NewV2GetHostOK() *V2GetHostOK {
	return &V2GetHostOK{}
}

/*
V2GetHostOK describes a response with status code 200, with default header values.

Success.
*/
type V2GetHostOK struct {
	Payload *models.Host
}

// IsSuccess returns true when this v2 get host o k response has a 2xx status code
func (o *V2GetHostOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this v2 get host o k response has a 3xx status code
func (o *V2GetHostOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v2 get host o k response has a 4xx status code
func (o *V2GetHostOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this v2 get host o k response has a 5xx status code
func (o *V2GetHostOK) IsServerError() bool {
	return false
}

// IsCode returns true when this v2 get host o k response a status code equal to that given
func (o *V2GetHostOK) IsCode(code int) bool {
	return code == 200
}

func (o *V2GetHostOK) Error() string {
	return fmt.Sprintf("[GET /v2/infra-envs/{infra_env_id}/hosts/{host_id}][%d] v2GetHostOK  %+v", 200, o.Payload)
}

func (o *V2GetHostOK) String() string {
	return fmt.Sprintf("[GET /v2/infra-envs/{infra_env_id}/hosts/{host_id}][%d] v2GetHostOK  %+v", 200, o.Payload)
}

func (o *V2GetHostOK) GetPayload() *models.Host {
	return o.Payload
}

func (o *V2GetHostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Host)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2GetHostUnauthorized creates a V2GetHostUnauthorized with default headers values
func NewV2GetHostUnauthorized() *V2GetHostUnauthorized {
	return &V2GetHostUnauthorized{}
}

/*
V2GetHostUnauthorized describes a response with status code 401, with default header values.

Unauthorized.
*/
type V2GetHostUnauthorized struct {
	Payload *models.InfraError
}

// IsSuccess returns true when this v2 get host unauthorized response has a 2xx status code
func (o *V2GetHostUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v2 get host unauthorized response has a 3xx status code
func (o *V2GetHostUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v2 get host unauthorized response has a 4xx status code
func (o *V2GetHostUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this v2 get host unauthorized response has a 5xx status code
func (o *V2GetHostUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this v2 get host unauthorized response a status code equal to that given
func (o *V2GetHostUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *V2GetHostUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v2/infra-envs/{infra_env_id}/hosts/{host_id}][%d] v2GetHostUnauthorized  %+v", 401, o.Payload)
}

func (o *V2GetHostUnauthorized) String() string {
	return fmt.Sprintf("[GET /v2/infra-envs/{infra_env_id}/hosts/{host_id}][%d] v2GetHostUnauthorized  %+v", 401, o.Payload)
}

func (o *V2GetHostUnauthorized) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *V2GetHostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2GetHostForbidden creates a V2GetHostForbidden with default headers values
func NewV2GetHostForbidden() *V2GetHostForbidden {
	return &V2GetHostForbidden{}
}

/*
V2GetHostForbidden describes a response with status code 403, with default header values.

Forbidden.
*/
type V2GetHostForbidden struct {
	Payload *models.InfraError
}

// IsSuccess returns true when this v2 get host forbidden response has a 2xx status code
func (o *V2GetHostForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v2 get host forbidden response has a 3xx status code
func (o *V2GetHostForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v2 get host forbidden response has a 4xx status code
func (o *V2GetHostForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this v2 get host forbidden response has a 5xx status code
func (o *V2GetHostForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this v2 get host forbidden response a status code equal to that given
func (o *V2GetHostForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *V2GetHostForbidden) Error() string {
	return fmt.Sprintf("[GET /v2/infra-envs/{infra_env_id}/hosts/{host_id}][%d] v2GetHostForbidden  %+v", 403, o.Payload)
}

func (o *V2GetHostForbidden) String() string {
	return fmt.Sprintf("[GET /v2/infra-envs/{infra_env_id}/hosts/{host_id}][%d] v2GetHostForbidden  %+v", 403, o.Payload)
}

func (o *V2GetHostForbidden) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *V2GetHostForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2GetHostNotFound creates a V2GetHostNotFound with default headers values
func NewV2GetHostNotFound() *V2GetHostNotFound {
	return &V2GetHostNotFound{}
}

/*
V2GetHostNotFound describes a response with status code 404, with default header values.

Error.
*/
type V2GetHostNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this v2 get host not found response has a 2xx status code
func (o *V2GetHostNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v2 get host not found response has a 3xx status code
func (o *V2GetHostNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v2 get host not found response has a 4xx status code
func (o *V2GetHostNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this v2 get host not found response has a 5xx status code
func (o *V2GetHostNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this v2 get host not found response a status code equal to that given
func (o *V2GetHostNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *V2GetHostNotFound) Error() string {
	return fmt.Sprintf("[GET /v2/infra-envs/{infra_env_id}/hosts/{host_id}][%d] v2GetHostNotFound  %+v", 404, o.Payload)
}

func (o *V2GetHostNotFound) String() string {
	return fmt.Sprintf("[GET /v2/infra-envs/{infra_env_id}/hosts/{host_id}][%d] v2GetHostNotFound  %+v", 404, o.Payload)
}

func (o *V2GetHostNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2GetHostNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2GetHostMethodNotAllowed creates a V2GetHostMethodNotAllowed with default headers values
func NewV2GetHostMethodNotAllowed() *V2GetHostMethodNotAllowed {
	return &V2GetHostMethodNotAllowed{}
}

/*
V2GetHostMethodNotAllowed describes a response with status code 405, with default header values.

Method Not Allowed.
*/
type V2GetHostMethodNotAllowed struct {
	Payload *models.Error
}

// IsSuccess returns true when this v2 get host method not allowed response has a 2xx status code
func (o *V2GetHostMethodNotAllowed) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v2 get host method not allowed response has a 3xx status code
func (o *V2GetHostMethodNotAllowed) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v2 get host method not allowed response has a 4xx status code
func (o *V2GetHostMethodNotAllowed) IsClientError() bool {
	return true
}

// IsServerError returns true when this v2 get host method not allowed response has a 5xx status code
func (o *V2GetHostMethodNotAllowed) IsServerError() bool {
	return false
}

// IsCode returns true when this v2 get host method not allowed response a status code equal to that given
func (o *V2GetHostMethodNotAllowed) IsCode(code int) bool {
	return code == 405
}

func (o *V2GetHostMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /v2/infra-envs/{infra_env_id}/hosts/{host_id}][%d] v2GetHostMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *V2GetHostMethodNotAllowed) String() string {
	return fmt.Sprintf("[GET /v2/infra-envs/{infra_env_id}/hosts/{host_id}][%d] v2GetHostMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *V2GetHostMethodNotAllowed) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2GetHostMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2GetHostInternalServerError creates a V2GetHostInternalServerError with default headers values
func NewV2GetHostInternalServerError() *V2GetHostInternalServerError {
	return &V2GetHostInternalServerError{}
}

/*
V2GetHostInternalServerError describes a response with status code 500, with default header values.

Error.
*/
type V2GetHostInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this v2 get host internal server error response has a 2xx status code
func (o *V2GetHostInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v2 get host internal server error response has a 3xx status code
func (o *V2GetHostInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v2 get host internal server error response has a 4xx status code
func (o *V2GetHostInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this v2 get host internal server error response has a 5xx status code
func (o *V2GetHostInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this v2 get host internal server error response a status code equal to that given
func (o *V2GetHostInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *V2GetHostInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v2/infra-envs/{infra_env_id}/hosts/{host_id}][%d] v2GetHostInternalServerError  %+v", 500, o.Payload)
}

func (o *V2GetHostInternalServerError) String() string {
	return fmt.Sprintf("[GET /v2/infra-envs/{infra_env_id}/hosts/{host_id}][%d] v2GetHostInternalServerError  %+v", 500, o.Payload)
}

func (o *V2GetHostInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2GetHostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2GetHostNotImplemented creates a V2GetHostNotImplemented with default headers values
func NewV2GetHostNotImplemented() *V2GetHostNotImplemented {
	return &V2GetHostNotImplemented{}
}

/*
V2GetHostNotImplemented describes a response with status code 501, with default header values.

Not implemented.
*/
type V2GetHostNotImplemented struct {
	Payload *models.Error
}

// IsSuccess returns true when this v2 get host not implemented response has a 2xx status code
func (o *V2GetHostNotImplemented) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v2 get host not implemented response has a 3xx status code
func (o *V2GetHostNotImplemented) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v2 get host not implemented response has a 4xx status code
func (o *V2GetHostNotImplemented) IsClientError() bool {
	return false
}

// IsServerError returns true when this v2 get host not implemented response has a 5xx status code
func (o *V2GetHostNotImplemented) IsServerError() bool {
	return true
}

// IsCode returns true when this v2 get host not implemented response a status code equal to that given
func (o *V2GetHostNotImplemented) IsCode(code int) bool {
	return code == 501
}

func (o *V2GetHostNotImplemented) Error() string {
	return fmt.Sprintf("[GET /v2/infra-envs/{infra_env_id}/hosts/{host_id}][%d] v2GetHostNotImplemented  %+v", 501, o.Payload)
}

func (o *V2GetHostNotImplemented) String() string {
	return fmt.Sprintf("[GET /v2/infra-envs/{infra_env_id}/hosts/{host_id}][%d] v2GetHostNotImplemented  %+v", 501, o.Payload)
}

func (o *V2GetHostNotImplemented) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2GetHostNotImplemented) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
