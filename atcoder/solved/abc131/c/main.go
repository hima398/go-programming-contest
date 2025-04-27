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

	a, b, c, d := nextInt(), nextInt(), nextInt(), nextInt()
	ans := solve(a, b, c, d)
	PrintInt(ans)
}

func solve(a, b, c, d int) int {
	//a-1以下の整数のうちCでもDでも割り切れないものの個数
	an := (a - 1) - ((a-1)/c + (a-1)/d - (a-1)/Lcm(c, d))
	//b以下の整数のうちCでもDでも割り切れないものの個数
	bn := b - (b/c + b/d - b/Lcm(c, d))
	ans := bn - an
	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
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
