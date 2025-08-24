package main

import (
	"fmt"
	"math"
	"strconv"
)

// Karatsuba乘法算法实现
func karatsuba(x, y int64) int64 {
	if x < 10 || y < 10 {
		return x * y
	}

	n := max(digits(x), digits(y))
	m := n / 2

	high1, low1 := split(x, m)
	high2, low2 := split(y, m)

	z0 := karatsuba(low1, low2)
	z1 := karatsuba((low1 + high1), (low2 + high2))
	z2 := karatsuba(high1, high2)

	return z2*int64(math.Pow10(2*m)) + (z1-z2-z0)*int64(math.Pow10(m)) + z0
}

func digits(n int64) int {
	return len(strconv.FormatInt(n, 10))
}

func split(n int64, m int) (int64, int64) {
	divisor := int64(math.Pow10(m))
	high := n / divisor
	low := n % divisor
	return high, low
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var a, b int64
	fmt.Print("请输入第一个数字: ")
	fmt.Scan(&a)
	fmt.Print("请输入第二个数字: ")
	fmt.Scan(&b)

	result := karatsuba(a, b)
	fmt.Printf("Karatsuba乘法结果: %d\n", result)
	fmt.Printf("标准乘法结果: %d\n", a*b)

}

