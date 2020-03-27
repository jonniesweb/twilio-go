package twilio

import (
	"context"
	"net/url"
)

////create
//curl -X POST https://studio.twilio.com/v2/Flows \
//--data-urlencode "CommitMessage=First draft" \
//--data-urlencode "FriendlyName=Main IVR" \
//--data-urlencode "Status=draft" \
//--data-urlencode "Definition=$DEFINITION" \
//-u AC316c55775ce12e12fbc6774fe4184990:your_auth_token

const FlowPathPart = "Flows"

type FlowService struct {
	client *Client
}

type Flow struct {
	Sid string `json:"sid"`
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
