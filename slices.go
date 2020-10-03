package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 { // 返す値は2次元配列
	m := make([][]uint8, dy) // 長さdyの可変長配列を作成
	for i := 0; i < dy; i++ { // dy分ループを回す
		m[i] = make([]uint8, dx) // 長さdxの可変長配列を作成
		for j := 0; j < dx; j++ { //dx分ループを回す
			m[i][j] = uint8(i+j)
		}
	}
	return m
}

func main() {
	pic.Show(Pic)
}
