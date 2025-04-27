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
	a := nextIntSlice(n)

	ans := solve(n, a)

	PrintHorizonaly(ans)
}

func solve(n int, a []int) []int {
	const p = 998244353
	//i番目まで見て、長さ1、初項a1の部分列の個数
	dp1 := make([]map[int]int, n+1)
	//i番目まで見て、長さj、初項a1、2番目の項a2の部分列の個数
	dp2 := make([][]map[int]map[int]int, n+1)
	for i := range dp2 {
		dp2[i] = make([]map[int]map[int]int, n+1)
		for j := range dp2 {
			dp2[i][j] = make(map[int]map[int]int)
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j <= n; j++ {
			//i番目の数字を採用しない場合
			if j == 0 {
				for a1 := range dp1[i] {
					if dp1[i+1] == nil {
						dp1[i+1] = make(map[int]int)
					}
					dp1[i+1][a1] += dp1[i][a1]
					dp1[i+1][a1] %= p
				}
			} else {
				for a1 := range dp2[i][j] {
					for a2 := range dp2[i][j][a1] {
						if dp2[i+1][j][a1] == nil {
							dp2[i+1][j][a1] = make(map[int]int)
						}
						dp2[i+1][j][a1][a2] += dp2[i][j][a1][a2]
						dp2[i+1][j][a1][a2] %= p
					}
				}
			}
			//i番目の数字を採用する
			if j == 0 {
				if dp1[i+1] == nil {
					dp1[i+1] = make(map[int]int)
				}
				dp1[i+1][a[i]]++
				dp1[i+1][a[i]] %= p
			} else if j == 1 {
				for a1 := range dp1[i] {
					if dp2[i+1][j+1][a1] == nil {
						dp2[i+1][j+1][a1] = make(map[int]int)
					}
					dp2[i+1][j+1][a1][a[i]] += dp1[i][a1]
					dp2[i+1][j+1][a1][a[i]] %= p
				}
			} else { // j>=2
				for a1 := range dp2[i][j] {
					for a2 := range dp2[i][j][a1] {
						//a[i]を部分列に採用できるかチェックする
						if a1+j*(a2-a1) != a[i] {
							continue
						}
						if dp2[i+1][j+1][a1] == nil {
							dp2[i+1][j+1][a1] = make(map[int]int)
						}
						dp2[i+1][j+1][a1][a2] += dp2[i][j][a1][a2]
						dp2[i+1][j+1][a1][a2] %= p
					}
				}
			}
		}
	}
	//fmt.Println(dp1)
	//fmt.Println(dp2[n])
	ans := make([]int, n)
	for _, v := range dp1[n] {
		ans[0] += v
		ans[0] %= p
	}
	for j := 2; j <= n; j++ {
		for a1 := range dp2[n][j] {
			for a2 := range dp2[n][j][a1] {
				ans[j-1] += dp2[n][j][a1][a2]
				ans[j-1] %= p
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

func PrintHorizonaly(x []int) {
	defer out.Flush()
	fmt.Fprintf(out, "%d", x[0])
	for i := 1; i < len(x); i++ {
		fmt.Fprintf(out, " %d", x[i])
	}
	fmt.Fprintln(out)
}
