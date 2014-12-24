package pinyin

import (
  "strings"
  "regexp"
)

var hzRegexp = regexp.MustCompile("^[\u4e00-\u9fa5]$")

// Translate chinese to pinyin
func T(args ...string) string {

  splitter, argsLen := " ", len(args)

  if argsLen == 0 {
    return ""
  }

  var str string
  var non = ""
  var fslice = []string{}

  if argsLen > 1 {
    splitter = args[argsLen - 1]
  }

  for _, rune := range args[0] {

    str = string(rune)

    if str == " " {
      continue
    }

    if hzRegexp.MatchString(str) { //chinese
      if non != ""{
        fslice = append(fslice, non)
        non = ""
      }
      fslice = append(fslice, strings.Split(PinyinDict[str], ",")[0])
    }else{
      non += str
    }
  }

  if non != ""{
    fslice = append(fslice, non)
  }

  return strings.Join(fslice, splitter)
}