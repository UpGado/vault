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

// AUTO-GENERATED!  DO NOT EDIT FILE DIRECTLY

package okta

import (
	"context"
	"fmt"
	"time"
)

type EventHookResource resource

type EventHook struct {
	Links              interface{}         `json:"_links,omitempty"`
	Channel            *EventHookChannel   `json:"channel,omitempty"`
	Created            *time.Time          `json:"created,omitempty"`
	CreatedBy          string              `json:"createdBy,omitempty"`
	Events             *EventSubscriptions `json:"events,omitempty"`
	Id                 string              `json:"id,omitempty"`
	LastUpdated        *time.Time          `json:"lastUpdated,omitempty"`
	Name               string              `json:"name,omitempty"`
	Status             string              `json:"status,omitempty"`
	VerificationStatus string              `json:"verificationStatus,omitempty"`
}

func (m *EventHookResource) CreateEventHook(ctx context.Context, body EventHook) (*EventHook, *Response, error) {
	url := fmt.Sprintf("/api/v1/eventHooks")

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}

	var eventHook *EventHook

	resp, err := m.client.requestExecutor.Do(ctx, req, &eventHook)
	if err != nil {
		return nil, resp, err
	}

	return eventHook, resp, nil
}

func (m *EventHookResource) GetEventHook(ctx context.Context, eventHookId string) (*EventHook, *Response, error) {
	url := fmt.Sprintf("/api/v1/eventHooks/%v", eventHookId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var eventHook *EventHook

	resp, err := m.client.requestExecutor.Do(ctx, req, &eventHook)
	if err != nil {
		return nil, resp, err
	}

	return eventHook, resp, nil
}

func (m *EventHookResource) UpdateEventHook(ctx context.Context, eventHookId string, body EventHook) (*EventHook, *Response, error) {
	url := fmt.Sprintf("/api/v1/eventHooks/%v", eventHookId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("PUT", url, body)
	if err != nil {
		return nil, nil, err
	}

	var eventHook *EventHook

	resp, err := m.client.requestExecutor.Do(ctx, req, &eventHook)
	if err != nil {
		return nil, resp, err
	}

	return eventHook, resp, nil
}

func (m *EventHookResource) DeleteEventHook(ctx context.Context, eventHookId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/eventHooks/%v", eventHookId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (m *EventHookResource) ListEventHooks(ctx context.Context) ([]*EventHook, *Response, error) {
	url := fmt.Sprintf("/api/v1/eventHooks")

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var eventHook []*EventHook

	resp, err := m.client.requestExecutor.Do(ctx, req, &eventHook)
	if err != nil {
		return nil, resp, err
	}

	return eventHook, resp, nil
}

func (m *EventHookResource) ActivateEventHook(ctx context.Context, eventHookId string) (*EventHook, *Response, error) {
	url := fmt.Sprintf("/api/v1/eventHooks/%v/lifecycle/activate", eventHookId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var eventHook *EventHook

	resp, err := m.client.requestExecutor.Do(ctx, req, &eventHook)
	if err != nil {
		return nil, resp, err
	}

	return eventHook, resp, nil
}

func (m *EventHookResource) DeactivateEventHook(ctx context.Context, eventHookId string) (*EventHook, *Response, error) {
	url := fmt.Sprintf("/api/v1/eventHooks/%v/lifecycle/deactivate", eventHookId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var eventHook *EventHook

	resp, err := m.client.requestExecutor.Do(ctx, req, &eventHook)
	if err != nil {
		return nil, resp, err
	}

	return eventHook, resp, nil
}

func (m *EventHookResource) VerifyEventHook(ctx context.Context, eventHookId string) (*EventHook, *Response, error) {
	url := fmt.Sprintf("/api/v1/eventHooks/%v/lifecycle/verify", eventHookId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var eventHook *EventHook

	resp, err := m.client.requestExecutor.Do(ctx, req, &eventHook)
	if err != nil {
		return nil, resp, err
	}

	return eventHook, resp, nil
}
