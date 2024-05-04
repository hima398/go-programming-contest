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

	a, b, c := nextInt(), nextInt(), nextInt()

	ans := solve(a, b, c)

	Print(ans)
}

func solve(a, b, c int) float64 {
	const n = 105 //100枚になるまで期待値を求めるのに十分な値
	var visited [n][n][n]bool
	var memo [n][n][n]float64

	var f func(a, b, c int) float64
	f = func(a, b, c int) float64 {
		if visited[a][b][c] {
			return memo[a][b][c]
		}
		if a >= 100 || b >= 100 || c >= 100 {
			visited[a][b][c] = true
			memo[a][b][c] = 0
			return memo[a][b][c]
		}
		fa, fb, fc := float64(a), float64(b), float64(c)
		fs := fa + fb + fc
		visited[a][b][c] = true
		memo[a][b][c] = fa*f(a+1, b, c) + fb*f(a, b+1, c) + fc*f(a, b, c+1)
		memo[a][b][c] /= fs
		memo[a][b][c] += 1.0
		return memo[a][b][c]
	}

	return f(a, b, c)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func Print(x any) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
