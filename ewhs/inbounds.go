package ewhs

import (
	"context"
	"encoding/json"
	"fmt"
)

type InboundsService service

type Inbound struct {
	ExternalReference string        `json:"external_reference,omitempty"`
	Note              string        `json:"note,omitempty"`
	InboundDate       string        `json:"inbound_date,omitempty"`
	InboundLines      []InboundLine `json:"inbound_lines,omitempty"`
}
type InboundLine struct {
	Quantity    int    `json:"quantity,omitempty"`
	ArticleCode string `json:"article_code,omitempty"`
}

type InboundListOptions struct {
	Reference string `url:"reference,omitempty"`
	Status    string `url:"status,omitempty"`
	Type      int    `url:"type,omitempty"`
	Page      int    `url:"page,omitempty"`
	From      string `url:"from,omitempty"`
	To        string `url:"to,omitempty"`
	Limit     int    `url:"limit,omitempty"`
	Sort      string `url:"sort,omitempty"`
	Direction string `url:"direction,omitempty"`
}

func (is *InboundsService) List(ctx context.Context, opts *InboundListOptions) (list *[]Inbound, res *Response, err error) {
	res, err = is.client.get(ctx, "wms/inbounds/", opts)
	if err != nil {
		return
	}

	if len(res.content) == 0 {
		return
	}

	if err = json.Unmarshal(res.content, &list); err != nil {
		return
	}

	return
}

func (is *InboundsService) Get(ctx context.Context, inboundID string) (inbound *Inbound, res *Response, err error) {
	res, err = is.client.get(ctx, fmt.Sprintf("wms/inbounds/%s/", inboundID), nil)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &inbound); err != nil {
		return
	}

	return
}

func (is *InboundsService) Create(ctx context.Context, inb Inbound) (inbound *Inbound, res *Response, err error) {
	res, err = is.client.post(ctx, "wms/inbounds/", inb, nil)

	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &inbound); err != nil {
		return
	}

	return
}

func (is *InboundsService) Update(ctx context.Context, inboundID string, inb Inbound) (inbound *Inbound, res *Response, err error) {
	res, err = is.client.patch(ctx, fmt.Sprintf("wms/inbounds/%s/", inboundID), inb, nil)

	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &inbound); err != nil {
		return
	}

	return
}

func (is *InboundsService) Cancel(ctx context.Context, inboundID string) (inbound *Inbound, res *Response, err error) {
	res, err = is.client.patch(ctx, fmt.Sprintf("wms/inbounds/%s/cancel/", inboundID), nil, nil)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &inbound); err != nil {
		return
	}

	return
}
