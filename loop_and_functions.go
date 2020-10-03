package main

import (
	"fmt"
	"math"
)

//閾値を使用する方法
func Sqrt_check(x float64) float64 {
	z := 0.1
	delta := 0.001
	for {
		next := z - (z*z-x)/2*z //次の値を確認
		if math.Abs(next-z) < delta { //次の値と、今のzの値の差の絶対値が、閾値よりも小さかったらbreak
			break
		}
		z = next
	}
	return z
}


// 10回ループを回す方法
func Sqrt(x float64) float64 {
	z := 1.0 // 浮動小数点の変数を初期化して宣言する
	for i := 0; i < 10; i++ {
		z = z - (z*z-x)/2*z
	}
	return z
}

func main() {
	fmt.Println(Sqrt(4))
	fmt.Println(Sqrt_check(4))
}
