package ewhs

import (
	"context"
	"encoding/json"
	"fmt"
)

type ShippingMethodsService service

type ShippingMethod struct {
	ID               string `json:"id,omitempty"`
	Shipper          string `json:"shipper,omitempty"`
	ShipperCode      string `json:"shipper_code,omitempty"`
	Code             string `json:"code,omitempty"`
	Description      string `json:"description,omitempty"`
	ShippingSoftware string `json:"shipping_software,omitempty"`
}

type ShippingMethodListOptions struct {
	From      string `url:"from,omitempty"`
	To        string `url:"to,omitempty"`
	Limit     int    `url:"limit,omitempty"`
	Direction string `url:"direction,omitempty"`
}

func (ss *ShippingMethodsService) List(ctx context.Context, opts *ShippingMethodListOptions) (list *[]ShippingMethod, res *Response, err error) {
	res, err = ss.client.get(ctx, "wms/shippingmethods/", opts)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &list); err != nil {
		return
	}

	return
}

func (ss *ShippingMethodsService) Get(ctx context.Context, shippingMethodID string) (shippingMethod *ShippingMethod, res *Response, err error) {
	res, err = ss.client.get(ctx, fmt.Sprintf("wms/shippingmethods/%s/", shippingMethodID), nil)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &shippingMethod); err != nil {
		return
	}

	return
}
