hownow
======

[![Build Status](https://travis-ci.org/ashmckenzie/go-hownow.svg?branch=master)](https://travis-ci.org/ashmckenzie/go-hownow)

How now should be displayed.  Uses the very awesome https://github.com/jinzhu/now library.

```
usage: hownow [<flags>] <command> [<args> ...]

How now should be displayed

Flags:
      --help         Show context-sensitive help (also try --help-long and --help-man).
  -n, --no-new-line  Don't print a newline at the end
  -e, --epoch        Format as seconds since epoch.
      --version      Show application version.

Commands:
  help [<command>...]
    Show help.

  bod
    Beginning of day.

  bow
    Beginning of week.

  eod
    End of day.

  eow
    End of week.
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
