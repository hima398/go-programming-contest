package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, t := nextInt(), nextInt()
	var a, b []int
	for i := 0; i < n; i++ {
		a, b = append(a, nextInt()), append(b, nextInt())
	}

	ans := solve(n, t, a, b)

	Print(ans)
}

func solve(n, t int, a, b []int) int {
	type dish struct {
		t, d int
	}
	var dishes []dish
	for i := 0; i < n; i++ {
		dishes = append(dishes, dish{a[i], b[i]})
	}
	sort.Slice(dishes, func(i, j int) bool {
		return dishes[i].t < dishes[j].t
	})

	//i番目の料理でj分以内に完食できる美味しさの最大
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, t+1)
	}

	var ans int
	for i := 0; i < n; i++ {
		for j := 0; j < t; j++ {
			dp[i+1][j] = Max(dp[i+1][j], dp[i][j])
			next := j + dishes[i].t
			if next >= t {
				continue
			}
			dp[i+1][next] = Max(dp[i+1][next], dp[i][j]+dishes[i].d)
		}
		ans = Max(ans, dp[i][t-1]+dishes[i].d)
	}
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

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
