package main

import (
	"bufio"
	"fmt"
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

	n, x, y := nextInt(), nextInt(), nextInt()
	var p, t []int
	for i := 0; i < n-1; i++ {
		p, t = append(p, nextInt()), append(t, nextInt())
	}

	lq := nextInt()
	var q []int
	for i := 0; i < lq; i++ {
		q = append(q, nextInt())
	}

	ans := solve(n, x, y, p, t, lq, q)

	PrintVertically(ans)
}

func solve(n, x, y int, p, t []int, lq int, q []int) []int {
	lcm := 1
	for i := 2; i <= 8; i++ {
		lcm = Lcm(lcm, i)
	}

	var s []int
	for mod := 0; mod < lcm; mod++ {
		v := mod
		for i := 0; i < n-1; i++ {
			for j := 0; j < p[i]; j++ {
				if (v+j)%p[i] == 0 {
					v += j + t[i]
					break
				}
			}
		}
		s = append(s, v-mod)
	}

	var ans []int
	for _, qi := range q {
		v := qi + x + s[(qi+x)%lcm] + y
		ans = append(ans, v)
	}
	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func PrintVertically(x []int) {
	defer out.Flush()
	for _, v := range x {
		fmt.Fprintln(out, v)
	}
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
