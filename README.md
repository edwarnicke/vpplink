vpplink is a api to simplify interacting with govpp

## Generation

To use govpp simply create a directory into which you wish to generate govpp and add a gen.go file:

```golang
//go:build tools

package vpplink

import (
	_ "github.com/edwarnicke/vpplink/cmd"
	_ "go.fd.io/govpp/binapi"
)

// Run using go generate -tags tools ./...
//go:generate go run github.com/edwarnicke/vpplink/cmd --binapi-package "go.fd.io/govpp/binapi"
```

Replace `go.fd.io/govpp/binapi` with the govpp binapi of your choice, then run:

```bash
go generate -tags tools ./...
go mod tidy
```

and it will generate your vpplink matching your binapi of choice.

## Generation Example

Alternately, you can simply fork https://github.com/edwarnicke/vpplink-example-consumer , edit the [gen.go](https://github.com/edwarnicke/vpplink-example-consumer/blob/main/vpplink/gen.go)
file to point to your binapi of choice, and run:

```bash
go generate -tags tools ./...
go mod tidy
```
