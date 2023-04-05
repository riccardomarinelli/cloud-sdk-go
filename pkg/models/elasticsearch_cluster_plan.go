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

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ElasticsearchClusterPlan The plan for the Elasticsearch cluster.
//
// swagger:model ElasticsearchClusterPlan
type ElasticsearchClusterPlan struct {

	// Enable autoscaling for this Elasticsearch cluster.
	AutoscalingEnabled *bool `json:"autoscaling_enabled,omitempty"`

	// cluster topology
	// Required: true
	ClusterTopology []*ElasticsearchClusterTopologyElement `json:"cluster_topology"`

	// Documents which deployment template was used in the creation of this plan
	DeploymentTemplate *DeploymentTemplateReference `json:"deployment_template,omitempty"`

	// elasticsearch
	// Required: true
	Elasticsearch *ElasticsearchConfiguration `json:"elasticsearch"`

	// transient
	Transient *TransientElasticsearchPlanConfiguration `json:"transient,omitempty"`
}

// Validate validates this elasticsearch cluster plan
func (m *ElasticsearchClusterPlan) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterTopology(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeploymentTemplate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateElasticsearch(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTransient(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ElasticsearchClusterPlan) validateClusterTopology(formats strfmt.Registry) error {

	if err := validate.Required("cluster_topology", "body", m.ClusterTopology); err != nil {
		return err
	}

	for i := 0; i < len(m.ClusterTopology); i++ {
		if swag.IsZero(m.ClusterTopology[i]) { // not required
			continue
		}

		if m.ClusterTopology[i] != nil {
			if err := m.ClusterTopology[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cluster_topology" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("cluster_topology" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ElasticsearchClusterPlan) validateDeploymentTemplate(formats strfmt.Registry) error {
	if swag.IsZero(m.DeploymentTemplate) { // not required
		return nil
	}

	if m.DeploymentTemplate != nil {
		if err := m.DeploymentTemplate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deployment_template")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("deployment_template")
			}
			return err
		}
	}

	return nil
}

func (m *ElasticsearchClusterPlan) validateElasticsearch(formats strfmt.Registry) error {

	if err := validate.Required("elasticsearch", "body", m.Elasticsearch); err != nil {
		return err
	}

	if m.Elasticsearch != nil {
		if err := m.Elasticsearch.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("elasticsearch")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("elasticsearch")
			}
			return err
		}
	}

	return nil
}

func (m *ElasticsearchClusterPlan) validateTransient(formats strfmt.Registry) error {
	if swag.IsZero(m.Transient) { // not required
		return nil
	}

	if m.Transient != nil {
		if err := m.Transient.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("transient")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("transient")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this elasticsearch cluster plan based on the context it is used
func (m *ElasticsearchClusterPlan) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateClusterTopology(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDeploymentTemplate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateElasticsearch(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTransient(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ElasticsearchClusterPlan) contextValidateClusterTopology(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ClusterTopology); i++ {

		if m.ClusterTopology[i] != nil {
			if err := m.ClusterTopology[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cluster_topology" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("cluster_topology" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ElasticsearchClusterPlan) contextValidateDeploymentTemplate(ctx context.Context, formats strfmt.Registry) error {

	if m.DeploymentTemplate != nil {
		if err := m.DeploymentTemplate.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deployment_template")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("deployment_template")
			}
			return err
		}
	}

	return nil
}

func (m *ElasticsearchClusterPlan) contextValidateElasticsearch(ctx context.Context, formats strfmt.Registry) error {

	if m.Elasticsearch != nil {
		if err := m.Elasticsearch.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("elasticsearch")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("elasticsearch")
			}
			return err
		}
	}

	return nil
}

func (m *ElasticsearchClusterPlan) contextValidateTransient(ctx context.Context, formats strfmt.Registry) error {

	if m.Transient != nil {
		if err := m.Transient.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("transient")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("transient")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ElasticsearchClusterPlan) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ElasticsearchClusterPlan) UnmarshalBinary(b []byte) error {
	var res ElasticsearchClusterPlan
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
