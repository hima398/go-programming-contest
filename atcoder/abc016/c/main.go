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
	PrintVertically(ans)
}

func solve(n, m int, a, b []int) []int {
	ans := make([]int, n)
	e := make([][]int, n)
	for i := 0; i < m; i++ {
		e[a[i]] = append(e[a[i]], b[i])
		e[b[i]] = append(e[b[i]], a[i])
	}
	bfs := func(x int) int {
		d := make([]int, n)
		for i := range d {
			d[i] = -1
		}
		var q []int
		q = append(q, x)
		d[x] = 0
		for len(q) > 0 {
			p := q[0]
			q = q[1:]
			for _, next := range e[p] {
				if d[next] >= 0 {
					continue
				}
				q = append(q, next)
				d[next] = d[p] + 1
			}
		}
		res := 0
		for i := 0; i < n; i++ {
			if d[i] == 2 {
				res++
			}
		}
		return res
	}
	for i := 0; i < n; i++ {
		ans[i] = bfs(i)
	}
	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func PrintVertically(x []int) {
	defer out.Flush()
	for _, v := range x {
		fmt.Fprintln(out, v)
	}
}
