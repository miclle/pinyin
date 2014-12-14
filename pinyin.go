package pinyin

import (
  "log"
  "os"
)

// Translate chinese to pinyin
func T(args ...string) string {
  pyStr, splitter, argsLen := "", " ", len(args)

  if argsLen == 0 {
    return pyStr
  }

  if argsLen > 1 {
    splitter = args[argsLen - 1]
  }

  for _, rune := range args[0] {
    pyStr += string(rune) + splitter
  }

  return pyStr
}


type Pinyin struct {

  initialized bool
}


func (py *Pinyin) Init() {
  if py.initialized {
    log.Fatal("Initialization can not be repeated!")
  }

  py.initialized = true

  file, err := os.Open("data/pinyin-utf8.dat")
  defer file.Close()
  if err != nil {
    log.Fatal(err)
  }

}