package pinyin

import (
  "strings"
  // "regexp"
)

// var hzRegexp = regexp.MustCompile("^[\u4e00-\u9fa5]$")

var dict = map[string]string{}

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


func init() {

  for _, line := range strings.Split(dictDate, "\n") {

    var words = strings.SplitN(line, " ", 2)

    dict[words[0]] = words[1]
  }
}