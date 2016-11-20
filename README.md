delocalize
==========
[![Build Status](https://travis-ci.org/ushios/delocalize.svg?branch=master)](https://travis-ci.org/ushios/delocalize)
[![Coverage Status](https://coveralls.io/repos/ushios/delocalize/badge.svg?branch=master&service=github)](https://coveralls.io/github/ushios/delocalize?branch=master)

remove .localized file

Installation
------------

```bash
$ go get github.com/ushios/delocalize
```


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
