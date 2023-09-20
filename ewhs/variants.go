package ewhs

import (
	"context"
	"encoding/json"
	"fmt"
)

type VariantsService service

type Variant struct {
	ID                 string  `json:"id,omitempty"`
	ArticleCode        string  `json:"article_code,omitempty"`
	Name               string  `json:"name,omitempty"`
	Description        string  `json:"description,omitempty"`
	Ean                string  `json:"ean,omitempty"`
	Sku                string  `json:"sku,omitempty"`
	HsTariffCode       string  `json:"hs_tariff_code,omitempty"`
	Height             int64   `json:"height,omitempty"`
	Depth              int64   `json:"depth,omitempty"`
	Width              int64   `json:"width,omitempty"`
	Weight             int64   `json:"weight,omitempty"`
	Expirable          bool    `json:"expirable,omitempty"`
	CountryOfOrigin    string  `json:"country_of_origin,omitempty"`
	UsingSerialNumbers bool    `json:"using_serial_numbers,omitempty"`
	Value              float64 `json:"value,omitempty"`
}

type VariantListOptions struct {
	From      string `url:"from,omitempty"`
	To        string `url:"to,omitempty"`
	Limit     int    `url:"limit,omitempty"`
	Direction string `url:"direction,omitempty"`
}

func (vs *VariantsService) List(ctx context.Context, opts *VariantListOptions) (list *[]Variant, res *Response, err error) {
	res, err = vs.client.get(ctx, "wms/variants/", opts)
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

func (vs *VariantsService) Get(ctx context.Context, variantID string) (variant *Variant, res *Response, err error) {
	res, err = vs.client.get(ctx, fmt.Sprintf("wms/variants/%s/", variantID), nil)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &variant); err != nil {
		return
	}

	return
}

func (vs *VariantsService) Create(ctx context.Context, ord Order) (order *Order, res *Response, err error) {
	res, err = vs.client.post(ctx, "wms/variants/", ord, nil)

	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &order); err != nil {
		return
	}

	return
}

func (vs *VariantsService) Update(ctx context.Context, variantID string, vr Variant) (variant *Variant, res *Response, err error) {
	res, err = vs.client.patch(ctx, fmt.Sprintf("wms/variants/%s/", variantID), vr, nil)

	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &variant); err != nil {
		return
	}

	return
}
