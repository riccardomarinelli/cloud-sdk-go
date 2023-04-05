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

package platform_configuration_snapshots

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

// NewGetSnapshotRepositoriesParams creates a new GetSnapshotRepositoriesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSnapshotRepositoriesParams() *GetSnapshotRepositoriesParams {
	return &GetSnapshotRepositoriesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSnapshotRepositoriesParamsWithTimeout creates a new GetSnapshotRepositoriesParams object
// with the ability to set a timeout on a request.
func NewGetSnapshotRepositoriesParamsWithTimeout(timeout time.Duration) *GetSnapshotRepositoriesParams {
	return &GetSnapshotRepositoriesParams{
		timeout: timeout,
	}
}

// NewGetSnapshotRepositoriesParamsWithContext creates a new GetSnapshotRepositoriesParams object
// with the ability to set a context for a request.
func NewGetSnapshotRepositoriesParamsWithContext(ctx context.Context) *GetSnapshotRepositoriesParams {
	return &GetSnapshotRepositoriesParams{
		Context: ctx,
	}
}

// NewGetSnapshotRepositoriesParamsWithHTTPClient creates a new GetSnapshotRepositoriesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSnapshotRepositoriesParamsWithHTTPClient(client *http.Client) *GetSnapshotRepositoriesParams {
	return &GetSnapshotRepositoriesParams{
		HTTPClient: client,
	}
}

/*
GetSnapshotRepositoriesParams contains all the parameters to send to the API endpoint

	for the get snapshot repositories operation.

	Typically these are written to a http.Request.
*/
type GetSnapshotRepositoriesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get snapshot repositories params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSnapshotRepositoriesParams) WithDefaults() *GetSnapshotRepositoriesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get snapshot repositories params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSnapshotRepositoriesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get snapshot repositories params
func (o *GetSnapshotRepositoriesParams) WithTimeout(timeout time.Duration) *GetSnapshotRepositoriesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get snapshot repositories params
func (o *GetSnapshotRepositoriesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get snapshot repositories params
func (o *GetSnapshotRepositoriesParams) WithContext(ctx context.Context) *GetSnapshotRepositoriesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get snapshot repositories params
func (o *GetSnapshotRepositoriesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get snapshot repositories params
func (o *GetSnapshotRepositoriesParams) WithHTTPClient(client *http.Client) *GetSnapshotRepositoriesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get snapshot repositories params
func (o *GetSnapshotRepositoriesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetSnapshotRepositoriesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
