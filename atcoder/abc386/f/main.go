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

	k := nextInt()
	s := nextString()
	t := nextString()

	//if solve01(k, s, t) {
	if solve(k, s, t) {
		Print("Yes")
	} else {
		Print("No")
	}
}

func solve(k int, s, t string) bool {
	const INF = 1 << 60
	if Abs(len(s)-len(t)) > k {
		return false
	}

	dp := make([][]int, len(s)+1)
	for i := range dp {
		dp[i] = make([]int, 2*k+1)
		for j := range dp[i] {
			dp[i][j] = INF
		}
	}
	//for i, v := range dp {
	//	fmt.Println(i, v)
	//}
	dp[0][k] = 0

	for si := 0; si <= len(s); si++ {
		for ti := si - k; ti <= si+k; ti++ {
			if ti < 0 {
				continue
			}
			if ti > len(t) {
				break
			}
			j := ti - si + k

			if si > 0 && j < 2*k {
				dp[si][j] = Min(dp[si][j], dp[si-1][j+1]+1)
			}
			if ti > 0 && j > 0 {
				dp[si][j] = Min(dp[si][j], dp[si][j-1]+1)
			}
			if si > 0 && ti > 0 {
				var d int
				if s[si-1] != t[ti-1] {
					d = 1
				}
				dp[si][j] = Min(dp[si][j], dp[si-1][j]+d)
			}
		}
	}
	//for i, v := range dp {
	//	fmt.Println(i, v)
	//}

	dist := dp[len(s)][len(t)-len(s)+k]
	//fmt.Println("dist = ", dist)
	return dist <= k
}

func solve01(k int, s, t string) bool {
	dist := computeLevenshteinDist(s, t)
	//fmt.Println("dist = ", dist)
	return dist <= k
}

func computeLevenshteinDist(s, t string) int {
	const INF = 1 << 60
	dp := make([][]int, len(s)+1)
	for i := range dp {
		dp[i] = make([]int, len(t)+1)
		for j := range dp[i] {
			dp[i][j] = INF
		}
	}
	dp[0][0] = 0

	for i := 1; i <= len(s); i++ {
		dp[i][0] = i
	}
	for j := 1; j <= len(t); j++ {
		dp[0][j] = j
	}

	for i := 0; i < len(s); i++ {
		for j := 0; j < len(t); j++ {
			var d int
			if s[i] != t[j] {
				d = 1
			}
			dp[i+1][j+1] = Min(Min(dp[i][j+1]+1, dp[i+1][j]+1), dp[i][j]+d)
		}
	}
	for _, v := range dp {
		fmt.Println(v)
	}
	return dp[len(s)][len(t)]
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

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
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
