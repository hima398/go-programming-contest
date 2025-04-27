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
	a := nextIntSlice(n)
	m := nextInt()
	b := nextIntSlice(m)
	x := nextInt()

	ans := solve(n, a, m, b, x)

	PrintString(ans)
}

func solve(n int, a []int, m int, b []int, x int) string {
	riceCakes := make(map[int]struct{})
	for _, v := range b {
		riceCakes[v] = struct{}{}
	}
	dp := make([]bool, x+1)
	dp[0] = true
	for i := 0; i < x; i++ {
		for j := 0; j < n; j++ {
			next := i + a[j]
			if next > x {
				continue
			}
			if _, found := riceCakes[next]; found {
				continue
			}
			dp[next] = dp[next] || dp[i]
		}
	}
	if dp[x] {
		return "Yes"
	} else {
		return "No"
	}
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

func PrintString(x string) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
