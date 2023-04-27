# Go Http Matchers
> This project defines a set of functions that return predicates and extractors for making decisions about http request.

[![Go Lang Version](https://img.shields.io/badge/go-1.11-00ADD8.svg?style=plastic)](http://golang.com)
[![Go Doc](https://img.shields.io/badge/godoc-reference-00ADD8.svg?style=plastic)](https://godoc.org/github.com/danapsimer/go-http-matchers)
[![Go Report Card](https://goreportcard.com/badge/github.com/danapsimer/go-http-matchers?style=plastic)](https://goreportcard.com/report/github.com/danapsimer/go-http-matchers)
[![codecov](https://img.shields.io/codecov/c/github/danapsimer/go-http-matchers.svg?style=plastic)](https://codecov.io/gh/danapsimer/go-http-matchers)
[![CircleCI](https://img.shields.io/circleci/project/github/danapsimer/go-http-matchers.svg?style=plastic)](https://circleci.com/gh/danapsimer/go-http-matchers/tree/master)

The process of making decisions on how to handle an HTTP Request as in a mocking framework like
[Mockery](http://github.com/danapsimer/mockery) or a reverse proxy like [Iluvitar](http://github.com/danapsimer/iluvitar)
can result in cumbersome code that is hard to follow, so we built this library to make this process easier.

See [Mockery](http://github.com/danapsimer/mockery) for an example usage.

## Installation

```sh
go get github.com/danapsimer/go-http-matchers
```

## Development setup

```
git clone https://github.com/danapsimer/go-http-matchers.git
cd go-http-matchers
go mod download
```

## Release History

* v0.0.6 - Filled out this README.

## Meta

* Maintainer: Dana P'Simer 
* Twitter: [@ComputersFearMe](https://twitter.com/computersfearme) 
* EMail: danap@dhptech.com

Distributed under the Apache 2.0 license. See ``LICENSE`` for more information.

## Contributing

1. Fork it (<https://github.com/danapsimer/go-http-matchers/fork>)
2. Create your feature branch (`git checkout -b feature/fooBar`)
3. Commit your changes (`git commit -am 'Add some fooBar'`)
4. Push to the branch (`git push origin feature/fooBar`)
5. Create a new Pull Request
