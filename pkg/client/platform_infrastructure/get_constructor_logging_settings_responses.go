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

// GetConstructorLoggingSettingsReader is a Reader for the GetConstructorLoggingSettings structure.
type GetConstructorLoggingSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConstructorLoggingSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConstructorLoggingSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetConstructorLoggingSettingsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetConstructorLoggingSettingsOK creates a GetConstructorLoggingSettingsOK with default headers values
func NewGetConstructorLoggingSettingsOK() *GetConstructorLoggingSettingsOK {
	return &GetConstructorLoggingSettingsOK{}
}

/*
GetConstructorLoggingSettingsOK describes a response with status code 200, with default header values.

The logging settings for the constructor instance
*/
type GetConstructorLoggingSettingsOK struct {

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

// IsSuccess returns true when this get constructor logging settings o k response has a 2xx status code
func (o *GetConstructorLoggingSettingsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get constructor logging settings o k response has a 3xx status code
func (o *GetConstructorLoggingSettingsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get constructor logging settings o k response has a 4xx status code
func (o *GetConstructorLoggingSettingsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get constructor logging settings o k response has a 5xx status code
func (o *GetConstructorLoggingSettingsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get constructor logging settings o k response a status code equal to that given
func (o *GetConstructorLoggingSettingsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get constructor logging settings o k response
func (o *GetConstructorLoggingSettingsOK) Code() int {
	return 200
}

func (o *GetConstructorLoggingSettingsOK) Error() string {
	return fmt.Sprintf("[GET /platform/infrastructure/constructors/{constructor_id}/logging_settings][%d] getConstructorLoggingSettingsOK  %+v", 200, o.Payload)
}

func (o *GetConstructorLoggingSettingsOK) String() string {
	return fmt.Sprintf("[GET /platform/infrastructure/constructors/{constructor_id}/logging_settings][%d] getConstructorLoggingSettingsOK  %+v", 200, o.Payload)
}

func (o *GetConstructorLoggingSettingsOK) GetPayload() *models.LoggingSettings {
	return o.Payload
}

func (o *GetConstructorLoggingSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetConstructorLoggingSettingsNotFound creates a GetConstructorLoggingSettingsNotFound with default headers values
func NewGetConstructorLoggingSettingsNotFound() *GetConstructorLoggingSettingsNotFound {
	return &GetConstructorLoggingSettingsNotFound{}
}

/*
GetConstructorLoggingSettingsNotFound describes a response with status code 404, with default header values.

The logging settings for this constructor were not found. (code: `constructors.logging_settings.not_found`)
*/
type GetConstructorLoggingSettingsNotFound struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

// IsSuccess returns true when this get constructor logging settings not found response has a 2xx status code
func (o *GetConstructorLoggingSettingsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get constructor logging settings not found response has a 3xx status code
func (o *GetConstructorLoggingSettingsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get constructor logging settings not found response has a 4xx status code
func (o *GetConstructorLoggingSettingsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get constructor logging settings not found response has a 5xx status code
func (o *GetConstructorLoggingSettingsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get constructor logging settings not found response a status code equal to that given
func (o *GetConstructorLoggingSettingsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get constructor logging settings not found response
func (o *GetConstructorLoggingSettingsNotFound) Code() int {
	return 404
}

func (o *GetConstructorLoggingSettingsNotFound) Error() string {
	return fmt.Sprintf("[GET /platform/infrastructure/constructors/{constructor_id}/logging_settings][%d] getConstructorLoggingSettingsNotFound  %+v", 404, o.Payload)
}

func (o *GetConstructorLoggingSettingsNotFound) String() string {
	return fmt.Sprintf("[GET /platform/infrastructure/constructors/{constructor_id}/logging_settings][%d] getConstructorLoggingSettingsNotFound  %+v", 404, o.Payload)
}

func (o *GetConstructorLoggingSettingsNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *GetConstructorLoggingSettingsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
