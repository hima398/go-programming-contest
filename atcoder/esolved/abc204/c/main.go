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

	n, m := nextInt(), nextInt()
	var a, b []int
	for i := 0; i < m; i++ {
		a = append(a, nextInt()-1)
		b = append(b, nextInt()-1)
	}
	ans := solve(n, m, a, b)
	PrintInt(ans)
}

func solve(n, m int, a, b []int) int {
	e := make([][]int, n)
	for i := 0; i < m; i++ {
		e[a[i]] = append(e[a[i]], b[i])
	}
	var visited []bool
	var dfs func(cur int) int
	dfs = func(cur int) int {
		visited[cur] = true
		res := 1
		for _, next := range e[cur] {
			if visited[next] {
				continue
			}
			res += dfs(next)
		}
		return res
	}
	var ans int
	for i := 0; i < n; i++ {
		visited = make([]bool, n)
		ans += dfs(i)
	}
	return ans
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
