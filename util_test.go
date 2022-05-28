package gotable

import (
	"testing"
	"unicode/utf8"
)

func TestRealWidth(t *testing.T) {
	testCases := []struct {
		arg  string
		want int
	}{
		{"123456789", 9},
		{"123456789\t", 9},
		{"测试长度", 8},
		{"测试长度\t", 8},
		{"测试长度", utf8.RuneCountInString("测试长度") * 2},
		{"1混合字符2", 10},
		{"1混合字符2\n", 10},
	}

	for _, v := range testCases {
		if e := RealWidth(v.arg); e != v.want {
			t.Errorf("%s want %d, but %d\n", v.arg, v.want, e)
		}
	}
}
