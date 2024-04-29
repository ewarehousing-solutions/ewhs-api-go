package ewhs

import (
	"context"
	"github.com/ewarehousing-solutions/ewhs-api-go/test/testdata"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
)

type webhooksServiceSuite struct{ suite.Suite }

func (os *webhooksServiceSuite) TestWebhooksSuite_Delete() {
	type args struct {
		ctx     context.Context
		webhook string
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
			"delete webhooks returns 204 as expected.",
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
				testMethod(os.T(), r, "DELETE")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				} else {
					w.WriteHeader(http.StatusNoContent)
				}
				_, _ = w.Write([]byte(testdata.DeleteWebhooksResponse))
			},
		},
	}

	for _, c := range cases {
		setup()
		defer teardown()

		os.T().Run(c.name, func(t *testing.T) {
			c.pre()
			tMux.HandleFunc("/webhooks/", c.handler)

			m, res, err := tClient.Webhooks.Delete(c.args.ctx, c.args.webhook)
			if c.wantErr {
				os.NotNil(err)
				os.EqualError(err, c.err.Error())
			} else {
				os.Nil(err)
				os.IsType(&Webhook{}, m)
				os.IsType(&http.Response{}, res.Response)
			}
		})
	}
}

func TestWebhooksSuite(t *testing.T) {
	suite.Run(t, new(webhooksServiceSuite))
}
