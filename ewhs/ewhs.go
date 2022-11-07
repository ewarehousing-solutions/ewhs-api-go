package ewhs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
	"io"
	"net/http"
	"net/url"
)

const (
	BaseURL            string = "https://eu.middleware.ewarehousing-solutions.com"
	RequestContentType string = "application/json"
)

// Client represents a client.
type Client struct {
	client  *http.Client
	baseURL *url.URL
	common  service
	config  *Config

	// Services
	Inbounds  *InboundsService
	Orders    *OrdersService
	Stock     *StockService
	Shipments *ShipmentsService
	Variants  *VariantsService
	Webhooks  *WebhooksService
}

type service struct {
	client *Client
}

func (c *Client) get(uri string, options interface{}) (res *Response, err error) {
	if options != nil {
		v, _ := query.Values(options)
		uri = fmt.Sprintf("%s?%s", uri, v.Encode())
	}

	req, err := c.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return
	}

	return c.Do(req)
}

func (c *Client) post(uri string, body interface{}, options interface{}) (res *Response, err error) {
	if options != nil {
		v, _ := query.Values(options)
		uri = fmt.Sprintf("%s?%s", uri, v.Encode())
	}

	req, err := c.NewRequest(http.MethodPost, uri, body)
	if err != nil {
		return
	}

	return c.Do(req)
}

func (c *Client) patch(uri string, body interface{}, options interface{}) (res *Response, err error) {
	if options != nil {
		v, _ := query.Values(options)
		uri = fmt.Sprintf("%s?%s", uri, v.Encode())
	}

	req, err := c.NewRequest(http.MethodPatch, uri, body)
	if err != nil {
		return
	}

	return c.Do(req)
}

func (c *Client) delete(uri string, options interface{}) (res *Response, err error) {
	if options != nil {
		v, _ := query.Values(options)
		uri = fmt.Sprintf("%s?%s", uri, v.Encode())
	}

	req, err := c.NewRequest(http.MethodDelete, uri, nil)
	if err != nil {
		return
	}

	return c.Do(req)
}

func (c *Client) NewRequest(method string, uri string, body interface{}) (req *http.Request, err error) {
	u, err := c.baseURL.Parse(uri)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err = http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", RequestContentType)
	req.Header.Set("Accept", RequestContentType)

	return req, nil
}

// Do sends an API request and returns the API response or returned as an
// error if an API error has occurred.
func (c *Client) Do(req *http.Request) (*Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("httperror: %w", err)
	}

	defer resp.Body.Close()

	response, err := newResponse(resp)
	if err != nil {
		return response, err
	}

	err = CheckResponse(response)
	if err != nil {
		return response, err
	}

	return response, nil
}

func NewClient(baseClient *http.Client, c *Config) Client {
	if baseClient == nil {
		baseClient = http.DefaultClient
	}

	u, _ := url.Parse(BaseURL)

	ewhs := Client{
		baseURL: u,
		client:  baseClient,
		config:  c,
	}

	ewhs.common.client = &ewhs

	// services for resources
	ewhs.Inbounds = (*InboundsService)(&ewhs.common)
	ewhs.Orders = (*OrdersService)(&ewhs.common)
	ewhs.Stock = (*StockService)(&ewhs.common)
	ewhs.Shipments = (*ShipmentsService)(&ewhs.common)
	ewhs.Variants = (*VariantsService)(&ewhs.common)
	ewhs.Webhooks = (*WebhooksService)(&ewhs.common)

	return ewhs
}

/*
Error reports details on a failed API request.
*/
type Error struct {
	Code     int       `json:"code"`
	Message  string    `json:"message"`
	Response *Response `json:"response"`
}

// Error function complies with the error interface
func (e *Error) Error() string {
	return fmt.Sprintf("response failed with status %v", e.Message)
}

func newError(r *Response) *Error {
	var e Error
	e.Response = r
	e.Code = r.StatusCode
	e.Message = r.Status
	return &e
}

type Response struct {
	*http.Response
	content []byte
}

func newResponse(rsp *http.Response) (*Response, error) {
	res := Response{Response: rsp}

	data, err := io.ReadAll(rsp.Body)
	if err != nil {
		return &res, err
	}

	res.content = data

	rsp.Body = io.NopCloser(bytes.NewBuffer(data))
	res.Response = rsp

	return &res, nil
}

func CheckResponse(r *Response) error {
	if r.StatusCode >= http.StatusMultipleChoices {
		return newError(r)
	}
	return nil
}
