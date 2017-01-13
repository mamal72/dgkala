[![Build Status](https://travis-ci.org/mamal72/dgkala.svg?branch=master)](https://travis-ci.org/mamal72/dgkala)
[![Go Report Card](https://goreportcard.com/badge/github.com/mamal72/dgkala)](https://goreportcard.com/report/github.com/mamal72/dgkala)
[![Coverage Status](https://coveralls.io/repos/github/mamal72/dgkala/badge.svg?branch=master)](https://coveralls.io/github/mamal72/dgkala?branch=master)
[![GoDoc](https://godoc.org/github.com/mamal72/dgkala?status.svg)](https://godoc.org/github.com/mamal72/dgkala)
[![license](https://img.shields.io/github/license/mamal72/dgkala.svg)](https://github.com/mamal72/dgkala/blob/master/LICENSE)

# dgkala

This is a simple Go package to interact with [Digikala](https://www.digikala.com) website. It's a WIP and more methods may be added in the future.


## Installation

```bash
go get github.com/mamal72/dgkala
```


## Usage

```go
package main

import "github.com/mamal72/dgkala"

func main() {
    // Get special offers
    offers, err := dgkala.SpecialOffers() // []SpecialOffer, error
}
```


## Tests

```bash
go test
```


## Ideas || Issues
Just fill an issue and describe it. I'll check it ASAP!


## Contribution

You can fork the repository, improve or fix some part of it and then send the pull requests back if you want to see them here. I really appreciate that. :heart:

Remember to write a few tests for your code before sending pull requests.


## License

Licensed under the [MIT License](https://github.com/mamal72/dgkala/blob/master/LICENSE).