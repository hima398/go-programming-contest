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

	//solveHonestly(n, k, a)
	ans := solve(n, k, a)

	Print(ans)
}

func solve(n, k int, a []int) int {
	const p = 998244353

	s := make([]int, n+1)
	for i := 0; i < n; i++ {
		s[i+1] = s[i] + a[i]
	}
	//i番目まで見て、部分列の合計がkになるものを含まない個数
	dp := make([]int, n+1)
	dp[0] = 1
	sum := 1
	//キーがs[i]となるdp[i]の総和
	m := make(map[int]int)
	m[0] = 1

	for i := 1; i <= n; i++ {
		dp[i] = (sum - m[s[i]-k]) % p
		m[s[i]] += dp[i]
		sum += dp[i]
	}
	//fmt.Println(dp)

	return dp[n]
}

func solveHonestly(n, k int, a []int) int {
	const p = 998244353

	s := make([]int, n+1)
	for i := 0; i < n; i++ {
		s[i+1] = s[i] + a[i]
	}
	//i番目まで見て、部分列の合計がkになるものを含まない個数
	dp := make([]int, n+1)
	dp[0] = 1

	for i := 0; i <= n; i++ {
		for j := 0; j < i; j++ {
			if s[i]-s[j] == k {
				continue
			}
			dp[i] += dp[j]
			dp[i] %= p
		}
	}
	//fmt.Println(dp)

	return dp[n]
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
