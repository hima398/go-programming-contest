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

	h, w, k := nextInt(), nextInt(), nextInt()
	si, sj := nextInt()-1, nextInt()-1
	var a [][]int
	for i := 0; i < h; i++ {
		a = append(a, nextIntSlice(w))
	}
	ans := solve(h, w, k, si, sj, a)
	Print(ans)
}

func solve(h, w, k, si, sj int, a [][]int) int {
	const INF = -(1 << 60)
	di := []int{-1, 0, 1, 0}
	dj := []int{0, -1, 0, 1}

	dp := make([][][]int, h*w+1)
	for l := range dp {
		dp[l] = make([][]int, h)
		for i := range dp[l] {
			dp[l][i] = make([]int, w)
			for j := range dp[l][i] {
				dp[l][i][j] = INF
			}
		}
	}
	dp[0][si][sj] = 0
	//fmt.Println(dp[0])

	for l := 0; l < Min(h*w, k); l++ {
		for i := 0; i < h; i++ {
			for j := 0; j < w; j++ {
				dp[l+1][i][j] = Max(dp[l+1][i][j], dp[l][i][j]+a[i][j])
				if dp[l][i][j] == INF {
					continue
				}
				for dir := 0; dir < 4; dir++ {
					ni, nj := i+di[dir], j+dj[dir]
					if ni < 0 || ni >= h || nj < 0 || nj >= w {
						continue
					}
					dp[l+1][ni][nj] = Max(dp[l+1][ni][nj], dp[l][i][j]+a[ni][nj])
				}
			}
		}
	}

	var ans int
	for l := 0; l <= Min(h*w, k); l++ {
		for i := 0; i < h; i++ {
			for j := 0; j < w; j++ {
				ans = Max(ans, dp[l][i][j]+(k-l)*a[i][j])
			}
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
