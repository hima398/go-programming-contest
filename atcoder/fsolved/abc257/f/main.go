package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

const INF = 1 << 60

func solve(n, m int, u, v []int) []int {
	var sv []int
	e := make([][]int, n+1)
	for i := 0; i < m; i++ {
		if u[i] == 0 {
			sv = append(sv, v[i])
		} else {
			e[u[i]] = append(e[u[i]], v[i])
			e[v[i]] = append(e[v[i]], u[i])
		}
	}
	d := make([][]int, 2)
	for i := 0; i < 2; i++ {
		d[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			d[i][j] = INF
		}
	}

	bfs := func(root, t int) {
		var q []int
		q = append(q, root)
		d[t][root] = 0
		for len(q) > 0 {
			cur := q[0]
			q = q[1:]
			for _, next := range e[cur] {
				if d[t][next] < INF {
					continue
				}
				q = append(q, next)
				d[t][next] = d[t][cur] + 1
			}
		}
	}
	bfs(1, 0)
	bfs(n, 1)

	//i1, in := 0, 0
	mn1, mnn := INF, INF
	for _, v := range sv {
		if mn1 > d[0][v] {
			mn1 = d[0][v]
			//i1 = v
		}
		if mnn > d[1][v] {
			mnn = d[1][v]
			//in = v
		}
	}
	var ans []int
	for i := 1; i <= n; i++ {
		v := d[0][n]
		v = Min(v, mn1+1+d[1][i])
		v = Min(v, mnn+1+d[0][i])
		v = Min(v, mn1+mnn+2)
		ans = append(ans, v)
	}
	for i := range ans {
		if ans[i] >= INF {
			ans[i] = -1
		}
	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, m := nextInt(), nextInt()
	var u, v []int
	for i := 0; i < m; i++ {
		u = append(u, nextInt())
		v = append(v, nextInt())
	}
	ans := solve(n, m, u, v)
	PrintHorizonaly(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func PrintHorizonaly(x []int) {
	defer out.Flush()
	fmt.Fprintf(out, "%d", x[0])
	for i := 1; i < len(x); i++ {
		fmt.Fprintf(out, " %d", x[i])
	}
	fmt.Fprintln(out)
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
