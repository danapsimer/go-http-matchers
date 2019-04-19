# Go Http Matchers
> This project defines a set of functions that return predicates and extractors for making decisions about http request.

[![Go Lang Version](https://img.shields.io/badge/go-1.11.0-blue.svg)](http://golang.com)
[![CircleCI](https://circleci.com/gh/bluesoftdev/go-http-matchers/tree/master.svg?style=svg)](https://circleci.com/gh/bluesoftdev/go-http-matchers/tree/master)

The process of making decisions on how to handle an HTTP Request as in a mocking framework like
[Mockery](http://github.com/bluesoftdev/mockery) or a reverse proxy like [Iluvitar](http://github.com/bluesoftdev/iluvitar)
can result in cumbersome code that is hard to follow, so we built this library to make this process easier.

See [Mockery](http://github.com/bluesoftdev/mockery) for an example usage.

## Installation

```sh
go get github.com/bluesoftdev/go-http-matchers
```

## Documentation

[GoDoc](https://godoc.org/github.com/bluesoftdev/go-http-matchers)

## Development setup

```
git clone https://github.com/bluesoftdev/go-http-matchers.git
cd go-http-matchers
go mod download
```

## Release History

* v0.0.6 - Filled out this README.

## Meta

* Dana P'Simer 
* [@ComputersFearMe](https://twitter.com/computersfearme) 
* danap@bluesoftdev.com

Distributed under the Apache 2.0 license. See ``LICENSE`` for more information.

## Contributing

1. Fork it (<https://github.com/bluesoftdev/go-http-matchers/fork>)
2. Create your feature branch (`git checkout -b feature/fooBar`)
3. Commit your changes (`git commit -am 'Add some fooBar'`)
4. Push to the branch (`git push origin feature/fooBar`)
5. Create a new Pull Request
