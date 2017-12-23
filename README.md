# go-every8d

[![Build Status](https://travis-ci.org/minchao/go-every8d.svg?branch=master)](https://travis-ci.org/minchao/go-every8d)
[![GoDoc](https://godoc.org/github.com/minchao/go-every8d?status.svg)](https://godoc.org/github.com/minchao/go-every8d)
[![Go Report Card](https://goreportcard.com/badge/github.com/minchao/go-every8d)](https://goreportcard.com/report/github.com/minchao/go-every8d)
[![codecov](https://codecov.io/gh/minchao/go-every8d/branch/master/graph/badge.svg)](https://codecov.io/gh/minchao/go-every8d)

A Go client library for accessing the [EVERY8D](http://global.every8d.com.tw/) SMS API.

## Installation

Use go get to install.

```
go get -u github.com/minchao/go-every8d
```

## Usage

Import the `go-every8d` package.

```go
import "github.com/minchao/go-every8d"
```

Construct a new API client, then use to access the EVERY8D API. For example:

```go
client := every8d.NewClient("UID", "PWD", nil)
```

### Send an SMS

```go
message := every8d.Message{
    Subject:         "Note",
    Content:         "Hello, 世界",
    Destination:     "+886987654321",
}

result, err := client.Send(context.Background(), message)
```

### Query to retrieve the delivery status

```go
resp, err := client.GetDeliveryStatus(context.Background(), batchID, pageNo)
```

### Query credit

Retrieve your account balance.

```go
credit, err := client.GetCredit(context.Background())
```

## Develop

### Run all tests

```go
go test -v -race .
```

## License

This library is distributed under the BSD-style license found in the [LICENSE](./LICENSE) file.
