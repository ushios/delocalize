delocalize
==========
[![CircleCI](https://circleci.com/gh/ushios/delocalize.svg?style=shield&circle-token=3e552577232060a1cbf22c98afc7885ec83cd558)](https://circleci.com/gh/ushios/delocalize)

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
