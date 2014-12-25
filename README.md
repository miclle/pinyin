Pinyin
=============

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

    fmt.Print(pinyin.TT("中国人", "-"))}
}
```

#License

Code released under the MIT license.