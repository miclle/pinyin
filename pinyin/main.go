package main

import (
  "os"
  "strings"
  "fmt"
  "pinyin"
)

func main() {
  str, splitter := "", " "

  if len(os.Args) == 2 {
    str = os.Args[1]
  }

  if len(os.Args) > 2 {
    str = strings.Join(os.Args[1:len(os.Args)-1], "")
    splitter = os.Args[len(os.Args)-1]
  }

  fmt.Print(pinyin.T(str, splitter))
}
