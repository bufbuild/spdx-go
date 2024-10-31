# spdx-go

[![Build](https://github.com/bufbuild/spdx-go/actions/workflows/ci.yaml/badge.svg?branch=main)](https://github.com/bufbuild/spdx-go/actions/workflows/ci.yaml)
[![Report Card](https://goreportcard.com/badge/buf.build/go/spdx)](https://goreportcard.com/report/buf.build/go/spdx)
[![GoDoc](https://pkg.go.dev/badge/buf.build/go/spdx.svg)](https://pkg.go.dev/buf.build/go/spdx)
[![Slack](https://img.shields.io/badge/slack-buf-%23e01563)](https://buf.build/links/slack)

A simple Golang library that contains license information from [SPDX](https://spdx.dev).

See [spdx.org/licenses](https://spdx.org/licenses) for more details.

```go
import (
	"buf.build/go/spdx"
)

if license, ok := spdx.LicenseForID("apache-2.0"); ok {
	fmt.Println(license.ID)              // Apache-2.0
	fmt.Println(license.Name)            // Apache License 2.0
	fmt.Println(license.Reference)       // https://spdx.org/licenses/Apache-2.0.html
	fmt.Println(license.ReferenceNumber) // 501
	fmt.Println(license.DetailsURL)      // https://spdx.org/licenses/Apache-2.0.json
	fmt.Println(license.Deprecated)      // false
	fmt.Println(license.SeeAlso)         // [https://www.apache.org/licenses/LICENSE-2.0 https://opensource.org/licenses/Apache-2.0]
	fmt.Println(license.OSIApproved)     // true
}

for _, license := range spdx.AllLicenses() {
    fmt.Println(license)
}
```

All IDs are guaranteed to match the regular expression `^[a-zA-Z0-9-.+]+$`.

## Status: Beta

This repository is still in beta, however will be promoted to stable very soon.

## Legal

The original data is sourced from [SPDX](https://spdx.dev).

This library is offered under the [Apache 2 license](https://github.com/bufbuild/spdx-go/blob/main/LICENSE).
