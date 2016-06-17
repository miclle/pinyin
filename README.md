Pinyin
=============

[![Coverage Status](https://coveralls.io/repos/github/miclle/pinyin/badge.svg?branch=master)](https://coveralls.io/github/miclle/pinyin?branch=master)


Translate chinese hanzi to pinyin.

Install
-------

    go get -u github.com/miclle/pinyin

Usage
-----

```
package main

import (
    "fmt"
    "github.com/miclle/pinyin"
)

func main() {

    fmt.Print(pinyin.T("中国人"))

    fmt.Print(pinyin.T("中国人", "-"))

    fmt.Print(pinyin.TT("中国人"))

    fmt.Print(pinyin.TT("中国人", "-"))
}
```

#License

Code released under the MIT license.