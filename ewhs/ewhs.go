package ewhs

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/go-querystring/query"
	"io"
	"net/http"
	"net/url"
	"runtime"
	"strings"
	"time"
)

const (
	BaseURL     string = "https://eu.middleware.ewarehousing-solutions.com/"
	TestBaseURL string = "https://eu-dev.middleware.ewarehousing-solutions.com/"

	RequestContentType string = "application/json"
	AuthHeader         string = "Authorization"
	CustomerCodeHeader string = "X-Customer-Code"
	WmsCodeHeader      string = "X-Wms-Code"
)

var (
	errEmptyAuthKey = errors.New("you must provide a non-empty authentication key")
	errBadBaseURL   = errors.New("malformed base url, it must contain a trailing slash")
)

// Client represents a client.
type Client struct {
	BaseURL   *url.URL
	client    *http.Client
	userAgent string
	common    service
	config    *Config

	authToken string

	// Services
	Articles        *ArticlesService
	Inbounds        *InboundsService
	Orders          *OrdersService
	Stock           *StockService
	Shipments       *ShipmentsService
	ShippingMethods *ShippingMethodsService
	Variants        *VariantsService
	Webhooks        *WebhooksService
}

type service struct {
	client *Client
}

func (c *Client) get(ctx context.Context, uri string, options interface{}) (res *Response, err error) {
	if options != nil {
		v, _ := query.Values(options)
		uri = fmt.Sprintf("%s?%s", uri, v.Encode())
	}

	req, err := c.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return
	}

	return c.Do(req)
}

func (c *Client) post(ctx context.Context, uri string, body interface{}, options interface{}) (res *Response, err error) {
	if options != nil {
		v, _ := query.Values(options)
		uri = fmt.Sprintf("%s?%s", uri, v.Encode())
	}

	req, err := c.NewRequest(ctx, http.MethodPost, uri, body)
	if err != nil {
		return
	}

	return c.Do(req)
}

func (c *Client) patch(ctx context.Context, uri string, body interface{}, options interface{}) (res *Response, err error) {
	if options != nil {
		v, _ := query.Values(options)
		uri = fmt.Sprintf("%s?%s", uri, v.Encode())
	}

	req, err := c.NewRequest(ctx, http.MethodPatch, uri, body)
	if err != nil {
		return
	}

	return c.Do(req)
}

func (c *Client) delete(ctx context.Context, uri string, options interface{}) (res *Response, err error) {
	if options != nil {
		v, _ := query.Values(options)
		uri = fmt.Sprintf("%s?%s", uri, v.Encode())
	}

	req, err := c.NewRequest(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return
	}

	return c.Do(req)
}

// Authorize fetches a new access token based on login credentials
func (c *Client) authorize(ctx context.Context) (res *Response, err error) {
	req, err := c.NewRequest(ctx, http.MethodPost, "wms/auth/login/", nil)
	if err != nil {
		return
	}

	res, err = c.Do(req)

	if err != nil {
		return res, err
	}

	authToken := AuthToken{}

	if err = json.Unmarshal(res.content, &authToken); err != nil {
		return
	}

	c.authToken = authToken.Token

	return res, nil
}

func (c *Client) WithAuthToken(k string) error {
	if k == "" {
		return errEmptyAuthKey
	}

	c.authToken = strings.TrimSpace(k)

	return nil
}

func (c *Client) NewRequest(ctx context.Context, method string, uri string, body interface{}) (req *http.Request, err error) {
	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, errBadBaseURL
	}

	u, err := c.BaseURL.Parse(uri)
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

	if ctx == nil {
		ctx = context.Background()
	}

	req, err = http.NewRequestWithContext(ctx, method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", RequestContentType)
	req.Header.Set("Content-Type", RequestContentType)

	if c.config != nil {
		req.Header.Set(CustomerCodeHeader, c.config.CustomerCode)
		req.Header.Set(WmsCodeHeader, c.config.WmsCode)
	}

	// TODO: allow expand headers
	//if ctx.Value("Expand") != nil {
	//}

	// if no auth token is found -> authorize first
	if c.authToken == "" && uri != "wms/auth/login/" {
		_, err := c.authorize(ctx)
		if err != nil {
			return nil, err
		}
	}

	req.Header.Add(AuthHeader, strings.Join([]string{"Bearer", c.authToken}, " "))

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

func NewClient(baseClient *http.Client, c *Config) (ewhs *Client, err error) {
	if baseClient == nil {
		baseClient = http.DefaultClient
		{
			baseClient.Timeout = 30 * time.Second
		}
	}

	bURL := BaseURL

	if c.Testing == true {
		bURL = TestBaseURL
	}

	u, _ := url.Parse(bURL)

	ewhs = &Client{
		BaseURL: u,
		client:  baseClient,
		config:  c,
	}

	ewhs.common.client = ewhs

	// services for resources
	ewhs.Articles = (*ArticlesService)(&ewhs.common)
	ewhs.Inbounds = (*InboundsService)(&ewhs.common)
	ewhs.Orders = (*OrdersService)(&ewhs.common)
	ewhs.Stock = (*StockService)(&ewhs.common)
	ewhs.Shipments = (*ShipmentsService)(&ewhs.common)
	ewhs.ShippingMethods = (*ShippingMethodsService)(&ewhs.common)
	ewhs.Variants = (*VariantsService)(&ewhs.common)
	ewhs.Webhooks = (*WebhooksService)(&ewhs.common)

	ewhs.userAgent = strings.Join([]string{
		runtime.GOOS,
		runtime.GOARCH,
		runtime.Version(),
	}, ";")

	return ewhs, nil
}

func newError(rsp *Response) error {
	merr := &BaseError{}

	//if rsp.ContentLength > 0 {
	//	err := json.Unmarshal(rsp.content, merr)
	//	if err != nil {
	//		return err
	//	}
	//} else {
	merr.Status = rsp.StatusCode
	merr.Title = rsp.Status
	merr.Detail = string(rsp.content)
	//}

	return merr
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
