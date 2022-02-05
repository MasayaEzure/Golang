package main

import (
	"fmt"
	"math"
)

func main() {
	// 円周率
	fmt.Println(math.Pi)

	// 2つの平方根
	fmt.Println(math.Sqrt2)

	// float32で表現可能な最大値
	fmt.Println(math.MaxFloat32)
	// float32で表現可能な0以外の最小値
	fmt.Println(math.SmallestNonzeroFloat32)
	// float64で表現可能な最大値
	fmt.Println(math.MaxFloat64)
	// float64で表現可能な0以外の最小値
	fmt.Println(math.SmallestNonzeroFloat64)
	// int64で表現可能な最大値
	fmt.Println(math.MaxInt64)
	// int64で表現可能な最小値
	fmt.Println(math.MinInt64)
	
	// 絶対値
	fmt.Println(math.Abs(100))
	fmt.Println(math.Abs(-100))
	
	// 累乗を求める
	fmt.Println(math.Pow(0,2))
	fmt.Println(math.Pow(2,0))
	
	// 平方根
	fmt.Println(math.Sqrt((2)))
	// 立方根
	fmt.Println(math.Cbrt((8)))
	
	// 最大値
	fmt.Println(math.Max(10,3))
	// 最小値
	fmt.Println(math.Min(10,3))

	// 小数点以下の切り捨て
	fmt.Println(math.Trunc(3.14))
	fmt.Println(math.Trunc(-3.14))

	// 引数の数値より小さい最大の整数を返す
	fmt.Println(math.Floor(3.5))
	fmt.Println(math.Floor(-3.5))

	// 引数の最大値より大きい最小の整数を返す
	fmt.Println(math.Ceil(2.8))
	fmt.Println(math.Ceil(-2.8))
}