# go-every8d

[![Build Status](https://travis-ci.org/minchao/go-every8d.svg?branch=master)](https://travis-ci.org/minchao/go-every8d)
[![Go Report Card](https://goreportcard.com/badge/github.com/minchao/go-every8d)](https://goreportcard.com/report/github.com/minchao/go-every8d)
[![codecov](https://codecov.io/gh/minchao/go-every8d/branch/master/graph/badge.svg)](https://codecov.io/gh/minchao/go-every8d)

A Go client library for accessing the EVERY8D SMS API.

## Installation

Use go get to install.

```
go get -u github.com/minchao/go-every8d
```

## Usage

```go
import "github.com/minchao/go-every8d"
```

Construct a new API client, then use to access the EVERY8D API. For example:

```go
client := every8d.NewClient("UID", "PWD", nil)
ctx := context.Background()

// Retrieve your account balance
credit, err := client.GetCredit(ctx)
```

## License

This library is distributed under the BSD-style license found in the [LICENSE](./LICENSE) file.
