package twilio

import (
	"context"
	"net/url"
)

const FlowPathPart = "Flows"

type FlowService struct {
	client *Client
}

type Flow struct {
	Sid          string     `json:"sid"`
	AccountSid   string     `json:"account_sid"`
	FriendlyName string     `json:"friendly_name"`
	Status       string     `json:"status"`
	Valid        string     `json:"valid"`
	RawErrors       interface{} `json:"errors"`
  Errors string
	WebhookUrl   string     `json:"webhook_url"`
	DateCreated  TwilioTime `json:"date_created"`
	DateUpdated  TwilioTime `json:"date_updated"`
	Url          string     `json:"url"`
	RawDefinition   interface{} `json:"definition"`
	Definition   string
}

func (f *FlowService) Create(ctx context.Context, data url.Values) (*Flow, error) {
  flow := new(Flow)

	err := f.client.CreateResource(ctx, FlowPathPart, data, flow)

	return flow, err
}

func (f *FlowService) Get(ctx context.Context, sid string) (*Flow, error) {
  flow := new(Flow)

	err := f.client.GetResource(ctx, FlowPathPart, sid, flow)

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
