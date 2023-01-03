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
	var u, v []int
	for i := 0; i < n-1; i++ {
		u = append(u, nextInt()-1)
		v = append(v, nextInt()-1)
	}
	ans := solve(n, u, v)
	PrintVertically(ans)
}

func solve(n int, u, v []int) [][2]int {
	e := make([][]int, n)
	for i := range u {
		e[u[i]] = append(e[u[i]], v[i])
		e[v[i]] = append(e[v[i]], u[i])
	}
	ans := make([][2]int, n)
	var dfs func(i, par, l int) int
	dfs = func(i, par, l int) int {
		ans[i][0] = l
		r := l
		for _, next := range e[i] {
			if next == par {
				continue
			}
			r = dfs(next, i, l)
			l = r + 1
		}
		ans[i][1] = r
		return r
	}
	dfs(0, -1, 1)
	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func PrintVertically(x [][2]int) {
	defer out.Flush()
	for _, v := range x {
		fmt.Fprintln(out, v[0], v[1])
	}
}
