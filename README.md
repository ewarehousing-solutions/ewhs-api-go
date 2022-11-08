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

```go
config := ewhs.NewConfig("username", "password", "wms_code", "customer_code")
client, err := ewhs.NewClient(nil, config)
if err != nil {
log.Fatal(err)
}
```

## Support

[www.ewarehousing-solutions.nl](https://ewarehousing-solutions.nl/) — info@ewarehousing-solutions.nl

## Credits

This package is heavily based on [https://github.com/VictorAvelar/mollie-api-go](VictorAvelar/mollie-api-go).