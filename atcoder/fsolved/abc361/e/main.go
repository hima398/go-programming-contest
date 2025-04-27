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
	var a, b, c []int
	for i := 0; i < n-1; i++ {
		a = append(a, nextInt()-1)
		b = append(b, nextInt()-1)
		c = append(c, nextInt())
	}
	ans := solve(n, a, b, c)
	Print(ans)
}

func solve(n int, a, b, c []int) int {
	type edge struct {
		to, cost int
	}
	e := make([][]edge, n)
	for i := 0; i < n-1; i++ {
		e[a[i]] = append(e[a[i]], edge{b[i], c[i]})
		e[b[i]] = append(e[b[i]], edge{a[i], c[i]})
	}

	dist := make([]int, n)
	var dfs func(cur, par int)
	dfs = func(cur, par int) {
		for _, next := range e[cur] {
			if next.to == par {
				continue
			}
			dist[next.to] = dist[cur] + next.cost
			dfs(next.to, cur)
		}
	}
	dfs(0, -1)
	var root, max int
	for i, v := range dist {
		if v > max {
			root = i
			max = v
		}
	}
	for i := range dist {
		dist[i] = 0
	}
	dfs(root, -1)
	var s int
	for i := 0; i < n-1; i++ {
		s += 2 * c[i]
	}
	var d int
	for _, v := range dist {
		d = Max(d, v)
	}

	return s - d
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

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
