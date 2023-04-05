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

package authentication

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

// NewMethodsParams creates a new MethodsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewMethodsParams() *MethodsParams {
	return &MethodsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewMethodsParamsWithTimeout creates a new MethodsParams object
// with the ability to set a timeout on a request.
func NewMethodsParamsWithTimeout(timeout time.Duration) *MethodsParams {
	return &MethodsParams{
		timeout: timeout,
	}
}

// NewMethodsParamsWithContext creates a new MethodsParams object
// with the ability to set a context for a request.
func NewMethodsParamsWithContext(ctx context.Context) *MethodsParams {
	return &MethodsParams{
		Context: ctx,
	}
}

// NewMethodsParamsWithHTTPClient creates a new MethodsParams object
// with the ability to set a custom HTTPClient for a request.
func NewMethodsParamsWithHTTPClient(client *http.Client) *MethodsParams {
	return &MethodsParams{
		HTTPClient: client,
	}
}

/*
MethodsParams contains all the parameters to send to the API endpoint

	for the methods operation.

	Typically these are written to a http.Request.
*/
type MethodsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the methods params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MethodsParams) WithDefaults() *MethodsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the methods params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MethodsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the methods params
func (o *MethodsParams) WithTimeout(timeout time.Duration) *MethodsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the methods params
func (o *MethodsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the methods params
func (o *MethodsParams) WithContext(ctx context.Context) *MethodsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the methods params
func (o *MethodsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the methods params
func (o *MethodsParams) WithHTTPClient(client *http.Client) *MethodsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the methods params
func (o *MethodsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *MethodsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
