vpplink is a high level API over the GoVPP binapi

## APIs

apis may be found in package github.com/edwarnicke/vpplink/api

## Generating vpplink implementation against your own binapi

To generate an implementation for vpplink against your binapi:

1.  Add a gen.go file to the directory where you want to generate your vpplink implementation:
```go
//go:build tools

package vpplink

import (
	_ "github.com/edwarnicke/vpplink/cmd"
	_ "go.fd.io/govpp/binapi"
)

// Run using go generate -tags tools ./...
//go:generate go run github.com/edwarnicke/vpplink/cmd --binapi-package "go.fd.io/govpp/binapi"
```
Substituting your binapi-package for go.fd.io/govpp/binapi
2. Run ```go mod tidy```
3. Run ```go generate -tags tools ./...```

## Developing vpplink templates

To develop vpplink templates, simply edit/add templates in the templates/ subdirectory

To test your work, run ```go test -v ./...``` which will generate a testdata/ dir containing code generated from the template.
