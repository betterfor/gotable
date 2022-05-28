package gotable

import (
	"math"
	"strings"

	"github.com/mattn/go-runewidth"
)

func WrapString(text string, limit int) ([]string, int) {
	words := strings.Fields(strings.ReplaceAll(text, "\n", " "))
	var lines []string
	for _, w := range words {
		max := runewidth.StringWidth(w)
		if max > limit {
			limit = max
		}
	}
	for _, line := range shortest(words, limit) {
		lines = append(lines, strings.Join(line, " "))
	}
	return lines, limit
}

// 把多个单词组合成多行字符串
// https://xxyxyz.org/line-breaking/
// Wrap formats text at the given width in linear time
func shortest(words []string, width int) (ans [][]string) {
	count := len(words)
	offsets := make([]int, count+1)
	for i := 1; i < count+1; i++ {
		offsets[i] += offsets[i-1] + len(words[i-1])
	}

	var minima = make([]int, count+1)
	var breaks = make([]int, count+1)
	for i := 1; i <= count; i++ {
		minima[i] = math.MaxInt32
	}

	for i := 0; i < count; i++ {
		for j := i + 1; j <= count; j++ {
			w := offsets[j] - offsets[i] + j - i - 1
			if w > width {
				break
			}
			cost := minima[i] + (width-w)*(width-w)
			if cost < minima[j] {
				minima[j] = cost
				breaks[j] = i
			}
		}
	}
	j := count
	for j > 0 {
		i := breaks[j]
		ans = append(ans, words[i:j])
		j = i
	}
	for i, j := 0, len(ans)-1; i < j; i, j = i+1, j-1 {
		ans[i], ans[j] = ans[j], ans[i]
	}
	return
}
