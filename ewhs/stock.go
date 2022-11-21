package ewhs

import (
	"context"
	"encoding/json"
	"time"
)

type StockService service

type Stock struct {
	ID              string      `json:"id,omitempty"`
	ArticleCode     string      `json:"article_code,omitempty"`
	Ean             string      `json:"ean,omitempty"`
	Sku             interface{} `json:"sku,omitempty"`
	StockPhysical   int         `json:"stock_physical,omitempty"`
	StockSalable    int         `json:"stock_salable,omitempty"`
	StockAvailable  int         `json:"stock_available,omitempty"`
	StockQuarantine int         `json:"stock_quarantine,omitempty"`
	StockPickable   int         `json:"stock_pickable,omitempty"`
	ModifiedAt      time.Time   `json:"modified_at,omitempty"`
	Variant         Variant     `json:"variant,omitempty"`
}

type StockListOptions struct {
	ArticleCode string `url:"article_code,omitempty"`
	Ean         string `url:"ean,omitempty"`
	Sku         string `url:"sku,omitempty"`
	Search      string `url:"search,omitempty"`
	ModifiedGte string `url:"modified_gte,omitempty"`
	Page        int    `url:"page,omitempty"`
	From        string `url:"from,omitempty"`
	To          string `url:"to,omitempty"`
	Limit       int    `url:"limit,omitempty"`
	Direction   string `url:"direction,omitempty"`
}

func (ss *StockService) List(ctx context.Context, opts *StockListOptions) (list *[]Stock, res *Response, err error) {
	res, err = ss.client.get(ctx, "wms/stock/", opts)
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
