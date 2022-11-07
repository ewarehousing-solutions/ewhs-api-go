package ewhs

type VariantsService service

type Variant struct {
	ID                 string  `json:"id,omitempty"`
	ArticleCode        string  `json:"article_code,omitempty"`
	Name               string  `json:"name,omitempty"`
	Description        string  `json:"description,omitempty"`
	Ean                string  `json:"ean,omitempty"`
	Sku                string  `json:"sku,omitempty"`
	HsTariffCode       string  `json:"hs_tariff_code,omitempty"`
	Height             string  `json:"height,omitempty"`
	Depth              string  `json:"depth,omitempty"`
	Width              string  `json:"width,omitempty"`
	Weight             string  `json:"weight,omitempty"`
	Expirable          bool    `json:"expirable,omitempty"`
	CountryOfOrigin    string  `json:"country_of_origin,omitempty"`
	UsingSerialNumbers bool    `json:"using_serial_numbers,omitempty"`
	Value              float64 `json:"value,omitempty"`
}
