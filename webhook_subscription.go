package pagerduty

import (
	"context"
	"fmt"
	"net/http"
)

type DeliveryMethod struct {
	ID                  string `json:"id"`
	Type                string `json:"type"`
	Url                 string `json:"url"`
	Secret              string `json:"secret,omitempty"`
	TemporarilyDisabled bool   `json:"temporarily_disabled" yaml:"temporarily_disabled"`
}

type Scope struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

type WebhookSubscription struct {
	ID             string         `json:"id"`
	Type           string         `json:"type"`
	Events         []string       `json:"events"`
	Description    string         `json:"description,omitempty"`
	Filter         Scope          `json:"filter"`
	DeliveryMethod DeliveryMethod `json:"delivery_method" yaml:"delivery_method"`
}

type ListWebhookSubscriptionResponse struct {
	APIListObject
	WebhookSubscriptions []WebhookSubscription `json:"webhook_subscriptions"`
}

func (c *Client) GetWebhookSubscription(id string) (*WebhookSubscription, error) {
	return c.GetWebhookSubscriptionWithContext(context.Background(), id)
}

func (c *Client) GetWebhookSubscriptionWithContext(ctx context.Context, id string) (*WebhookSubscription, error) {
	resp, err := c.getEarlyAccess(ctx, "/webhook_subscriptions/"+id, "webhooks-v3")

	return getWebhookSubscriptionFromResponse(c, resp, err)
}

func (c *Client) ListWebhookSubscriptionsWithContext(ctx context.Context) (*ListWebhookSubscriptionResponse, error) {
	resp, err := c.getEarlyAccess(ctx, "/webhook_subscriptions", "webhooks-v3")
	if err != nil {
		return nil, err
	}

	var result ListWebhookSubscriptionResponse
	if err = c.decodeJSON(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func getWebhookSubscriptionFromResponse(c *Client, resp *http.Response, err error) (*WebhookSubscription, error) {
	if err != nil {
		return nil, err
	}

	var target map[string]WebhookSubscription
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("could not decode JSON response: %v", dErr)
	}

	const rootNode = "webhook_subscription"

	t, nodeOK := target[rootNode]
	if !nodeOK {
		return nil, fmt.Errorf("JSON response does not have %s field", rootNode)
	}

	return &t, nil
}
