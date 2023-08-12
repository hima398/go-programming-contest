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

	ln := nextInt()
	n := 2 * ln
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			a[i][j] = nextInt()
		}
	}
	ans := solve(n, a)
	PrintInt(ans)
}

func solve(n int, a [][]int) int {
	mask := (1 << n) - 1

	var ans int
	var dfs func(i, pat, score int)
	dfs = func(i, pat, score int) {
		//全員ペアが組めたら楽しさの合計を使って解を更新
		if pat == mask {
			ans = Max(ans, score)
			return
		}

		//i番目の人がまだペアを作れていないとき
		if pat>>i&1 == 0 {
			for j := i + 1; j < n; j++ {
				if pat>>j&1 == 0 {
					dfs(i+1, pat|(1<<i)|(1<<j), score^a[i][j])
				}
			}
		} else {
			//i番目の人がすでにペアが作れたとき
			dfs(i+1, pat, score)
		}
	}
	dfs(0, 0, 0)
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

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
