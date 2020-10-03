package main

import (
	"fmt"
)

type ErrNegativeSqrt float64 // エラー分用に新しい型を定義

func (e ErrNegativeSqrt) Error() string { // ErrNegativeSqrtをレシーバーに設定
	return fmt.Sprintf("cannot Sqrt negative number: %g", e) 
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	return 0, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
