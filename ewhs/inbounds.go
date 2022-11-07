package ewhs

import (
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

func (is *InboundsService) List(opts *InboundListOptions) (list *[]Inbound, res *Response, err error) {
	res, err = is.client.get("wms/inbounds/", opts)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &list); err != nil {
		return
	}

	return
}

func (is *InboundsService) Get(inboundID string) (inbound *Inbound, res *Response, err error) {
	res, err = is.client.get(fmt.Sprintf("wms/inbounds/%s/", inboundID), nil)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &inbound); err != nil {
		return
	}

	return
}

func (is *InboundsService) Create(inb Inbound) (inbound *Inbound, res *Response, err error) {
	res, err = is.client.post("wms/inbounds/", inb, nil)

	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &inbound); err != nil {
		return
	}

	return
}

func (is *InboundsService) Update(inboundID string, inb Inbound) (inbound *Inbound, res *Response, err error) {
	res, err = is.client.patch(fmt.Sprintf("wms/inbounds/%s/", inboundID), inb, nil)

	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &inbound); err != nil {
		return
	}

	return
}

func (is *InboundsService) Cancel(inboundID string) (inbound *Inbound, res *Response, err error) {
	res, err = is.client.patch(fmt.Sprintf("wms/inbounds/%s/cancel", inboundID), nil, nil)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &inbound); err != nil {
		return
	}

	return
}
