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

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// UpdateOrganizationReader is a Reader for the UpdateOrganization structure.
type UpdateOrganizationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateOrganizationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateOrganizationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateOrganizationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateOrganizationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateOrganizationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateOrganizationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateOrganizationOK creates a UpdateOrganizationOK with default headers values
func NewUpdateOrganizationOK() *UpdateOrganizationOK {
	return &UpdateOrganizationOK{}
}

/*
UpdateOrganizationOK describes a response with status code 200, with default header values.

Organization updated successfully
*/
type UpdateOrganizationOK struct {
	Payload *models.Organization
}

// IsSuccess returns true when this update organization o k response has a 2xx status code
func (o *UpdateOrganizationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update organization o k response has a 3xx status code
func (o *UpdateOrganizationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update organization o k response has a 4xx status code
func (o *UpdateOrganizationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update organization o k response has a 5xx status code
func (o *UpdateOrganizationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update organization o k response a status code equal to that given
func (o *UpdateOrganizationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update organization o k response
func (o *UpdateOrganizationOK) Code() int {
	return 200
}

func (o *UpdateOrganizationOK) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organization_id}][%d] updateOrganizationOK  %+v", 200, o.Payload)
}

func (o *UpdateOrganizationOK) String() string {
	return fmt.Sprintf("[PUT /organizations/{organization_id}][%d] updateOrganizationOK  %+v", 200, o.Payload)
}

func (o *UpdateOrganizationOK) GetPayload() *models.Organization {
	return o.Payload
}

func (o *UpdateOrganizationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Organization)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateOrganizationBadRequest creates a UpdateOrganizationBadRequest with default headers values
func NewUpdateOrganizationBadRequest() *UpdateOrganizationBadRequest {
	return &UpdateOrganizationBadRequest{}
}

/*
	UpdateOrganizationBadRequest describes a response with status code 400, with default header values.

	* Name must be between 2 and 30 characters. (code: `organization.invalid_name`)

* User already has an organization. (code: `organization.user_organization_already_exists`)
*/
type UpdateOrganizationBadRequest struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

// IsSuccess returns true when this update organization bad request response has a 2xx status code
func (o *UpdateOrganizationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update organization bad request response has a 3xx status code
func (o *UpdateOrganizationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update organization bad request response has a 4xx status code
func (o *UpdateOrganizationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update organization bad request response has a 5xx status code
func (o *UpdateOrganizationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update organization bad request response a status code equal to that given
func (o *UpdateOrganizationBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update organization bad request response
func (o *UpdateOrganizationBadRequest) Code() int {
	return 400
}

func (o *UpdateOrganizationBadRequest) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organization_id}][%d] updateOrganizationBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateOrganizationBadRequest) String() string {
	return fmt.Sprintf("[PUT /organizations/{organization_id}][%d] updateOrganizationBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateOrganizationBadRequest) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *UpdateOrganizationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateOrganizationUnauthorized creates a UpdateOrganizationUnauthorized with default headers values
func NewUpdateOrganizationUnauthorized() *UpdateOrganizationUnauthorized {
	return &UpdateOrganizationUnauthorized{}
}

/*
UpdateOrganizationUnauthorized describes a response with status code 401, with default header values.

You are not authorized to perform this action
*/
type UpdateOrganizationUnauthorized struct {
	Payload *models.BasicFailedReply
}

// IsSuccess returns true when this update organization unauthorized response has a 2xx status code
func (o *UpdateOrganizationUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update organization unauthorized response has a 3xx status code
func (o *UpdateOrganizationUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update organization unauthorized response has a 4xx status code
func (o *UpdateOrganizationUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update organization unauthorized response has a 5xx status code
func (o *UpdateOrganizationUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update organization unauthorized response a status code equal to that given
func (o *UpdateOrganizationUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the update organization unauthorized response
func (o *UpdateOrganizationUnauthorized) Code() int {
	return 401
}

func (o *UpdateOrganizationUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organization_id}][%d] updateOrganizationUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateOrganizationUnauthorized) String() string {
	return fmt.Sprintf("[PUT /organizations/{organization_id}][%d] updateOrganizationUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateOrganizationUnauthorized) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *UpdateOrganizationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateOrganizationForbidden creates a UpdateOrganizationForbidden with default headers values
func NewUpdateOrganizationForbidden() *UpdateOrganizationForbidden {
	return &UpdateOrganizationForbidden{}
}

/*
UpdateOrganizationForbidden describes a response with status code 403, with default header values.

The current user does not have access to the requested organization. (code: `organization.invalid_access`)
*/
type UpdateOrganizationForbidden struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

// IsSuccess returns true when this update organization forbidden response has a 2xx status code
func (o *UpdateOrganizationForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update organization forbidden response has a 3xx status code
func (o *UpdateOrganizationForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update organization forbidden response has a 4xx status code
func (o *UpdateOrganizationForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update organization forbidden response has a 5xx status code
func (o *UpdateOrganizationForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update organization forbidden response a status code equal to that given
func (o *UpdateOrganizationForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update organization forbidden response
func (o *UpdateOrganizationForbidden) Code() int {
	return 403
}

func (o *UpdateOrganizationForbidden) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organization_id}][%d] updateOrganizationForbidden  %+v", 403, o.Payload)
}

func (o *UpdateOrganizationForbidden) String() string {
	return fmt.Sprintf("[PUT /organizations/{organization_id}][%d] updateOrganizationForbidden  %+v", 403, o.Payload)
}

func (o *UpdateOrganizationForbidden) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *UpdateOrganizationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateOrganizationNotFound creates a UpdateOrganizationNotFound with default headers values
func NewUpdateOrganizationNotFound() *UpdateOrganizationNotFound {
	return &UpdateOrganizationNotFound{}
}

/*
UpdateOrganizationNotFound describes a response with status code 404, with default header values.

Organization not found. (code: `organization.not_found`)
*/
type UpdateOrganizationNotFound struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

// IsSuccess returns true when this update organization not found response has a 2xx status code
func (o *UpdateOrganizationNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update organization not found response has a 3xx status code
func (o *UpdateOrganizationNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update organization not found response has a 4xx status code
func (o *UpdateOrganizationNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update organization not found response has a 5xx status code
func (o *UpdateOrganizationNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update organization not found response a status code equal to that given
func (o *UpdateOrganizationNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update organization not found response
func (o *UpdateOrganizationNotFound) Code() int {
	return 404
}

func (o *UpdateOrganizationNotFound) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organization_id}][%d] updateOrganizationNotFound  %+v", 404, o.Payload)
}

func (o *UpdateOrganizationNotFound) String() string {
	return fmt.Sprintf("[PUT /organizations/{organization_id}][%d] updateOrganizationNotFound  %+v", 404, o.Payload)
}

func (o *UpdateOrganizationNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *UpdateOrganizationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
