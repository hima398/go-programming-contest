package main

import (
	"bufio"
	"fmt"
	"math"
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

	n := nextInt()
	a := nextIntSlice(n)

	ans := solve(n, a)

	Print(ans)
}

func solve(n int, a []int) int {
	const INF = math.MaxInt
	//先頭から順番に見た数列で、単調増加な部分列で長さiであるものの最小値
	//k==1のときは操作を実施済み
	dp := make([][]int, 2)
	for k := range dp {
		dp[k] = make([]int, n)
		for i := range dp[k] {
			dp[k][i] = INF
		}
	}

	dp[0][0] = a[0]
	dp[1][0] = 0
	for i := 1; i < n; i++ {
		replaced := sort.Search(n, func(j int) bool {
			return a[i-1]+1 <= dp[0][j]
		})

		idx1 := sort.Search(n, func(j int) bool {
			return a[i] <= dp[0][j]
		})
		idx2 := sort.Search(n, func(j int) bool {
			return a[i] <= dp[1][j]
		})

		//置き換える操作を実施する
		dp[1][replaced] = Min(dp[1][replaced], a[i-1]+1)

		//置き換える操作をする
		dp[0][idx1] = a[i]
		dp[1][idx2] = Min(dp[1][idx2], a[i])
	}

	for _, v := range dp {
		fmt.Println(v)
	}

	var ans int
	for i, v := range dp[1] {
		if v < INF {
			ans = Max(ans, i+1)
		}
	}
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
