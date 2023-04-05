// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by go-swagger; DO NOT EDIT.

package platform_configuration_templates

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

// NewGetGlobalDeploymentTemplatesParams creates a new GetGlobalDeploymentTemplatesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetGlobalDeploymentTemplatesParams() *GetGlobalDeploymentTemplatesParams {
	return &GetGlobalDeploymentTemplatesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetGlobalDeploymentTemplatesParamsWithTimeout creates a new GetGlobalDeploymentTemplatesParams object
// with the ability to set a timeout on a request.
func NewGetGlobalDeploymentTemplatesParamsWithTimeout(timeout time.Duration) *GetGlobalDeploymentTemplatesParams {
	return &GetGlobalDeploymentTemplatesParams{
		timeout: timeout,
	}
}

// NewGetGlobalDeploymentTemplatesParamsWithContext creates a new GetGlobalDeploymentTemplatesParams object
// with the ability to set a context for a request.
func NewGetGlobalDeploymentTemplatesParamsWithContext(ctx context.Context) *GetGlobalDeploymentTemplatesParams {
	return &GetGlobalDeploymentTemplatesParams{
		Context: ctx,
	}
}

// NewGetGlobalDeploymentTemplatesParamsWithHTTPClient creates a new GetGlobalDeploymentTemplatesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetGlobalDeploymentTemplatesParamsWithHTTPClient(client *http.Client) *GetGlobalDeploymentTemplatesParams {
	return &GetGlobalDeploymentTemplatesParams{
		HTTPClient: client,
	}
}

/*
GetGlobalDeploymentTemplatesParams contains all the parameters to send to the API endpoint

	for the get global deployment templates operation.

	Typically these are written to a http.Request.
*/
type GetGlobalDeploymentTemplatesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get global deployment templates params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetGlobalDeploymentTemplatesParams) WithDefaults() *GetGlobalDeploymentTemplatesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get global deployment templates params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetGlobalDeploymentTemplatesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get global deployment templates params
func (o *GetGlobalDeploymentTemplatesParams) WithTimeout(timeout time.Duration) *GetGlobalDeploymentTemplatesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get global deployment templates params
func (o *GetGlobalDeploymentTemplatesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get global deployment templates params
func (o *GetGlobalDeploymentTemplatesParams) WithContext(ctx context.Context) *GetGlobalDeploymentTemplatesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get global deployment templates params
func (o *GetGlobalDeploymentTemplatesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get global deployment templates params
func (o *GetGlobalDeploymentTemplatesParams) WithHTTPClient(client *http.Client) *GetGlobalDeploymentTemplatesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get global deployment templates params
func (o *GetGlobalDeploymentTemplatesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetGlobalDeploymentTemplatesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
