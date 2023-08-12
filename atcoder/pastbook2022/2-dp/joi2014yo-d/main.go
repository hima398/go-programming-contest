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
	s := nextString()
	ans := solve(n, s)
	PrintInt(ans)
}

func solve(n int, s string) int {
	const p = int(1e4) + 7
	shift := map[byte]int{'J': 0, 'O': 1, 'I': 2}
	//i日目まで見て、3人の出席パターンがjの時のスケジュール表のパターン
	dp := make([][8]int, n+1)
	dp[0][1] = 1
	for i := 0; i < n; i++ {
		for j := 1; j < 8; j++ {
			for nj := 1; nj < 8; nj++ {
				//鍵を持って帰った人が次の日にこないパターンなので
				//以降の処理を無視する
				if nj&j == 0 {
					continue
				}
				//責任者が出席していない
				if (nj>>shift[s[i]])&1 == 0 {
					continue
				}
				dp[i+1][nj] += dp[i][j]
				dp[i+1][nj] %= p
			}
		}
	}
	var ans int
	for j := 0; j < 8; j++ {
		ans += dp[n][j]
		ans %= p
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

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
