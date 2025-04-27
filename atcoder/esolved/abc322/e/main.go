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

	n, k, p := nextInt(), nextInt(), nextInt()
	var c []int
	var a [][]int
	for i := 0; i < n; i++ {
		c = append(c, nextInt())
		a = append(a, nextIntSlice(k))
	}

	ans, err := solve(k, n, p, c, a)

	if err != nil {
		Print(-1)
	} else {
		Print(ans)
	}
}

func solve(k, n, p int, c []int, a [][]int) (int, error) {
	const INF = 1 << 60
	//開発案xまで見てパラメータの集合をyにする最小のコスト
	dp := make([]map[int]int, n+1)
	for x := range dp {
		dp[x] = make(map[int]int)
	}
	dp[0][0] = 0

	for x := 0; x < n; x++ {
		for params, cost := range dp[x] {
			//x番目の開発案を実行しない
			if _, found := dp[x+1][params]; !found {
				dp[x+1][params] = cost
			} else {
				dp[x+1][params] = Min(dp[x+1][params], cost)
			}

			//x番目の開発案を実行する
			var nextParams int
			for i := 0; i < k; i++ {
				param := params >> (i * 3) & 0b111
				nextParam := Min(param+a[x][i], p)
				nextParams |= nextParam << (i * 3)
			}
			if _, found := dp[x+1][nextParams]; !found {
				dp[x+1][nextParams] = cost + c[x]
			} else {
				dp[x+1][nextParams] = Min(dp[x+1][nextParams], cost+c[x])
			}
		}
	}
	var maxParams int
	for i := 0; i < k; i++ {
		maxParams |= p << (i * 3)
	}
	if _, found := dp[n][maxParams]; !found {
		return -1, errors.New("Impossible")
	}
	return dp[n][maxParams], nil
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
