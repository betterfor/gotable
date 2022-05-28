package gotable

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	ESC = "\033"
	SEP = ";"
)

// Background color
const (
	BgBlack int = 40 + iota
	BgRed
	BgGreen
	BgYellow
	BgBlue
	BgMagenta
	BgCyan
	ByWhite
)

// Foreground color
const (
	FgBlack int = 30 + iota
	FgRed
	FgGreen
	FgYellow
	FgMagenta
	FgCyan
	FgWhite
)

// Background Hi color
const (
	BgHiBlack int = 100 + iota
	BgHiRed
	BgHiGreen
	BgHiYellow
	BgHiBlue
	BgHiMagenta
	BgHiCyan
	BgHiWhite
)

// Foreground Hi color
const (
	FgHiBlack int = iota + 90
	FgHiRed
	FgHiGreen
	FgHiYellow
	FgHiBlue
	FgHiMagenta
	FgHiCyan
	FgHiWhite
)

const (
	Normal          = 0
	Bold            = 1
	UnderlineSingle = 4
	Italic
)

func startColor(seq string) string {
	return fmt.Sprintf("%s[%sm", ESC, seq)
}

func stopColor() string {
	return fmt.Sprintf("%s[%dm", ESC, Normal)
}

// Colors
type Colors []int

// format change the color of string
func Format(s string, codes interface{}) string {
	var seq string
	switch v := codes.(type) {
	case string:
		seq = v
	case []int:
		seq = makeSeqence(v)
	case Colors:
		seq = makeSeqence(v)
	default:
		return s
	}
	if len(seq) == 0 {
		return s
	}
	return startColor(seq) + s + stopColor()
}

func makeSeqence(codes []int) string {
	var codesString []string
	for _, code := range codes {
		codesString = append(codesString, strconv.Itoa(code))
	}
	return strings.Join(codesString, SEP)
}

func Color(colors ...int) Colors {
	return colors
}
