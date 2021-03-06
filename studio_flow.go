package twilio

import (
	"context"
	"encoding/json"
	"net/url"
)

const FlowPathPart = "Flows"

type FlowService struct {
	client *Client
}

type Flow struct {
	Sid           string      `json:"sid"`
	AccountSid    string      `json:"account_sid"`
	FriendlyName  string      `json:"friendly_name"`
	Status        string      `json:"status"`
	Valid         bool        `json:"valid"`
	WebhookUrl    string      `json:"webhook_url"`
	DateCreated   TwilioTime  `json:"date_created"`
	DateUpdated   TwilioTime  `json:"date_updated"`
	Url           string      `json:"url"`
	RawDefinition interface{} `json:"definition"`
	Definition    string      `json:"-"`
}

func (f *FlowService) Create(ctx context.Context, data url.Values) (*Flow, error) {
	flow := new(Flow)

	err := f.client.CreateResource(ctx, FlowPathPart, data, flow)

	return flow, err
}

func (f *FlowService) Get(ctx context.Context, sid string) (*Flow, error) {
	flow := new(Flow)

	err := f.client.GetResource(ctx, FlowPathPart, sid, flow)

	if err != nil {
		return flow, err
	}

	json, err := json.Marshal(flow.RawDefinition)
	if err != nil {
		return flow, err
	}
	flow.Definition = string(json)

	return flow, err
}

func (f *FlowService) Update(ctx context.Context, sid string, data url.Values) (*Flow, error) {
	flow := new(Flow)

	err := f.client.UpdateResource(ctx, FlowPathPart, sid, data, flow)

	return flow, err
}

func (f *FlowService) Delete(ctx context.Context, sid string) error {
	return f.client.DeleteResource(ctx, FlowPathPart, sid)
}
