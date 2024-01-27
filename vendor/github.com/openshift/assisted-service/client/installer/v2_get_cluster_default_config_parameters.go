// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewV2GetClusterDefaultConfigParams creates a new V2GetClusterDefaultConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewV2GetClusterDefaultConfigParams() *V2GetClusterDefaultConfigParams {
	return &V2GetClusterDefaultConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewV2GetClusterDefaultConfigParamsWithTimeout creates a new V2GetClusterDefaultConfigParams object
// with the ability to set a timeout on a request.
func NewV2GetClusterDefaultConfigParamsWithTimeout(timeout time.Duration) *V2GetClusterDefaultConfigParams {
	return &V2GetClusterDefaultConfigParams{
		timeout: timeout,
	}
}

// NewV2GetClusterDefaultConfigParamsWithContext creates a new V2GetClusterDefaultConfigParams object
// with the ability to set a context for a request.
func NewV2GetClusterDefaultConfigParamsWithContext(ctx context.Context) *V2GetClusterDefaultConfigParams {
	return &V2GetClusterDefaultConfigParams{
		Context: ctx,
	}
}

// NewV2GetClusterDefaultConfigParamsWithHTTPClient creates a new V2GetClusterDefaultConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewV2GetClusterDefaultConfigParamsWithHTTPClient(client *http.Client) *V2GetClusterDefaultConfigParams {
	return &V2GetClusterDefaultConfigParams{
		HTTPClient: client,
	}
}

/*
V2GetClusterDefaultConfigParams contains all the parameters to send to the API endpoint

	for the v2 get cluster default config operation.

	Typically these are written to a http.Request.
*/
type V2GetClusterDefaultConfigParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the v2 get cluster default config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *V2GetClusterDefaultConfigParams) WithDefaults() *V2GetClusterDefaultConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the v2 get cluster default config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *V2GetClusterDefaultConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the v2 get cluster default config params
func (o *V2GetClusterDefaultConfigParams) WithTimeout(timeout time.Duration) *V2GetClusterDefaultConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v2 get cluster default config params
func (o *V2GetClusterDefaultConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v2 get cluster default config params
func (o *V2GetClusterDefaultConfigParams) WithContext(ctx context.Context) *V2GetClusterDefaultConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v2 get cluster default config params
func (o *V2GetClusterDefaultConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v2 get cluster default config params
func (o *V2GetClusterDefaultConfigParams) WithHTTPClient(client *http.Client) *V2GetClusterDefaultConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v2 get cluster default config params
func (o *V2GetClusterDefaultConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *V2GetClusterDefaultConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
