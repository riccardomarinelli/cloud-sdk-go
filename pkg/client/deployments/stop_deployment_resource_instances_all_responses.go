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

package deployments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// StopDeploymentResourceInstancesAllReader is a Reader for the StopDeploymentResourceInstancesAll structure.
type StopDeploymentResourceInstancesAllReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StopDeploymentResourceInstancesAllReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewStopDeploymentResourceInstancesAllAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewStopDeploymentResourceInstancesAllForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStopDeploymentResourceInstancesAllNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 449:
		result := NewStopDeploymentResourceInstancesAllRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStopDeploymentResourceInstancesAllInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStopDeploymentResourceInstancesAllAccepted creates a StopDeploymentResourceInstancesAllAccepted with default headers values
func NewStopDeploymentResourceInstancesAllAccepted() *StopDeploymentResourceInstancesAllAccepted {
	return &StopDeploymentResourceInstancesAllAccepted{}
}

/*
StopDeploymentResourceInstancesAllAccepted describes a response with status code 202, with default header values.

The stop command was issued successfully.
*/
type StopDeploymentResourceInstancesAllAccepted struct {
	Payload *models.DeploymentResourceCommandResponse
}

// IsSuccess returns true when this stop deployment resource instances all accepted response has a 2xx status code
func (o *StopDeploymentResourceInstancesAllAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this stop deployment resource instances all accepted response has a 3xx status code
func (o *StopDeploymentResourceInstancesAllAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stop deployment resource instances all accepted response has a 4xx status code
func (o *StopDeploymentResourceInstancesAllAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this stop deployment resource instances all accepted response has a 5xx status code
func (o *StopDeploymentResourceInstancesAllAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this stop deployment resource instances all accepted response a status code equal to that given
func (o *StopDeploymentResourceInstancesAllAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the stop deployment resource instances all accepted response
func (o *StopDeploymentResourceInstancesAllAccepted) Code() int {
	return 202
}

func (o *StopDeploymentResourceInstancesAllAccepted) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/{resource_kind}/{ref_id}/instances/_stop][%d] stopDeploymentResourceInstancesAllAccepted  %+v", 202, o.Payload)
}

func (o *StopDeploymentResourceInstancesAllAccepted) String() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/{resource_kind}/{ref_id}/instances/_stop][%d] stopDeploymentResourceInstancesAllAccepted  %+v", 202, o.Payload)
}

func (o *StopDeploymentResourceInstancesAllAccepted) GetPayload() *models.DeploymentResourceCommandResponse {
	return o.Payload
}

func (o *StopDeploymentResourceInstancesAllAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeploymentResourceCommandResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStopDeploymentResourceInstancesAllForbidden creates a StopDeploymentResourceInstancesAllForbidden with default headers values
func NewStopDeploymentResourceInstancesAllForbidden() *StopDeploymentResourceInstancesAllForbidden {
	return &StopDeploymentResourceInstancesAllForbidden{}
}

/*
StopDeploymentResourceInstancesAllForbidden describes a response with status code 403, with default header values.

The stop maintenance mode command was prohibited for the given Resource. (code: `deployments.instance_update_prohibited_error`)
*/
type StopDeploymentResourceInstancesAllForbidden struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

// IsSuccess returns true when this stop deployment resource instances all forbidden response has a 2xx status code
func (o *StopDeploymentResourceInstancesAllForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stop deployment resource instances all forbidden response has a 3xx status code
func (o *StopDeploymentResourceInstancesAllForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stop deployment resource instances all forbidden response has a 4xx status code
func (o *StopDeploymentResourceInstancesAllForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this stop deployment resource instances all forbidden response has a 5xx status code
func (o *StopDeploymentResourceInstancesAllForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this stop deployment resource instances all forbidden response a status code equal to that given
func (o *StopDeploymentResourceInstancesAllForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the stop deployment resource instances all forbidden response
func (o *StopDeploymentResourceInstancesAllForbidden) Code() int {
	return 403
}

func (o *StopDeploymentResourceInstancesAllForbidden) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/{resource_kind}/{ref_id}/instances/_stop][%d] stopDeploymentResourceInstancesAllForbidden  %+v", 403, o.Payload)
}

func (o *StopDeploymentResourceInstancesAllForbidden) String() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/{resource_kind}/{ref_id}/instances/_stop][%d] stopDeploymentResourceInstancesAllForbidden  %+v", 403, o.Payload)
}

func (o *StopDeploymentResourceInstancesAllForbidden) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *StopDeploymentResourceInstancesAllForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewStopDeploymentResourceInstancesAllNotFound creates a StopDeploymentResourceInstancesAllNotFound with default headers values
func NewStopDeploymentResourceInstancesAllNotFound() *StopDeploymentResourceInstancesAllNotFound {
	return &StopDeploymentResourceInstancesAllNotFound{}
}

/*
	StopDeploymentResourceInstancesAllNotFound describes a response with status code 404, with default header values.

	* The Deployment specified by {deployment_id} cannot be found. (code: `deployments.deployment_not_found`)

* The Resource specified by {ref_id} cannot be found. (code: `deployments.deployment_resource_not_found`)
* One or more instances of the given resource type are missing. (code: `deployments.instances_missing_on_update_error`)
*/
type StopDeploymentResourceInstancesAllNotFound struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

// IsSuccess returns true when this stop deployment resource instances all not found response has a 2xx status code
func (o *StopDeploymentResourceInstancesAllNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stop deployment resource instances all not found response has a 3xx status code
func (o *StopDeploymentResourceInstancesAllNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stop deployment resource instances all not found response has a 4xx status code
func (o *StopDeploymentResourceInstancesAllNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this stop deployment resource instances all not found response has a 5xx status code
func (o *StopDeploymentResourceInstancesAllNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this stop deployment resource instances all not found response a status code equal to that given
func (o *StopDeploymentResourceInstancesAllNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the stop deployment resource instances all not found response
func (o *StopDeploymentResourceInstancesAllNotFound) Code() int {
	return 404
}

func (o *StopDeploymentResourceInstancesAllNotFound) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/{resource_kind}/{ref_id}/instances/_stop][%d] stopDeploymentResourceInstancesAllNotFound  %+v", 404, o.Payload)
}

func (o *StopDeploymentResourceInstancesAllNotFound) String() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/{resource_kind}/{ref_id}/instances/_stop][%d] stopDeploymentResourceInstancesAllNotFound  %+v", 404, o.Payload)
}

func (o *StopDeploymentResourceInstancesAllNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *StopDeploymentResourceInstancesAllNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewStopDeploymentResourceInstancesAllRetryWith creates a StopDeploymentResourceInstancesAllRetryWith with default headers values
func NewStopDeploymentResourceInstancesAllRetryWith() *StopDeploymentResourceInstancesAllRetryWith {
	return &StopDeploymentResourceInstancesAllRetryWith{}
}

/*
StopDeploymentResourceInstancesAllRetryWith describes a response with status code 449, with default header values.

Elevated permissions are required. (code: `root.unauthorized.rbac.elevated_permissions_required`)
*/
type StopDeploymentResourceInstancesAllRetryWith struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

// IsSuccess returns true when this stop deployment resource instances all retry with response has a 2xx status code
func (o *StopDeploymentResourceInstancesAllRetryWith) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stop deployment resource instances all retry with response has a 3xx status code
func (o *StopDeploymentResourceInstancesAllRetryWith) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stop deployment resource instances all retry with response has a 4xx status code
func (o *StopDeploymentResourceInstancesAllRetryWith) IsClientError() bool {
	return true
}

// IsServerError returns true when this stop deployment resource instances all retry with response has a 5xx status code
func (o *StopDeploymentResourceInstancesAllRetryWith) IsServerError() bool {
	return false
}

// IsCode returns true when this stop deployment resource instances all retry with response a status code equal to that given
func (o *StopDeploymentResourceInstancesAllRetryWith) IsCode(code int) bool {
	return code == 449
}

// Code gets the status code for the stop deployment resource instances all retry with response
func (o *StopDeploymentResourceInstancesAllRetryWith) Code() int {
	return 449
}

func (o *StopDeploymentResourceInstancesAllRetryWith) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/{resource_kind}/{ref_id}/instances/_stop][%d] stopDeploymentResourceInstancesAllRetryWith  %+v", 449, o.Payload)
}

func (o *StopDeploymentResourceInstancesAllRetryWith) String() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/{resource_kind}/{ref_id}/instances/_stop][%d] stopDeploymentResourceInstancesAllRetryWith  %+v", 449, o.Payload)
}

func (o *StopDeploymentResourceInstancesAllRetryWith) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *StopDeploymentResourceInstancesAllRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewStopDeploymentResourceInstancesAllInternalServerError creates a StopDeploymentResourceInstancesAllInternalServerError with default headers values
func NewStopDeploymentResourceInstancesAllInternalServerError() *StopDeploymentResourceInstancesAllInternalServerError {
	return &StopDeploymentResourceInstancesAllInternalServerError{}
}

/*
StopDeploymentResourceInstancesAllInternalServerError describes a response with status code 500, with default header values.

A Resource that was previously stored no longer exists. (code: `deployments.deployment_resource_no_longer_exists`)
*/
type StopDeploymentResourceInstancesAllInternalServerError struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

// IsSuccess returns true when this stop deployment resource instances all internal server error response has a 2xx status code
func (o *StopDeploymentResourceInstancesAllInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stop deployment resource instances all internal server error response has a 3xx status code
func (o *StopDeploymentResourceInstancesAllInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stop deployment resource instances all internal server error response has a 4xx status code
func (o *StopDeploymentResourceInstancesAllInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this stop deployment resource instances all internal server error response has a 5xx status code
func (o *StopDeploymentResourceInstancesAllInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this stop deployment resource instances all internal server error response a status code equal to that given
func (o *StopDeploymentResourceInstancesAllInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the stop deployment resource instances all internal server error response
func (o *StopDeploymentResourceInstancesAllInternalServerError) Code() int {
	return 500
}

func (o *StopDeploymentResourceInstancesAllInternalServerError) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/{resource_kind}/{ref_id}/instances/_stop][%d] stopDeploymentResourceInstancesAllInternalServerError  %+v", 500, o.Payload)
}

func (o *StopDeploymentResourceInstancesAllInternalServerError) String() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/{resource_kind}/{ref_id}/instances/_stop][%d] stopDeploymentResourceInstancesAllInternalServerError  %+v", 500, o.Payload)
}

func (o *StopDeploymentResourceInstancesAllInternalServerError) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *StopDeploymentResourceInstancesAllInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
