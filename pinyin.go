package main

import (
  "strings"
  // "regexp"
)

// var hzRegexp = regexp.MustCompile("^[\u4e00-\u9fa5]$")

var dict = map[string]string{}

// Translate chinese to pinyin
func T(args ...string) string {
  splitter, argsLen := " ", len(args)

  var fslice = []string{}

  if argsLen == 0 {
    return ""
  }

  if argsLen > 1 {
    splitter = args[argsLen - 1]
  }

  for _, rune := range args[0] {
    fslice = append(fslice, string(rune))
  }

  return strings.Join(fslice, splitter)
}


func init() {

  for _, line := range strings.Split(dictDate, "\n") {

    var words = strings.SplitN(line, "\t", 2)

    dict[words[0]] = words[1]
  }
}