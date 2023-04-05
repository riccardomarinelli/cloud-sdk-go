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

package deployments_traffic_filter

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// CreateTrafficFilterRulesetReader is a Reader for the CreateTrafficFilterRuleset structure.
type CreateTrafficFilterRulesetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateTrafficFilterRulesetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateTrafficFilterRulesetCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewCreateTrafficFilterRulesetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateTrafficFilterRulesetCreated creates a CreateTrafficFilterRulesetCreated with default headers values
func NewCreateTrafficFilterRulesetCreated() *CreateTrafficFilterRulesetCreated {
	return &CreateTrafficFilterRulesetCreated{}
}

/*
CreateTrafficFilterRulesetCreated describes a response with status code 201, with default header values.

The ruleset definition is valid and the creation has started.
*/
type CreateTrafficFilterRulesetCreated struct {
	Payload *models.TrafficFilterRulesetResponse
}

// IsSuccess returns true when this create traffic filter ruleset created response has a 2xx status code
func (o *CreateTrafficFilterRulesetCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create traffic filter ruleset created response has a 3xx status code
func (o *CreateTrafficFilterRulesetCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create traffic filter ruleset created response has a 4xx status code
func (o *CreateTrafficFilterRulesetCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create traffic filter ruleset created response has a 5xx status code
func (o *CreateTrafficFilterRulesetCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create traffic filter ruleset created response a status code equal to that given
func (o *CreateTrafficFilterRulesetCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create traffic filter ruleset created response
func (o *CreateTrafficFilterRulesetCreated) Code() int {
	return 201
}

func (o *CreateTrafficFilterRulesetCreated) Error() string {
	return fmt.Sprintf("[POST /deployments/traffic-filter/rulesets][%d] createTrafficFilterRulesetCreated  %+v", 201, o.Payload)
}

func (o *CreateTrafficFilterRulesetCreated) String() string {
	return fmt.Sprintf("[POST /deployments/traffic-filter/rulesets][%d] createTrafficFilterRulesetCreated  %+v", 201, o.Payload)
}

func (o *CreateTrafficFilterRulesetCreated) GetPayload() *models.TrafficFilterRulesetResponse {
	return o.Payload
}

func (o *CreateTrafficFilterRulesetCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TrafficFilterRulesetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTrafficFilterRulesetInternalServerError creates a CreateTrafficFilterRulesetInternalServerError with default headers values
func NewCreateTrafficFilterRulesetInternalServerError() *CreateTrafficFilterRulesetInternalServerError {
	return &CreateTrafficFilterRulesetInternalServerError{}
}

/*
CreateTrafficFilterRulesetInternalServerError describes a response with status code 500, with default header values.

Error creating the traffic filter ruleset. (code: `traffic_filter.request_execution_failed`)
*/
type CreateTrafficFilterRulesetInternalServerError struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

// IsSuccess returns true when this create traffic filter ruleset internal server error response has a 2xx status code
func (o *CreateTrafficFilterRulesetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create traffic filter ruleset internal server error response has a 3xx status code
func (o *CreateTrafficFilterRulesetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create traffic filter ruleset internal server error response has a 4xx status code
func (o *CreateTrafficFilterRulesetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create traffic filter ruleset internal server error response has a 5xx status code
func (o *CreateTrafficFilterRulesetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create traffic filter ruleset internal server error response a status code equal to that given
func (o *CreateTrafficFilterRulesetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create traffic filter ruleset internal server error response
func (o *CreateTrafficFilterRulesetInternalServerError) Code() int {
	return 500
}

func (o *CreateTrafficFilterRulesetInternalServerError) Error() string {
	return fmt.Sprintf("[POST /deployments/traffic-filter/rulesets][%d] createTrafficFilterRulesetInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateTrafficFilterRulesetInternalServerError) String() string {
	return fmt.Sprintf("[POST /deployments/traffic-filter/rulesets][%d] createTrafficFilterRulesetInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateTrafficFilterRulesetInternalServerError) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *CreateTrafficFilterRulesetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
