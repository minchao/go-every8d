# go-every8d

[![Build Status](https://travis-ci.org/minchao/go-every8d.svg?branch=master)](https://travis-ci.org/minchao/go-every8d)
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

### Credit query

Retrieve your account balance.

```go
credit, err := client.GetCredit(context.Background())
```

## License

This library is distributed under the BSD-style license found in the [LICENSE](./LICENSE) file.
