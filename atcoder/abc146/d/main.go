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
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	var a, b []int
	for i := 0; i < n-1; i++ {
		a = append(a, nextInt()-1)
		b = append(b, nextInt()-1)
	}
	k, ans := solve(n, a, b)
	PrintInt(k)
	PrintVertically(ans)
}
func solve(n int, a, b []int) (int, []int) {
	type edge struct {
		i, s, t int
	}
	m := n - 1
	e := make([][]edge, n)
	for i := 0; i < m; i++ {
		e[a[i]] = append(e[a[i]], edge{i, a[i], b[i]})
		e[b[i]] = append(e[b[i]], edge{i, b[i], a[i]})
	}
	k := 0
	for i := 0; i < n; i++ {
		k = Max(k, len(e[i]))
	}
	ans := make([]int, m)
	var dfs func(cur, par, parColor int)
	dfs = func(cur, par, parColor int) {
		nextColor := 1
		for _, next := range e[cur] {
			if next.t == par {
				continue
			}
			if nextColor == parColor {
				nextColor++
			}
			ans[next.i] = nextColor
			dfs(next.t, cur, nextColor)
			nextColor++
		}
	}

	dfs(0, -1, 0)

	return k, ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}

func PrintVertically(x []int) {
	defer out.Flush()
	for _, v := range x {
		fmt.Fprintln(out, v)
	}
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
