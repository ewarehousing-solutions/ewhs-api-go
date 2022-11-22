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


### Webhook verification
The package provides a helper which can be used to easily verify the webhooks
```go
func ValidateWebhook(httpRequest *http.Request) bool {
    secret := "super-secret-password"
    return VerifyWebhookRequest(httpRequest, secret)
}
```


## Support

[www.ewarehousing-solutions.nl](https://ewarehousing-solutions.nl/) — info@ewarehousing-solutions.nl

## Credits

This package is heavily based on [https://github.com/VictorAvelar/mollie-api-go](VictorAvelar/mollie-api-go).