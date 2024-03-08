/*
* Copyright 2018 - Present Okta, Inc.
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*      http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */

// Code generated by okta openapi generator. DO NOT EDIT.

package okta

import (
	"context"
	"fmt"

	"github.com/n-h-n/okta-sdk-golang/v2/okta/query"
)

type ProfileMappingResource resource

type ProfileMapping struct {
	Links      interface{}                        `json:"_links,omitempty"`
	Id         string                             `json:"id,omitempty"`
	Properties map[string]*ProfileMappingProperty `json:"properties,omitempty"`
	Source     *ProfileMappingSource              `json:"source,omitempty"`
	Target     *ProfileMappingSource              `json:"target,omitempty"`
}

// Fetches a single Profile Mapping referenced by its ID.
func (m *ProfileMappingResource) GetProfileMapping(ctx context.Context, mappingId string) (*ProfileMapping, *Response, error) {
	url := fmt.Sprintf("/api/v1/mappings/%v", mappingId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var profileMapping *ProfileMapping

	resp, err := rq.Do(ctx, req, &profileMapping)
	if err != nil {
		return nil, resp, err
	}

	return profileMapping, resp, nil
}

// Updates an existing Profile Mapping by adding, updating, or removing one or many Property Mappings.
func (m *ProfileMappingResource) UpdateProfileMapping(ctx context.Context, mappingId string, body ProfileMapping) (*ProfileMapping, *Response, error) {
	url := fmt.Sprintf("/api/v1/mappings/%v", mappingId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}

	var profileMapping *ProfileMapping

	resp, err := rq.Do(ctx, req, &profileMapping)
	if err != nil {
		return nil, resp, err
	}

	return profileMapping, resp, nil
}

// Enumerates Profile Mappings in your organization with pagination.
func (m *ProfileMappingResource) ListProfileMappings(ctx context.Context, qp *query.Params) ([]*ProfileMapping, *Response, error) {
	url := fmt.Sprintf("/api/v1/mappings")
	if qp != nil {
		url = url + qp.String()
	}

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var profileMapping []*ProfileMapping

	resp, err := rq.Do(ctx, req, &profileMapping)
	if err != nil {
		return nil, resp, err
	}

	return profileMapping, resp, nil
}
