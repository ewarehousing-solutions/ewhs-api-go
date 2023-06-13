# eWarehousing Solutions Go Library

This library provides convenient access to the eWarehousing Solutions API from applications written in the Go
language.

## Documentation

https://api.docs.ewarehousing-solutions.com/

## Installation

```
go get -u github.com/ewarehousing-solutions/ewhs-api-go
```

### Requirements

- Go 19+


## Usage

### Client setup

```go
config := ewhs.NewConfig("username", "password", "wms_code", "customer_code", false)
client, err := ewhs.NewClient(nil, config)
if err != nil {
    log.Fatal(err)
}
```

If you want to use the testing API, set the `testing` parameter to true
```go
config := ewhs.NewConfig("username", "password", "wms_code", "customer_code", true)
```

### Filtering
Collection routes allow query parameters to filter the dataset. Query params can be added using the corresponding `ListOptions{}` type. For example, if you want to search for a specific order reference:
```go
options := &ewhs.OrderListOptions{
	Reference: "SEARCH_REFERENCE",
}
orders, res, err := client.Orders.List(context.Background(), options)
```


### Webhook verification
The package provides a helper which can be used to easily verify the webhooks
```go
func ValidateWebhook(httpRequest *http.Request) bool {
    secret := "super-secret-password"
    return VerifyWebhookRequest(httpRequest, secret)
}
```

### Expanding responses
Many objects allow you to request additional information as an expanded response by using the expand header. This parameter is available on all API requests, and applies to the response of that request only. Checkout the [documentation](https://api.docs.ewarehousing-solutions.com/expanding-responses) for all possible options.
```go
ctx := context.Background()

order, res, err := client.Orders.Get(context.WithValue(ctx, "Expand", "order_lines"))
```

## Support

[www.ewarehousing-solutions.nl](https://ewarehousing-solutions.nl/) — info@ewarehousing-solutions.nl

## Credits

This package is heavily based on [https://github.com/VictorAvelar/mollie-api-go](VictorAvelar/mollie-api-go).