package gotable

import "io"

type Table struct {
	out io.Writer

	widthMap  map[int]int // 每列宽度最大集合
	heightMap map[int]int // 每行高度最大集合

	// 支持多行显示
	lines   [][][]string // len(lines)行数
	headers [][]string
	footers [][]string

	columnNum int // 列的数量
}

func NewTable(w io.Writer) *Table {
	return &Table{
		out:       w,
		widthMap:  make(map[int]int),
		heightMap: make(map[int]int),
		lines:     [][][]string{},
		headers:   [][]string{},
		footers:   [][]string{},
		columnNum: 0,
	}
}

func (t *Table) AddHeader(words []string) {

}
