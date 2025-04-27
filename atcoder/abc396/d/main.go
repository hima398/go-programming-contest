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

	n, m := nextInt(), nextInt()
	var u, v, w []int
	for i := 0; i < m; i++ {
		u = append(u, nextInt()-1)
		v = append(v, nextInt()-1)
		w = append(w, nextInt())
	}

	ans := solve(n, m, u, v, w)

	Print(ans)
}

func solve(n, m int, u, v, w []int) int {
	type edge struct {
		to, label int
	}
	e := make([][]edge, n)
	for i := 0; i < m; i++ {
		e[u[i]] = append(e[u[i]], edge{v[i], w[i]})
		e[v[i]] = append(e[v[i]], edge{u[i], w[i]})
	}
	visited := make([]bool, n)
	ans := 1 << 60
	var dfs func(cur, x int)
	dfs = func(cur, x int) {
		if cur == n-1 {
			ans = Min(ans, x)
			return
		}
		visited[cur] = true
		for _, next := range e[cur] {
			if visited[next.to] {
				continue
			}
			dfs(next.to, x^next.label)
		}
		visited[cur] = false
	}
	dfs(0, 0)
	return ans
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

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
