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

	n, k, d := nextInt(), nextInt(), nextInt()
	a := nextIntSlice(n)

	ans, err := solve(n, k, d, a)

	if err != nil {
		PrintInt(-1)
	} else {
		PrintInt(ans)
	}
}

func solve(n, k, d int, a []int) (int, error) {
	//i番目の値まで見て、j個選び、mod dがkになる値の最大値
	dp := make([][][]int, n+10)
	for i := range dp {
		dp[i] = make([][]int, k+10)
		for j := range dp[i] {
			dp[i][j] = make([]int, d)
			for ii := range dp[i][j] {
				dp[i][j][ii] = -1
			}
		}
	}
	//fmt.Println(dp)
	dp[0][0][0] = 0
	for i := 0; i < n; i++ {
		for j := 0; j <= Min(i, k); j++ {
			for ii := 0; ii < d; ii++ {
				//i番目を選ばない
				dp[i+1][j][ii] = Max(dp[i+1][j][ii], dp[i][j][ii])
				if dp[i][j][ii] < 0 {
					continue
				}
				//i番目を選ぶ
				ni, nj := i+1, j+1

				nii := (ii + a[i]) % d
				dp[ni][nj][nii] = Max(dp[ni][nj][nii], dp[i][j][ii]+a[i])
			}
		}
	}
	return dp[n][k][0], nil
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

func PrintInt(x int) {
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
