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

	n, k := nextInt(), nextInt()
	a := nextIntSlice(n)

	ans := solve(n, k, a)

	Print(ans)
}

func solve(n, k int, a []int) int {
	var sum int
	for _, ai := range a {
		sum ^= ai
	}
	var swapped bool
	if k > n-k {
		k = n - k
		swapped = true
	}

	var ans int
	var dfs func(i, d, s int)
	dfs = func(i, d, s int) {
		if d == k {
			if swapped {
				ans = Max(ans, sum^s)
			} else {
				ans = Max(ans, s)
			}
			return

		}
		if i == n {
			return
		}
		dfs(i+1, d+1, s^a[i])
		dfs(i+1, d, s)
	}
	dfs(0, 0, 0)
	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func nextIntSlice(n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = nextInt()
	}
	return s
}

func Print(x any) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
