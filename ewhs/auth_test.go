package ewhs

import (
	"context"
	"fmt"
	"github.com/ewarehousing-solutions/ewhs-api-go/test/testdata"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
)

type authServiceSuite struct{ suite.Suite }

func (as *authServiceSuite) TestAuthService_Authorize() {
	type args struct {
		ctx context.Context
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
			"authorize works as expected.",
			args{
				context.Background(),
			},
			false,
			nil,
			noPre,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(as.T(), r, CustomerCodeHeader, "test_customer")
				testHeader(as.T(), r, WmsCodeHeader, "test_wms")
				testMethod(as.T(), r, "POST")

				w.WriteHeader(http.StatusOK)
				_, _ = w.Write([]byte(testdata.CreateAuthTokenResponse))
			},
		},
		{
			"wrong login credentials gives an error.",
			args{
				context.Background(),
			},
			true,
			fmt.Errorf("401 - 401 Unauthorized"),
			noPre,
			unauthorizedHandler,
		},
	}

	for _, c := range cases {
		setup()
		defer teardown()

		as.T().Run(c.name, func(t *testing.T) {
			c.pre()
			tMux.HandleFunc("/wms/auth/login/", c.handler)

			_, err := tClient.authorize(c.args.ctx)
			if c.wantErr {
				as.NotNil(err)
				as.EqualError(err, c.err.Error())
			} else {
				as.Nil(err)
				as.NotNil(tClient.authToken)
			}
		})
	}
}

func TestAuthService(t *testing.T) {
	suite.Run(t, new(authServiceSuite))
}
