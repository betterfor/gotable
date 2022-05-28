package gotable

import (
	"regexp"

	"github.com/mattn/go-runewidth"
)

// 匹配颜色字符
var ansi = regexp.MustCompile(`(\x9B|\x1B\[)[0-?]*[ -/]*[@-~]`)

func RealWidth(str string) int {
	return runewidth.StringWidth(ansi.ReplaceAllString(str, ""))
}
