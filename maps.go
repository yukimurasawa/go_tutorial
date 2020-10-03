package main

import (
	"strings"
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int) // keyの型はstring型、valueの型はint型であるmapを作成
	for _, f := range strings.Fields(s) { // strings.Fields(s)で引数の文字列を単語で区切った配列を取得する。その配列に対してrangeでループを回す。
		m[f] += 1
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
