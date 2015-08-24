# musta [![Build Status](https://secure.travis-ci.org/shiwano/musta.png?branch=master)](http://travis-ci.org/shiwano/musta)

A command-line tool for Mustache template, written in Go.

## Installation

Via binary [releases](https://github.com/shiwano/musta/releases).

Via `go-get`:

```bash
$ go get github.com/shiwano/musta
```

Via [Homebrew](http://brew.sh/):

```bash
brew tap shiwano/formulas
brew install musta
```

## Usage

With a template file.

```bash
$ cat template.mustache
foo: {{foo}}
$ musta -t template.mustache foo=bar
foo: bar
```

With a template string.

```bash
$ musta -T "foobar: {{foo.bar}}, qux: {{#qux}}{{value}}{{^last}},{{/last}}{{/qux}}" foo.bar=1 qux=1,2,3
foobar: 1, qux: 1,2,3
```

With piped JSON data.

```bash
$ cat tokyo-weather.json | musta -t weather.mustache >> tokyo-weather.html
```

### Parsing Array

musta will parse a comma-separated string as an array.

For example, `values=a,b,c` will be parsed like the following JSON data.

```json
{
  "values": [
    { "index": 0, "value": "a", "first": true,  "last": false },
    { "index": 1, "value": "b", "first": false, "last": false },
    { "index": 2, "value": "c", "first": false, "last": true }
  ]
}
```

## License

Copyright (c) 2015 Shogo Iwano
Licensed under the MIT license.
