package ewhs

import (
	"encoding/json"
	"fmt"
)

type WebhooksService service

type Webhook struct {
	URL        string `json:"url,omitempty"`
	Group      string `json:"group,omitempty"`
	Action     string `json:"action,omitempty"`
	HashSecret string `json:"hash_secret,omitempty"`
}

func (ws *WebhooksService) List() (list *[]Webhook, res *Response, err error) {
	res, err = ws.client.get("wms/webhooks/", nil)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &list); err != nil {
		return
	}

	return
}

func (ws *WebhooksService) Get(webhookID string) (webhook *Webhook, res *Response, err error) {
	res, err = ws.client.get(fmt.Sprintf("webhooks/%s/", webhookID), nil)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &webhook); err != nil {
		return
	}

	return
}

func (ws *WebhooksService) Create(wh Webhook) (webhook *Webhook, res *Response, err error) {
	res, err = ws.client.post("webhooks/", wh, nil)

	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &webhook); err != nil {
		return
	}

	return
}

func (ws *WebhooksService) Update(webhookID string, wh Webhook) (webhook *Webhook, res *Response, err error) {
	res, err = ws.client.patch(fmt.Sprintf("webhooks/%s/", webhookID), wh, nil)

	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &webhook); err != nil {
		return
	}

	return
}

func (os *OrdersService) Delete(webhookID string) (webhook *Webhook, res *Response, err error) {
	res, err = os.client.delete(fmt.Sprintf("webhooks/%s/", webhookID), nil)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &webhook); err != nil {
		return
	}

	return
}
