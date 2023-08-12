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
	var t, x, a []int
	for i := 0; i < n; i++ {
		t = append(t, nextInt())
		x = append(x, nextInt())
		a = append(a, nextInt())
	}
	ans := solve(n, t, x, a)
	PrintInt(ans)
}

func solve(n int, t, x, a []int) int {
	dp := make([][5]int, int(1e5)+1)
	mx, ma := make(map[int]int), make(map[int]int)
	for i, ti := range t {
		mx[ti] = x[i]
		ma[ti] = a[i]
	}
	for i := 1; i <= int(1e5); i++ {
		for pj := 0; pj < Min(i, 5); pj++ {
			for j := Max(pj-1, 0); j <= Min(pj+1, 4); j++ {
				if _, found := mx[i]; found {
					if Abs(mx[i]-j) == 0 {
						dp[i][j] = Max(dp[i][j], dp[i-1][pj]+ma[i])
					} else {
						dp[i][j] = Max(dp[i][j], dp[i-1][pj])
					}
				} else {
					dp[i][j] = Max(dp[i][j], dp[i-1][pj])
				}
			}
		}
	}
	var ans int
	for j := 0; j < 5; j++ {
		ans = Max(ans, dp[t[n-1]][j])
	}
	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func PrintInt(x int) {
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
