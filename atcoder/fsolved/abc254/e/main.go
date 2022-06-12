package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

var e [][]int

var s [][4]int

type item struct {
	i, d int
}

func bfs(i, ki int) {
	var visited [15*int(1e4) + 1]bool

	root := i
	var q []item
	q = append(q, item{i, 0})

	visited[i] = true
	for len(q) > 0 {
		it := q[0]
		q = q[1:]

		s[root][it.d] += it.i
		if it.d >= ki {
			continue
		}
		for _, next := range e[it.i] {

			if visited[next] {
				continue
			}
			q = append(q, item{next, it.d + 1})

			visited[next] = true
		}
	}
}

func solve(n, m, q int, a, b, x, k []int) []int {

	e = make([][]int, n+1)
	for i := 0; i < m; i++ {
		e[a[i]] = append(e[a[i]], b[i])
		e[b[i]] = append(e[b[i]], a[i])
	}
	s = make([][4]int, n+1)
	mq := make(map[int]int)
	for i, v := range x {
		mq[v] = Max(mq[v], k[i])
	}

	for k, v := range mq {
		bfs(k, v)
	}

	for i := 0; i <= n; i++ {
		for j := 1; j <= 3; j++ {
			s[i][j] += s[i][j-1]
		}
	}

	var ans []int
	for i := 0; i < q; i++ {
		ans = append(ans, s[x[i]][k[i]])
	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, m := nextInt(), nextInt()
	var a, b []int
	for i := 0; i < m; i++ {
		ai, bi := nextInt(), nextInt()
		a = append(a, ai)
		b = append(b, bi)
	}
	q := nextInt()
	var x, k []int
	for i := 0; i < q; i++ {
		xi, ki := nextInt(), nextInt()
		x = append(x, xi)
		k = append(k, ki)
	}
	ans := solve(n, m, q, a, b, x, k)

	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for _, v := range ans {
		fmt.Fprintln(out, v)
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
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
