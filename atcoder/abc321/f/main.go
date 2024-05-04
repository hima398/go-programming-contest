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

	q, k := nextInt(), nextInt()
	var s []string
	var x []int
	for i := 0; i < q; i++ {
		s = append(s, nextString())
		x = append(x, nextInt())
	}
	ans := solve(q, k, s, x)
	for _, v := range ans {
		Print(v)
	}
}

func solve(q, k int, s []string, x []int) []int {
	const p = 998244353
	dp := make([]int, k+1)
	var ans []int
	dp[0] = 1
	for i := 0; i < q; i++ {
		switch s[i] {
		case "+":
			for j := k - x[i]; j >= 0; j-- {
				nj := j + x[i]
				dp[nj] += dp[j]
				dp[nj] %= p
			}
		case "-":
			for j := 0; j <= k-x[i]; j++ {
				nj := j + x[i]
				dp[nj] += p - dp[j]
				dp[nj] %= p
			}
		}
		//fmt.Println(dp)
		ans = append(ans, dp[k])
	}
	return ans
}

func firstsolve(q, k int, s []string, x []int) []int {
	const p = 998244353
	dp := make([][]int, q+1)
	for i := range dp {
		dp[i] = make([]int, k+1) //make(map[int]int)
	}
	var ans []int
	dp[0][0] = 1
	for i := 0; i < q; i++ {
		copy(dp[i+1], dp[i])
		switch s[i] {
		case "+":
			for j := 0; j <= k; j++ {
				nj := j + x[i]
				if nj > k {
					continue
				}
				dp[i+1][nj] += dp[i][j]
				dp[i+1][nj] %= p
			}
		case "-":
			for j := x[i]; j <= k; j++ {
				nj := j - x[i]
				dp[i+1][nj] += p - dp[i][j]
				dp[i+1][nj] %= p
			}
		}
		fmt.Println(dp[i+1])
		ans = append(ans, dp[i+1][k])
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
