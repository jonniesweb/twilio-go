package twilio

import (
  
)

const flowPathPart = "Flows"

type FlowService struct {
  client *Client
}

type Flow struct {
  Sid string `json:"sid"`
}
