package pinyin

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
