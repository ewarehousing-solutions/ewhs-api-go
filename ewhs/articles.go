package ewhs

import (
	"encoding/json"
	"fmt"
)

type ArticlesService service

type Article struct {
	Name     string           `json:"name,omitempty"`
	Variants []ArticleVariant `json:"variants,omitempty"`
}
type ArticleVariant struct {
	Name               string  `json:"name,omitempty"`
	ArticleCode        string  `json:"article_code,omitempty"`
	Description        string  `json:"description,omitempty"`
	Ean                string  `json:"ean,omitempty"`
	Sku                string  `json:"sku,omitempty"`
	HsTariffCode       string  `json:"hs_tariff_code,omitempty"`
	Height             int     `json:"height,omitempty"`
	Depth              int     `json:"depth,omitempty"`
	Width              int     `json:"width,omitempty"`
	Weight             int     `json:"weight,omitempty"`
	Expirable          bool    `json:"expirable,omitempty"`
	CountryOfOrigin    string  `json:"country_of_origin,omitempty"`
	UsingSerialNumbers bool    `json:"using_serial_numbers,omitempty"`
	Value              float64 `json:"value,omitempty"`
}

type ArticleListOptions struct {
	From      string `url:"from,omitempty"`
	To        string `url:"to,omitempty"`
	Limit     int    `url:"limit,omitempty"`
	Direction string `url:"direction,omitempty"`
}

func (as *ArticlesService) List(opts *ArticleListOptions) (list *[]Article, res *Response, err error) {
	res, err = as.client.get("wms/articles/", opts)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &list); err != nil {
		return
	}

	return
}

func (as *ArticlesService) Get(articleID string) (article *Article, res *Response, err error) {
	res, err = as.client.get(fmt.Sprintf("wms/articles/%s/", articleID), nil)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &article); err != nil {
		return
	}

	return
}

func (as *ArticlesService) Create(art Article) (article *Article, res *Response, err error) {
	res, err = as.client.post("wms/articles/", art, nil)

	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &article); err != nil {
		return
	}

	return
}

func (as *ArticlesService) Update(articleID string, art Article) (article *Article, res *Response, err error) {
	res, err = as.client.patch(fmt.Sprintf("wms/articles/%s/", articleID), art, nil)

	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &article); err != nil {
		return
	}

	return
}
