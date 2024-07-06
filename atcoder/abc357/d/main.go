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

	n := nextInt()

	ans := solve(n)

	Print(ans)
}

func dig(x int) int {
	if x == 0 {
		return 1
	}
	var res int
	for x > 0 {
		res++
		x /= 10
	}
	return res
}

func solve(n int) int {
	const p = 998244353
	d := dig(n)
	//TODO:ここを高速化
	r := Pow(10, d, p)
	w := (Pow(r, n, p) - 1 + p) % p
	w *= Inv((r-1+p)%p, p)
	w %= p

	ans := w * (n % p) % p
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

func Pow(x, y, p int) int {
	ret := 1
	for y > 0 {
		if y%2 == 1 {
			ret = ret * x % p
		}
		y >>= 1
		x = x * x % p
	}
	return ret
}

func Inv(x, p int) int {
	return Pow(x, p-2, p)
}
