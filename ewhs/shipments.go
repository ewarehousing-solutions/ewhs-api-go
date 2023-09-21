package ewhs

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type ShipmentsService service

type Shipment struct {
	ID                        string                 `json:"id,omitempty"`
	Customer                  string                 `json:"customer,omitempty"`
	CreatedAt                 time.Time              `json:"created_at,omitempty"`
	OrderID                   string                 `json:"order_id,omitempty"`
	OrderExternalReference    string                 `json:"order_external_reference,omitempty"`
	ShipmentExternalReference interface{}            `json:"shipment_external_reference,omitempty"`
	Reference                 string                 `json:"reference,omitempty"`
	ShippingMethod            ShipmentShippingMethod `json:"shipping_method,omitempty"`
	ShipmentLabels            []ShipmentLabels       `json:"shipment_labels,omitempty"`
	ShipmentLines             []ShipmentLines        `json:"shipment_lines,omitempty"`
	ShippingAddress           ShippingAddress        `json:"shipping_address,omitempty"`
}
type ShipmentShippingMethod struct {
	ID               string `json:"id,omitempty"`
	Shipper          string `json:"shipper,omitempty"`
	ShipperCode      string `json:"shipper_code,omitempty"`
	Code             string `json:"code,omitempty"`
	Description      string `json:"description,omitempty"`
	ShippingSoftware string `json:"shipping_software,omitempty"`
}
type ShipmentLabels struct {
	LabelCode    string `json:"label_code,omitempty"`
	TrackingCode string `json:"tracking_code,omitempty"`
	TrackingURL  string `json:"tracking_url,omitempty"`
}

type ShipmentLines struct {
	ShippedQuantity    int           `json:"shipped_quantity,omitempty"`
	ShippedEan         string        `json:"shipped_ean,omitempty"`
	ShippedArticleCode string        `json:"shipped_article_code,omitempty"`
	ShippedSku         string        `json:"shipped_sku,omitempty"`
	SerialNumbers      []interface{} `json:"serial_numbers,omitempty"`
	Variant            Variant       `json:"variant,omitempty"`
}

type ShipmentListOptions struct {
	OrderExternalReference string `url:"order_external_reference,omitempty"`
	From                   string `url:"from,omitempty"`
	To                     string `url:"to,omitempty"`
	Limit                  int    `url:"limit,omitempty"`
	Direction              string `url:"direction,omitempty"`
	Page                   int    `url:"page,omitempty"`
}

func (ss *ShipmentsService) List(ctx context.Context, opts *ShipmentListOptions) (list *[]Shipment, res *Response, err error) {
	res, err = ss.client.get(ctx, "wms/shipments/", opts)
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

func (ss *ShipmentsService) Get(ctx context.Context, shipmentID string) (shipment *Shipment, res *Response, err error) {
	res, err = ss.client.get(ctx, fmt.Sprintf("wms/shipments/%s/", shipmentID), nil)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &shipment); err != nil {
		return
	}

	return
}
