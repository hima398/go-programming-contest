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

	n, x := nextInt(), nextInt()
	var a, b []int
	for i := 0; i < n; i++ {
		a = append(a, nextInt())
		b = append(b, nextInt())
	}
	ok := solve(n, x, a, b)
	if ok {
		PrintString("Yes")
	} else {
		PrintString("No")
	}
}

func solve(n, x int, a, b []int) bool {
	//i番目の硬貨まで見て、j円払えるか？
	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, x+1)
	}
	dp[0][0] = true
	for i := 0; i < n; i++ {
		for j := 0; j <= b[i]; j++ {
			for k := 0; k <= x; k++ {
				nk := k + a[i]*j
				if nk > x {
					continue
				}
				dp[i+1][nk] = dp[i+1][nk] || dp[i][k]
			}
		}
	}
	return dp[n][x]
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func PrintString(x string) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
