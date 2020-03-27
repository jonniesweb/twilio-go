package twilio

import (
	"context"
	"net/url"
)

const FlowPathPart = "Flows"

type FlowService struct {
	client *Client
}

type EasyFlowType = map[string]interface{}

// TODO: This should really map to the correct types, but theres some nested
// JSON that Twilio responds with that isn't worth the trouble.

// type Flow struct {
// 	Sid          string     `json:"sid"`
// 	FriendlyName string     `json:"friendly_name"`
// 	Status       string     `json:status""`
// 	Valid        string     `json:"valid"`
// 	Errors       string     `json:"errors"`
// 	WebhookUrl   string     `json:"webhook_url"`
// 	DateCreated  TwilioTime `json:"date_created"`
// 	DateUpdated  TwilioTime `json:"date_updated"`
// 	Url          string     `json:"url"`
// 	// Definition   map[string]string `json:"definition"`
// }

func (f *FlowService) Create(ctx context.Context, data url.Values) (EasyFlowType, error) {
	var flow EasyFlowType

	err := f.client.CreateResource(ctx, FlowPathPart, data, &flow)

	return flow, err
}

func (f *FlowService) Get(ctx context.Context, sid string) (EasyFlowType, error) {
	var flow EasyFlowType

	err := f.client.GetResource(ctx, FlowPathPart, sid, &flow)

	return flow, err
}

func (f *FlowService) Update(ctx context.Context, sid string, data url.Values) (EasyFlowType, error) {
	var flow EasyFlowType

	err := f.client.UpdateResource(ctx, FlowPathPart, sid, data, &flow)

	return flow, err
}

func (f *FlowService) Delete(ctx context.Context, sid string) error {
	return f.client.DeleteResource(ctx, FlowPathPart, sid)
}
