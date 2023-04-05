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

package platform_infrastructure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// SetRunnerLoggingSettingsReader is a Reader for the SetRunnerLoggingSettings structure.
type SetRunnerLoggingSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetRunnerLoggingSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSetRunnerLoggingSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewSetRunnerLoggingSettingsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSetRunnerLoggingSettingsOK creates a SetRunnerLoggingSettingsOK with default headers values
func NewSetRunnerLoggingSettingsOK() *SetRunnerLoggingSettingsOK {
	return &SetRunnerLoggingSettingsOK{}
}

/*
SetRunnerLoggingSettingsOK describes a response with status code 200, with default header values.

The updated logging settings for the runner instance
*/
type SetRunnerLoggingSettingsOK struct {

	/* The date-time when the resource was created (ISO format relative to UTC)
	 */
	XCloudResourceCreated string

	/* The date-time when the resource was last modified (ISO format relative to UTC)
	 */
	XCloudResourceLastModified string

	/* The resource version, which is used to avoid update conflicts with concurrent operations
	 */
	XCloudResourceVersion string

	Payload *models.LoggingSettings
}

// IsSuccess returns true when this set runner logging settings o k response has a 2xx status code
func (o *SetRunnerLoggingSettingsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this set runner logging settings o k response has a 3xx status code
func (o *SetRunnerLoggingSettingsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this set runner logging settings o k response has a 4xx status code
func (o *SetRunnerLoggingSettingsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this set runner logging settings o k response has a 5xx status code
func (o *SetRunnerLoggingSettingsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this set runner logging settings o k response a status code equal to that given
func (o *SetRunnerLoggingSettingsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the set runner logging settings o k response
func (o *SetRunnerLoggingSettingsOK) Code() int {
	return 200
}

func (o *SetRunnerLoggingSettingsOK) Error() string {
	return fmt.Sprintf("[PUT /platform/infrastructure/runners/{runner_id}/logging_settings][%d] setRunnerLoggingSettingsOK  %+v", 200, o.Payload)
}

func (o *SetRunnerLoggingSettingsOK) String() string {
	return fmt.Sprintf("[PUT /platform/infrastructure/runners/{runner_id}/logging_settings][%d] setRunnerLoggingSettingsOK  %+v", 200, o.Payload)
}

func (o *SetRunnerLoggingSettingsOK) GetPayload() *models.LoggingSettings {
	return o.Payload
}

func (o *SetRunnerLoggingSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-cloud-resource-created
	hdrXCloudResourceCreated := response.GetHeader("x-cloud-resource-created")

	if hdrXCloudResourceCreated != "" {
		o.XCloudResourceCreated = hdrXCloudResourceCreated
	}

	// hydrates response header x-cloud-resource-last-modified
	hdrXCloudResourceLastModified := response.GetHeader("x-cloud-resource-last-modified")

	if hdrXCloudResourceLastModified != "" {
		o.XCloudResourceLastModified = hdrXCloudResourceLastModified
	}

	// hydrates response header x-cloud-resource-version
	hdrXCloudResourceVersion := response.GetHeader("x-cloud-resource-version")

	if hdrXCloudResourceVersion != "" {
		o.XCloudResourceVersion = hdrXCloudResourceVersion
	}

	o.Payload = new(models.LoggingSettings)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetRunnerLoggingSettingsNotFound creates a SetRunnerLoggingSettingsNotFound with default headers values
func NewSetRunnerLoggingSettingsNotFound() *SetRunnerLoggingSettingsNotFound {
	return &SetRunnerLoggingSettingsNotFound{}
}

/*
SetRunnerLoggingSettingsNotFound describes a response with status code 404, with default header values.

The logging settings for this runner were not found. (code: `runners.logging_settings.not_found`)
*/
type SetRunnerLoggingSettingsNotFound struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

// IsSuccess returns true when this set runner logging settings not found response has a 2xx status code
func (o *SetRunnerLoggingSettingsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this set runner logging settings not found response has a 3xx status code
func (o *SetRunnerLoggingSettingsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this set runner logging settings not found response has a 4xx status code
func (o *SetRunnerLoggingSettingsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this set runner logging settings not found response has a 5xx status code
func (o *SetRunnerLoggingSettingsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this set runner logging settings not found response a status code equal to that given
func (o *SetRunnerLoggingSettingsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the set runner logging settings not found response
func (o *SetRunnerLoggingSettingsNotFound) Code() int {
	return 404
}

func (o *SetRunnerLoggingSettingsNotFound) Error() string {
	return fmt.Sprintf("[PUT /platform/infrastructure/runners/{runner_id}/logging_settings][%d] setRunnerLoggingSettingsNotFound  %+v", 404, o.Payload)
}

func (o *SetRunnerLoggingSettingsNotFound) String() string {
	return fmt.Sprintf("[PUT /platform/infrastructure/runners/{runner_id}/logging_settings][%d] setRunnerLoggingSettingsNotFound  %+v", 404, o.Payload)
}

func (o *SetRunnerLoggingSettingsNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *SetRunnerLoggingSettingsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-cloud-error-codes
	hdrXCloudErrorCodes := response.GetHeader("x-cloud-error-codes")

	if hdrXCloudErrorCodes != "" {
		o.XCloudErrorCodes = hdrXCloudErrorCodes
	}

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
