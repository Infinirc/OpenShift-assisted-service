// Code generated by go-swagger; DO NOT EDIT.

package operators

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"

	"github.com/go-openapi/strfmt"
)

// V2ReportMonitoredOperatorStatusURL generates an URL for the v2 report monitored operator status operation
type V2ReportMonitoredOperatorStatusURL struct {
	ClusterID strfmt.UUID

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *V2ReportMonitoredOperatorStatusURL) WithBasePath(bp string) *V2ReportMonitoredOperatorStatusURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *V2ReportMonitoredOperatorStatusURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *V2ReportMonitoredOperatorStatusURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/v2/clusters/{cluster_id}/monitored-operators"

	clusterID := o.ClusterID.String()
	if clusterID != "" {
		_path = strings.Replace(_path, "{cluster_id}", clusterID, -1)
	} else {
		return nil, errors.New("clusterId is required on V2ReportMonitoredOperatorStatusURL")
	}

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/api/assisted-install"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *V2ReportMonitoredOperatorStatusURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *V2ReportMonitoredOperatorStatusURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *V2ReportMonitoredOperatorStatusURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on V2ReportMonitoredOperatorStatusURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on V2ReportMonitoredOperatorStatusURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *V2ReportMonitoredOperatorStatusURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
