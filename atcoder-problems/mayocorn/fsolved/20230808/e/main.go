package main

import (
	"bufio"
	"errors"
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

	n, m, k := nextInt(), nextInt(), nextInt()
	var a, b, c []int
	for i := 0; i < m; i++ {
		a = append(a, nextInt()-1)
		b = append(b, nextInt()-1)
		c = append(c, nextInt())
	}
	e := nextIntSlice(k)
	for i := range e {
		e[i]--
	}

	ans, err := solve(n, m, k, a, b, c, e)
	if err != nil {
		PrintInt(-1)
	} else {
		PrintInt(ans)
	}
}

func solve(n, m, k int, a, b, c, e []int) (int, error) {
	const INF = 1 << 60
	//良い経路を辿って頂点iにたどり着くまでに通る道の長さの最小値
	dp := make([]int, n)
	for i := range dp {
		dp[i] = INF
	}
	dp[0] = 0

	for i := 0; i < k; i++ {
		dp[b[e[i]]] = Min(dp[b[e[i]]], dp[a[e[i]]]+c[e[i]])
	}

	if dp[n-1] == INF {
		return -1, errors.New("Impossible")
	} else {
		return dp[n-1], nil
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
