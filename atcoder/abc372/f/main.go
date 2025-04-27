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

	n, m, k := nextInt(), nextInt(), nextInt()
	var x, y []int
	for i := 0; i < m; i++ {
		x = append(x, nextInt()-1)
		y = append(y, nextInt()-1)
	}

	ans := solve(n, m, k, x, y)

	Print(ans)
}

func solve(n, m, k int, x, y []int) int {
	const p = 998244353
	e := make([][]int, n)
	for i := 0; i < n; i++ {
		e[i] = append(e[i], (i+1)%n)
	}
	for i := 0; i < m; i++ {
		e[x[i]] = append(e[x[i]], y[i])
	}
	dp := make([]map[int]int, k+1)
	dp[0] = make(map[int]int)
	dp[0][0] = 1
	for i := 0; i < k; i++ {
		dp[i+1] = make(map[int]int)
		for j, v := range dp[i] {
			for _, next := range e[j] {
				dp[i+1][next] += v
				dp[i+1][next] %= p
			}
		}
		if i <= 10 {
			fmt.Println(i+1, dp[i+1])
		}
	}
	var ans int
	for _, v := range dp[k] {
		ans += v
		ans %= p
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
