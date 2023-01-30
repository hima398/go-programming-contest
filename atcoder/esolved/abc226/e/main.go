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
	var u, v []int
	for i := 0; i < m; i++ {
		u = append(u, nextInt()-1)
		v = append(v, nextInt()-1)
	}
	ans := solve(n, m, u, v)
	PrintInt(ans)
}

func solve(n, m int, u, v []int) int {
	const p = 998244353
	e := make([][]int, n)
	for i := range u {
		e[u[i]] = append(e[u[i]], v[i])
		e[v[i]] = append(e[v[i]], u[i])
	}
	visited := make([]bool, n)
	//点xを含む連結成分の頂点数と辺の数が一致しているかを探索する
	bfs := func(x int) bool {
		var n, m int
		var q []int
		q = append(q, x)
		visited[x] = true
		n++
		for len(q) > 0 {
			cur := q[0]
			q = q[1:]
			for _, next := range e[cur] {
				m++
				if visited[next] {
					continue
				}
				q = append(q, next)
				visited[next] = true
				n++
			}
		}
		//辺数は2回数えられているので、2*nと比較する
		//fmt.Println(x, n, m)
		return 2*n == m
	}
	ans := 1
	for i := 0; i < n; i++ {
		if visited[i] {
			continue
		}
		if bfs(i) {
			ans = (ans * 2) % p
		} else {
			ans = 0
		}
	}
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
