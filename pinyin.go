package pinyin

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

var hzRegexp = regexp.MustCompile("^[\u4e00-\u9fa5]$")

// Translate chinese to pinyin
func T(args ...string) string {
	return translate(args, false)
}

func TT(args ...string) string {
	return translate(args, true)
}

func translate(args []string, tone bool) string {

	splitter, argsLen := " ", len(args)

	if argsLen == 0 {
		return ""
	}

	var str string
	var non = ""
	var fslice = []string{}

	if argsLen > 1 {
		splitter = args[argsLen-1]
	}

	for _, rune := range args[0] {

		str = string(rune)

		if str == " " {
			continue
		}

		if hzRegexp.MatchString(str) { //chinese
			if non != "" {
				fslice = append(fslice, non)
				non = ""
			}

			fslice = append(fslice, spell(str, tone))

		} else {
			non += str
		}
	}

	if non != "" {
		fslice = append(fslice, non)
	}

	return strings.Join(fslice, splitter)
}

func spell(str string, tone bool) string {

	spell := strings.Split(PinyinDict[str], ",")[0]

	if tone {
		return spell
	}

	output := make([]string, utf8.RuneCountInString(spell))

	count := 0

	for _, tone := range spell {
		neutral, found := tones[string(tone)]
		if found {
			output[count] = neutral
		} else {
			output[count] = string(tone)
		}
		count++
	}

	return strings.Join(output, "")

}
