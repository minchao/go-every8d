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

### Use webhook to receive the sending report and reply message

```go
func main() {
	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		report, err := every8d.ParseReportMessage(r)
		if err != nil {
			// Handle error...
			return
		}
		// Process report message...
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Printf("ListenAndServe error: %v", err)
	}
}
```

## Develop

### Command-line Tool

Change directory to [cmd/every8d](./cmd/every8d), and build the tool:

```
$ cd cmd/every8d
$ go build
``` 

Run:

```
$ ./every8d
EVERY8D SMS CLI tool

Usage:
  every8d [command]

Available Commands:
  credit          Query credit
  delivery-status Query to retrieve the delivery status
  help            Help about any command
  send            Send an SMS
  webhook         Webhook to receive the sending report and reply message

Flags:
  -h, --help              help for every8d
      --password string   EVERY8D Password
      --username string   EVERY8D Username

Use "every8d [command] --help" for more information about a command.

```

Example to send SMS:

```
$ ./every8d send --username=0987654321 --password=password --dest=0987654321 --msg="Hello, 世界"
Credit: 79.00
Sent: 1
Cost: 1.00
Unsent: 0
BatchID: 00000000-00000-0000-0000-000000000000

```

### Run all tests

```
$ go test -v -race .
```

## License

This library is distributed under the BSD-style license found in the [LICENSE](./LICENSE) file.
