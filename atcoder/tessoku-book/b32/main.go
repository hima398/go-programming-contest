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
	a := nextIntSlice(k)

	firstCanWin := solve(n, k, a)
	//firstCanWin := solveRecursively(n, k, a)

	if firstCanWin {
		PrintString("First")
	} else {
		PrintString("Second")
	}
}

func solve(n, k int, a []int) bool {
	dp := make([]bool, n+1)

	for i := 0; i < n; i++ {
		for j := 0; j < k; j++ {
			next := i + a[j]
			if next > n {
				continue
			}
			dp[next] = dp[next] || !dp[i]
		}
	}
	return dp[n]
}

func solveRecursively(n, k int, a []int) bool {
	//memo := make([]bool, n+1)
	memo := make(map[int]bool)
	//石がx個積まれている時、先手が勝てるかを返す関数
	var f func(x int) bool
	f = func(x int) bool {
		if _, found := memo[x]; found {
			return memo[x]
		}
		if x == 0 {
			memo[0] = false
			return memo[0]
		}
		canWin := false
		for i := 0; i < k; i++ {
			next := x - a[i]
			if next < 0 {
				continue
			}
			//石がnext個積まれている時、先手が負けるものを相手に渡せるのであれば勝てる
			canWin = canWin || !f(next)
		}
		memo[x] = canWin
		return memo[x]
	}

	return f(n)
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

func PrintString(x string) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
