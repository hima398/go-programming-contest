package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func reverse(x int) int {
	var res int
	for x > 0 {
		sx := x % 10
		res = res*10 + sx
		x /= 10
	}
	return res
}

func F(x int) int {
	res := x
	rx := x
	for i := 0; i < 20; i++ {
		rx = reverse(rx)
		res = Min(res, rx)
	}
	return res
}

func solveHonestly(n, k int) (ans int) {
	for i := 1; i <= n; i++ {
		if F(i) == k {
			ans++
		}
	}
	return ans
}

func solve(n, k int) (ans int) {
	if k%10 == 0 {
		return 0
	}
	if k == reverse(k) {
		x := k
		for x <= n {
			ans++
			x *= 10
		}
		return ans
	} else {
		mn, mx := k, reverse(k)
		if mn > mx {
			return 0
		}
		for mn <= n {
			ans++
			if mx <= n {
				ans++
			}
			mn *= 10
			mx *= 10
		}
		return ans
	}
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, k := nextInt(), nextInt()
	ans := solve(n, k)
	PrintInt(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
