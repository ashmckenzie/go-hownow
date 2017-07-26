hownow
======

[![Build Status](https://travis-ci.org/ashmckenzie/go-hownow.svg?branch=master)](https://travis-ci.org/ashmckenzie/go-hownow)
[![Go Report Card](https://goreportcard.com/badge/github.com/ashmckenzie/go-hownow)](https://goreportcard.com/report/github.com/ashmckenzie/go-hownow)
[![Codecov](https://img.shields.io/codecov/c/github/ashmckenzie/go-hownow.svg)](https://codecov.io/gh/ashmckenzie/go-hownow)

How now should be displayed.  Uses the very awesome https://github.com/jinzhu/now library.

```
How now should be displayed

Usage:
  hownow [flags]
  hownow [command]

Available Commands:
  bod         Beginning of day
  bom         Beginning of month
  bow         Beginning of week
  eod         End of day
  eom         End of month
  eow         End of week
  help        Help about any command

Flags:
      --debug         enable debugging
  -e, --epoch         format as seconds since epoch
  -h, --help          help for hownow
  -n, --no-new-line   don't print a newline at the end
      --offset int    offset now
  -v, --version       show version
```

Examples
--------

Beginning of day

```
$ bin/hownow bod
2017-06-14 00:00:00 +1000 AEST
```

Beginning of day, formatted as epoch

```
$ bin/hownow -e bod
1497362400
```

Beginning of day, formatted as epoch and no newline (useful for scripts, % at the end us how zsh shows no new line)

```
$ bin/hownow -en bod
1497362400%
```

Install
-------

`go get github.com/ashmckenzie/go-hownow/hownow`

or download a release:

https://github.com/ashmckenzie/go-hownow/releases

## Building

1. `make`

## Thanks

* [github.com/spf13/cobra](https://github.com/spf13/cobra)
* [github.com/jinzhu/now](https://github.com/jinzhu/now)

License
-------

MIT License

Copyright (c) 2017 Ash McKenzie

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
