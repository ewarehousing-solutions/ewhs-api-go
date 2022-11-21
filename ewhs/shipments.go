package ewhs

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type ShipmentsService service

type Shipment struct {
	ID                        string      `json:"id,omitempty"`
	Customer                  string      `json:"customer,omitempty"`
	CreatedAt                 time.Time   `json:"created_at,omitempty"`
	OrderExternalReference    string      `json:"order_external_reference,omitempty"`
	ShipmentExternalReference interface{} `json:"shipment_external_reference,omitempty"`
	Reference                 string      `json:"reference,omitempty"`
}

type ShipmentListOptions struct {
	OrderExternalReference string `url:"order_external_reference,omitempty"`
	From                   string `url:"from,omitempty"`
	To                     string `url:"to,omitempty"`
	Limit                  int    `url:"limit,omitempty"`
	Direction              string `url:"direction,omitempty"`
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
