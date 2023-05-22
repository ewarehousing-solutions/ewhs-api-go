package ewhs

import (
	"context"
	"encoding/json"
	"fmt"
)

type OrdersService service

type Order struct {
	ID                    string          `json:"id,omitempty"`
	ExternalID            string          `json:"external_id,omitempty"`
	ExternalReference     string          `json:"external_reference,omitempty"`
	ShippingContactperson string          `json:"shipping_contactperson,omitempty"`
	RequestedDeliveryDate string          `json:"requested_delivery_date,omitempty"`
	CustomerNote          string          `json:"customer_note,omitempty"`
	ShippingEmail         string          `json:"shipping_email,omitempty"`
	ShippingMethod        string          `json:"shipping_method,omitempty"`
	Note                  string          `json:"note,omitempty"`
	Currency              string          `json:"currency,omitempty"`
	OrderAmount           int64           `json:"order_amount,omitempty"`
	Documents             []Document      `json:"documents,omitempty"`
	OrderLines            []OrderLine     `json:"order_lines,omitempty"`
	ShippingAddress       ShippingAddress `json:"shipping_address,omitempty"`
}
type Document struct {
	ShippingLabel bool   `json:"shipping_label,omitempty"`
	Title         string `json:"title,omitempty"`
	Quantity      int    `json:"quantity,omitempty"`
	OrderPrice    int    `json:"orderPrice,omitempty"`
	File          string `json:"file,omitempty"`
}
type OrderLine struct {
	Price       float64 `json:"price,omitempty"`
	Quantity    int     `json:"quantity,omitempty"`
	Description string  `json:"description,omitempty"`
	ArticleCode string  `json:"article_code,omitempty"`
}
type ShippingAddress struct {
	City                 string `json:"city,omitempty"`
	State                string `json:"state,omitempty"`
	Street               string `json:"street,omitempty"`
	Country              string `json:"country,omitempty"`
	Street2              string `json:"street2,omitempty"`
	Zipcode              string `json:"zipcode,omitempty"`
	FaxNumber            string `json:"fax_number,omitempty"`
	AddressedTo          string `json:"addressed_to,omitempty"`
	PhoneNumber          string `json:"phone_number,omitempty"`
	MobileNumber         string `json:"mobile_number,omitempty"`
	StreetNumber         string `json:"street_number,omitempty"`
	StreetNumberAddition string `json:"street_number_addition,omitempty"`
}

type OrderListOptions struct {
	Reference string `url:"reference,omitempty"`
	Status    string `url:"status,omitempty"`
	Page      int    `url:"page,omitempty"`
	From      string `url:"from,omitempty"`
	To        string `url:"to,omitempty"`
	Limit     int    `url:"limit,omitempty"`
	Sort      string `url:"sort,omitempty"`
	Direction string `url:"direction,omitempty"`
}

func (os *OrdersService) List(ctx context.Context, opts *OrderListOptions) (list *[]Order, res *Response, err error) {
	res, err = os.client.get(ctx, "wms/orders/", opts)
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

func (os *OrdersService) Get(ctx context.Context, orderID string) (order *Order, res *Response, err error) {
	res, err = os.client.get(ctx, fmt.Sprintf("wms/orders/%s/", orderID), nil)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &order); err != nil {
		return
	}

	return
}

func (os *OrdersService) Create(ctx context.Context, ord Order) (order *Order, res *Response, err error) {
	res, err = os.client.post(ctx, "wms/orders/", ord, nil)

	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &order); err != nil {
		return
	}

	return
}

func (os *OrdersService) Update(ctx context.Context, orderID string, ord Order) (order *Order, res *Response, err error) {
	res, err = os.client.patch(ctx, fmt.Sprintf("wms/orders/%s/", orderID), ord, nil)

	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &order); err != nil {
		return
	}

	return
}

func (os *OrdersService) Cancel(ctx context.Context, orderID string) (order *Order, res *Response, err error) {
	res, err = os.client.patch(ctx, fmt.Sprintf("wms/orders/%s/cancel", orderID), nil, nil)
	if err != nil {
		return
	}

	if err = json.Unmarshal(res.content, &order); err != nil {
		return
	}

	return
}
