package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	t := nextString()
	n := nextInt()
	var a []int
	var s [][]string
	for i := 0; i < n; i++ {
		a = append(a, nextInt())
		var ss []string
		for j := 0; j < a[i]; j++ {
			ss = append(ss, nextString())
		}
		s = append(s, ss)
	}

	ans := solve(t, n, a, s)

	Print(ans)
}

func solve(t string, n int, a []int, s [][]string) int {
	const INF = math.MaxInt - 1
	//袋iまで見てtのj文字目まで一致させるのに最小の金額
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, len(t)+1)
		for j := range dp[i] {
			dp[i][j] = INF
		}
	}
	dp[0][0] = 0

	//O(T*T*Max(Ai)) ~ int(1e5)
	for i := 0; i < n; i++ {
		//何もしない場合
		copy(dp[i+1], dp[i])
		//袋iから使える文字を使う
		for j := 0; j < len(t); j++ {
			for _, ss := range s[i] {
				tt := t[j:]
				if !strings.HasPrefix(tt, ss) {
					continue
				}
				nj := j + len(ss)
				dp[i+1][nj] = Min(dp[i+1][nj], dp[i][j]+1)
			}
		}
	}
	ans := dp[n][len(t)]
	if ans == INF {
		ans = -1
	}
	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func nextString() string {
	sc.Scan()
	return sc.Text()
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
