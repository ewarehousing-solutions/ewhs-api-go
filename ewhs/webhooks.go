package ewhs

import (
	"context"
	"encoding/json"
	"fmt"
)

type WebhooksService service

type WebhookResults struct {
	Count    int         `json:"count,omitempty"`
	Next     interface{} `json:"next,omitempty"`
	Previous interface{} `json:"previous,omitempty"`
	Results  []Webhook   `json:"results,omitempty"`
}

type Webhook struct {
	ID         string `json:"id,omitempty"`
	URL        string `json:"url,omitempty"`
	Group      string `json:"group,omitempty"`
	Action     string `json:"action,omitempty"`
	HashSecret string `json:"hash_secret,omitempty"`
}

func (ws *WebhooksService) List(ctx context.Context) (list *[]Webhook, res *Response, err error) {
	res, err = ws.client.get(ctx, "webhooks/", nil)
	if err != nil {
		return
	}

	var i WebhookResults

	if err = json.Unmarshal(res.content, &i); err != nil {
		return
	}

	return &i.Results, nil, nil
}

func (ws *WebhooksService) Get(ctx context.Context, webhookID string) (webhook *Webhook, res *Response, err error) {
	res, err = ws.client.get(ctx, fmt.Sprintf("webhooks/%s/", webhookID), nil)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &webhook); err != nil {
		return
	}

	return
}

func (ws *WebhooksService) Create(ctx context.Context, wh Webhook) (webhook *Webhook, res *Response, err error) {
	res, err = ws.client.post(ctx, "webhooks/", wh, nil)

	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &webhook); err != nil {
		return
	}

	return
}

func (ws *WebhooksService) Update(ctx context.Context, webhookID string, wh Webhook) (webhook *Webhook, res *Response, err error) {
	res, err = ws.client.patch(ctx, fmt.Sprintf("webhooks/%s/", webhookID), wh, nil)

	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &webhook); err != nil {
		return
	}

	return
}

func (os *WebhooksService) Delete(ctx context.Context, webhookID string) (webhook *Webhook, res *Response, err error) {
	res, err = os.client.delete(ctx, fmt.Sprintf("webhooks/%s/", webhookID), nil)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &webhook); err != nil {
		return
	}

	return
}
