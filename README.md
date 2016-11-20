delocalize
==========
[![CircleCI](https://circleci.com/gh/ushios/delocalize.svg?style=svg&circle-token=3e552577232060a1cbf22c98afc7885ec83cd558)](https://circleci.com/gh/ushios/delocalize)
[![Build Status](https://travis-ci.org/ushios/delocalize.svg?branch=master)](https://travis-ci.org/ushios/delocalize)
[![Coverage Status](https://coveralls.io/repos/ushios/delocalize/badge.svg?branch=master&service=github)](https://coveralls.io/github/ushios/delocalize?branch=master)


remove .localized file

Installation
------------

```bash
$ go get github.com/ushios/delocalize
```

Documentation
-------------

[![GoDoc](https://godoc.org/github.com/ushios/delocalize?status.svg)](https://godoc.org/github.com/ushios/delocalize)

Usage
-----

### Print localized file list under home directory

```bash
$ delocalize -t ~/

/Users/xxxxx/Documents/.localized
/Users/xxxxx/Downloads/.localized
...
```

### Delete localized files

using `-d` option.

``` bash
$ delocalize -d -t ~/

deleted: /Users/xxxxx/Documents/.localized
deleted: /Users/xxxxx/Downloads/.localized
...
```
