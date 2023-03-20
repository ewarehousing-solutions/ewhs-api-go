package ewhs

import (
	"context"
	"encoding/json"
)

type GdprService service

type RequestPersonData struct {
	Email string `json:"email,omitempty"`
}

type RedactPersonData struct {
	Email string `json:"email,omitempty"`
}

type Message struct {
	Message string `json:"message,omitempty"`
}

func (gs *GdprService) RequestPersonData(ctx context.Context, req RequestPersonData) (message *Message, res *Response, err error) {
	res, err = gs.client.post(ctx, "wms/gdpr/request-person-data/", req, nil)

	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &message); err != nil {
		return
	}

	return
}

func (gs *GdprService) RedactPersonData(ctx context.Context, req RedactPersonData) (message *Message, res *Response, err error) {
	res, err = gs.client.post(ctx, "wms/gdpr/redact-person-data/", req, nil)

	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &message); err != nil {
		return
	}

	return
}
