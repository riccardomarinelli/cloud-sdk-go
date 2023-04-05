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

package platform_configuration_security

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// CreateSamlConfigurationReader is a Reader for the CreateSamlConfiguration structure.
type CreateSamlConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateSamlConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateSamlConfigurationCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateSamlConfigurationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 449:
		result := NewCreateSamlConfigurationRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateSamlConfigurationCreated creates a CreateSamlConfigurationCreated with default headers values
func NewCreateSamlConfigurationCreated() *CreateSamlConfigurationCreated {
	return &CreateSamlConfigurationCreated{}
}

/*
CreateSamlConfigurationCreated describes a response with status code 201, with default header values.

The SAML configuration was successfully created
*/
type CreateSamlConfigurationCreated struct {

	/* The date-time when the resource was created (ISO format relative to UTC)
	 */
	XCloudResourceCreated string

	/* The date-time when the resource was last modified (ISO format relative to UTC)
	 */
	XCloudResourceLastModified string

	/* The resource version, which is used to avoid update conflicts with concurrent operations
	 */
	XCloudResourceVersion string

	Payload models.EmptyResponse
}

// IsSuccess returns true when this create saml configuration created response has a 2xx status code
func (o *CreateSamlConfigurationCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create saml configuration created response has a 3xx status code
func (o *CreateSamlConfigurationCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create saml configuration created response has a 4xx status code
func (o *CreateSamlConfigurationCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create saml configuration created response has a 5xx status code
func (o *CreateSamlConfigurationCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create saml configuration created response a status code equal to that given
func (o *CreateSamlConfigurationCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create saml configuration created response
func (o *CreateSamlConfigurationCreated) Code() int {
	return 201
}

func (o *CreateSamlConfigurationCreated) Error() string {
	return fmt.Sprintf("[POST /platform/configuration/security/realms/saml][%d] createSamlConfigurationCreated  %+v", 201, o.Payload)
}

func (o *CreateSamlConfigurationCreated) String() string {
	return fmt.Sprintf("[POST /platform/configuration/security/realms/saml][%d] createSamlConfigurationCreated  %+v", 201, o.Payload)
}

func (o *CreateSamlConfigurationCreated) GetPayload() models.EmptyResponse {
	return o.Payload
}

func (o *CreateSamlConfigurationCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSamlConfigurationBadRequest creates a CreateSamlConfigurationBadRequest with default headers values
func NewCreateSamlConfigurationBadRequest() *CreateSamlConfigurationBadRequest {
	return &CreateSamlConfigurationBadRequest{}
}

/*
	CreateSamlConfigurationBadRequest describes a response with status code 400, with default header values.

	* The realm id is already in use. (code: `security_realm.id_conflict`)

* The selected id is not valid. (code: `security_realm.invalid_id`)
* Order must be greater than zero. (code: `security_realm.invalid_order`)
* Invalid Elasticsearch Security realm type. (code: `security_realm.invalid_type`)
* The realm order is already in use. (code: `security_realm.order_conflict`)
* Advanced YAML format is invalid. (code: `security_realm.invalid_yaml`)
* The SAML IDP metadata endpoint returned an error response code 200 OK. (code: `security_realm.saml.invalid_idp_metadata_url`)
* Invalid certificate bundle URL. (code: `security_realm.invalid_bundle_url`)
*/
type CreateSamlConfigurationBadRequest struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

// IsSuccess returns true when this create saml configuration bad request response has a 2xx status code
func (o *CreateSamlConfigurationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create saml configuration bad request response has a 3xx status code
func (o *CreateSamlConfigurationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create saml configuration bad request response has a 4xx status code
func (o *CreateSamlConfigurationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create saml configuration bad request response has a 5xx status code
func (o *CreateSamlConfigurationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create saml configuration bad request response a status code equal to that given
func (o *CreateSamlConfigurationBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create saml configuration bad request response
func (o *CreateSamlConfigurationBadRequest) Code() int {
	return 400
}

func (o *CreateSamlConfigurationBadRequest) Error() string {
	return fmt.Sprintf("[POST /platform/configuration/security/realms/saml][%d] createSamlConfigurationBadRequest  %+v", 400, o.Payload)
}

func (o *CreateSamlConfigurationBadRequest) String() string {
	return fmt.Sprintf("[POST /platform/configuration/security/realms/saml][%d] createSamlConfigurationBadRequest  %+v", 400, o.Payload)
}

func (o *CreateSamlConfigurationBadRequest) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *CreateSamlConfigurationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateSamlConfigurationRetryWith creates a CreateSamlConfigurationRetryWith with default headers values
func NewCreateSamlConfigurationRetryWith() *CreateSamlConfigurationRetryWith {
	return &CreateSamlConfigurationRetryWith{}
}

/*
CreateSamlConfigurationRetryWith describes a response with status code 449, with default header values.

Elevated permissions are required. (code: `root.unauthorized.rbac.elevated_permissions_required`)
*/
type CreateSamlConfigurationRetryWith struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

// IsSuccess returns true when this create saml configuration retry with response has a 2xx status code
func (o *CreateSamlConfigurationRetryWith) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create saml configuration retry with response has a 3xx status code
func (o *CreateSamlConfigurationRetryWith) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create saml configuration retry with response has a 4xx status code
func (o *CreateSamlConfigurationRetryWith) IsClientError() bool {
	return true
}

// IsServerError returns true when this create saml configuration retry with response has a 5xx status code
func (o *CreateSamlConfigurationRetryWith) IsServerError() bool {
	return false
}

// IsCode returns true when this create saml configuration retry with response a status code equal to that given
func (o *CreateSamlConfigurationRetryWith) IsCode(code int) bool {
	return code == 449
}

// Code gets the status code for the create saml configuration retry with response
func (o *CreateSamlConfigurationRetryWith) Code() int {
	return 449
}

func (o *CreateSamlConfigurationRetryWith) Error() string {
	return fmt.Sprintf("[POST /platform/configuration/security/realms/saml][%d] createSamlConfigurationRetryWith  %+v", 449, o.Payload)
}

func (o *CreateSamlConfigurationRetryWith) String() string {
	return fmt.Sprintf("[POST /platform/configuration/security/realms/saml][%d] createSamlConfigurationRetryWith  %+v", 449, o.Payload)
}

func (o *CreateSamlConfigurationRetryWith) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *CreateSamlConfigurationRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
