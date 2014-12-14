package main

import (
  "os"
  "strings"
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

  println(pinyin.T(str, splitter))

  var py pinyin.Pinyin

  py.Init()

}
