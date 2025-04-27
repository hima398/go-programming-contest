package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()

	ans := solve(n)

	Print(ans)

}

func pow(x, y int) (int, error) {
	res := 1
	for i := 0; i < y; i++ {
		res *= x
		if res > int(2e18) {
			return -1, errors.New("Overflow")
		}
		if res < 0 {
			return -1, errors.New("Overflow")
		}
	}
	return res, nil
}

func solve(n int) int {
	var ans int
	counted := make(map[int]struct{})
	for i := 3; i < 60; i++ {
		var cnt int
		for x := 2; ; x++ {
			xx, err := pow(x, i)
			if err != nil {
				break
			}
			if xx > n {
				break
			}

			//if xx == int(math.Sqrt(float64(xx)))*int(math.Sqrt(float64(xx))) {
			if xx == FloorSqrt(xx)*FloorSqrt(xx) {
				continue
			}
			if _, found := counted[xx]; found {
				continue
			}
			cnt++
			counted[xx] = struct{}{}
		}
		ans += cnt

	}

	ans += FloorSqrt(n) //int(math.Sqrt(float64(n)))
	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func Print(x any) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func Floor(x, y int) int {
	/*
		return x / y
	*/
	r := (x%y + y) % y
	return (x - r) / y
}

func Ceil(x, y int) int {
	return (x + y - 1) / y
}

func Sqrt(x int) int {
	x2 := int(math.Sqrt(float64(x))) - 1
	for (x2+1)*(x2+1) <= x {
		x2++
	}
	return x2
}

// xの平方根の整数部分を計算する
func FloorSqrt(x int) int {
	l, r := 0, int(2e9)
	for l <= r {
		mid := (l + r) / 2
		if mid*mid > x {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return r
}

// xのy乗根の整数部分を計算する
func FloorNthRoot(x, y int) int {
	var res int
	return res
}

func Gcd(x, y int) int {
	if x == 0 {
		return y
	}
	if y == 0 {
		return x
	}
	/*
		if x < y {
			x, y = y, x
		}
	*/
	return Gcd(y, x%y)
}

func Lcm(x, y int) int {
	// x*yのオーバーフロー対策のため先にGcdで割る
	// Gcd(x, y)はxの約数のため割り切れる
	ret := x / Gcd(x, y)
	ret *= y
	return ret
}
