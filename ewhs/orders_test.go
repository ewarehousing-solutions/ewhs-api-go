package ewhs

import (
	"context"
	"fmt"
	"github.com/ewarehousing-solutions/ewhs-api-go/test/testdata"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
)

type ordersServiceSuite struct{ suite.Suite }

func (os *ordersServiceSuite) TestOrdersService_Get() {
	type args struct {
		ctx   context.Context
		order string
	}
	cases := []struct {
		name    string
		args    args
		wantErr bool
		err     error
		pre     func()
		handler http.HandlerFunc
	}{
		{
			"get orders works as expected.",
			args{
				context.Background(),
				"c9165f93-8301-4aaa-9f64-27f191c0c778",
			},
			false,
			nil,
			func() {
				tClient.WithAuthToken("eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9")
			},
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(os.T(), r, AuthHeader, "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9")
				testHeader(os.T(), r, CustomerCodeHeader, "test_customer")
				testHeader(os.T(), r, WmsCodeHeader, "test_wms")
				testMethod(os.T(), r, "GET")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				_, _ = w.Write([]byte(testdata.GetOrderResponse))
			},
		},
		{
			"get orders works as expected.",
			args{
				context.Background(),
				"c9165f93-8301-4aaa-9f64-27f191c0c778",
			},
			false,
			nil,
			func() {
				tClient.WithAuthToken("eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9")
			},
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(os.T(), r, AuthHeader, "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9")
				testMethod(os.T(), r, "GET")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				_, _ = w.Write([]byte(testdata.GetOrderResponse))
			},
		},
		{
			"get orders, an error is returned from the server",
			args{
				context.Background(),
				"c9165f93-8301-4aaa-9f64-27f191c0c778",
			},
			true,
			fmt.Errorf("500 - 500 Internal Server Error"),
			func() {
				tClient.WithAuthToken("eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9")
			},
			errorHandler,
		},
		{
			"get orders, an error occurs when parsing json",
			args{
				context.Background(),
				"c9165f93-8301-4aaa-9f64-27f191c0c778",
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			func() {
				tClient.WithAuthToken("eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9")
			},
			encodingHandler,
		},
		{
			"get orders, invalid url when building request",
			args{
				context.Background(),
				"c9165f93-8301-4aaa-9f64-27f191c0c778",
			},
			true,
			errBadBaseURL,
			crashSrv,
			errorHandler,
		},
	}

	for _, c := range cases {
		setup()
		defer teardown()

		os.T().Run(c.name, func(t *testing.T) {
			c.pre()
			tMux.HandleFunc(fmt.Sprintf("/wms/orders/%s/", c.args.order), c.handler)

			m, res, err := tClient.Orders.Get(c.args.ctx, c.args.order)
			if c.wantErr {
				os.NotNil(err)
				os.EqualError(err, c.err.Error())
			} else {
				os.Nil(err)
				os.IsType(&Order{}, m)
				os.IsType(&http.Response{}, res.Response)
			}
		})
	}
}

func (os *ordersServiceSuite) TestOrdersService_Create() {
	type args struct {
		ctx   context.Context
		order Order
	}
	cases := []struct {
		name    string
		args    args
		wantErr bool
		err     error
		pre     func()
		handler http.HandlerFunc
	}{
		{
			"create orders works as expected.",
			args{
				context.Background(),
				Order{
					ShippingAddress: ShippingAddress{},
				},
			},
			false,
			nil,
			func() {
				tClient.WithAuthToken("eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9")
			},
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(os.T(), r, AuthHeader, "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9")
				testHeader(os.T(), r, CustomerCodeHeader, "test_customer")
				testHeader(os.T(), r, WmsCodeHeader, "test_wms")
				testMethod(os.T(), r, "POST")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				_, _ = w.Write([]byte(testdata.GetOrderResponse))
			},
		},
		{
			"create orders works as expected.",
			args{
				context.Background(),
				Order{},
			},
			false,
			nil,
			func() {
				tClient.WithAuthToken("eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9")
			},
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(os.T(), r, AuthHeader, "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9")
				testMethod(os.T(), r, "POST")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				_, _ = w.Write([]byte(testdata.GetOrderResponse))
			},
		},
		{
			"create orders, an error is returned from the server",
			args{
				context.Background(),
				Order{},
			},
			true,
			fmt.Errorf("500 - 500 Internal Server Error"),
			func() {
				tClient.WithAuthToken("eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9")
			},
			errorHandler,
		},
		{
			"create orders, an error occurs when parsing json",
			args{
				context.Background(),
				Order{},
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			func() {
				tClient.WithAuthToken("eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9")
			},
			encodingHandler,
		},
	}

	for _, c := range cases {
		setup()
		defer teardown()

		os.T().Run(c.name, func(t *testing.T) {
			c.pre()
			tMux.HandleFunc("/wms/orders/", c.handler)

			m, res, err := tClient.Orders.Create(c.args.ctx, c.args.order)
			if c.wantErr {
				os.NotNil(err)
				os.EqualError(err, c.err.Error())
			} else {
				os.Nil(err)
				os.IsType(&Order{}, m)
				os.IsType(&http.Response{}, res.Response)
			}
		})
	}
}

func TestOrdersService(t *testing.T) {
	suite.Run(t, new(ordersServiceSuite))
}
