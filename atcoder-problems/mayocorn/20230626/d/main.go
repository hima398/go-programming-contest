package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, s := nextInt(), nextInt()
	var a, b []int
	for i := 0; i < n; i++ {
		a = append(a, nextInt())
		b = append(b, nextInt())
	}
	ans, ok := solve(n, s, a, b)
	if ok {
		PrintString("Yes")
		PrintString(ans)
	} else {
		PrintString("No")
	}
}

func solve(n, s int, a, b []int) (string, bool) {
	//i番目のカードまで見て、合計をjにできるかどうか
	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, s+1)
	}
	dp[0][0] = true
	for i := 0; i < n; i++ {
		for j := 0; j <= s; j++ {
			if !dp[i][j] {
				continue
			}
			nj := j + a[i]
			if nj <= s {
				dp[i+1][nj] = true
			}
			nj = j + b[i]
			if nj <= s {
				dp[i+1][nj] = true
			}
		}
	}
	if !dp[n][s] {
		return "", false
	}
	var ans []string
	cur := s
	for i := n - 1; i >= 0; i-- {
		if cur-a[i] >= 0 && dp[i][cur-a[i]] {
			ans = append(ans, "H")
			cur -= a[i]
		} else if cur-b[i] >= 0 && dp[i][cur-b[i]] {
			ans = append(ans, "T")
			cur -= b[i]
		}
	}
	for i := 0; i < len(ans)/2; i++ {
		j := len(ans) - 1 - i
		ans[i], ans[j] = ans[j], ans[i]
	}
	return strings.Join(ans, ""), true
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
