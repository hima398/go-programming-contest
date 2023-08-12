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
	var s, t []int
	for i := 0; i < m; i++ {
		s = append(s, nextInt()-1)
		t = append(t, nextInt()-1)
	}
	//ans := solveHonestly(n, m, s, t)
	ans := solve(n, m, s, t)
	PrintVertically(ans)
}

func solveHonestly(n, m int, s, t []int) []int {
	//愚直にやるとO(M(N*M))程度で間に合わない
	const INF = 1 << 60
	var ans []int
	bfs := func(e [][]int) int {
		d := make([]int, n)
		for i := range d {
			d[i] = INF
		}
		var q []int
		q = append(q, 0)
		d[0] = 0
		for len(q) > 0 {
			cur := q[0]
			q = q[1:]
			for _, next := range e[cur] {
				if d[next] < INF {
					continue
				}
				q = append(q, next)
				d[next] = d[cur] + 1
			}
		}
		if d[n-1] == INF {
			return -1
		} else {
			return d[n-1]
		}
	}
	for q := 0; q < m; q++ {
		e := make([][]int, n)
		for i := 0; i < m; i++ {
			if i == q {
				continue
			}
			e[s[i]] = append(e[s[i]], t[i])
		}
		ans = append(ans, bfs(e))
	}
	return ans
}

func solve(n, m int, s, t []int) []int {
	const INF = 1 << 60
	type edge struct {
		i, t int
	}
	ans := make([]int, m)
	d1 := make([]int, n)
	for i := range d1 {
		d1[i] = INF
	}
	e1 := make([][]edge, n)
	for i := 0; i < m; i++ {
		e1[s[i]] = append(e1[s[i]], edge{i, t[i]})
	}
	type node struct {
		i     int
		route []int
	}
	bfs1 := func() (int, []int) {
		d := make([]int, n)
		for i := range d {
			d[i] = INF
		}
		var q []node
		q = append(q, node{0, []int{}})
		d[0] = 0
		for len(q) > 0 {
			cur := q[0]
			q = q[1:]
			if cur.i == n-1 {
				return d[n-1], cur.route
			}
			for _, next := range e1[cur.i] {
				if d[next.t] < INF {
					continue
				}
				q = append(q, node{next.t, append(cur.route, next.i)})
				d[next.t] = d[cur.i] + 1
			}
		}
		return -1, []int{}
	}
	dist, route := bfs1()

	for i := range ans {
		ans[i] = dist
	}
	//すべての辺を使ってもゴールにたどり着けない場合はそこから1本辺を取っても
	//たどり着けない。ので、ここで解答を返しておく
	if dist < 0 {
		return ans
	}

	bfs2 := func(e [][]int) int {
		d := make([]int, n)
		for i := range d {
			d[i] = INF
		}
		var q []int
		q = append(q, 0)
		d[0] = 0
		for len(q) > 0 {
			cur := q[0]
			q = q[1:]
			for _, next := range e[cur] {
				if d[next] < INF {
					continue
				}
				q = append(q, next)
				d[next] = d[cur] + 1
			}
		}
		if d[n-1] == INF {
			return -1
		} else {
			return d[n-1]
		}
	}
	for _, q := range route {
		e := make([][]int, n)
		for i := 0; i < m; i++ {
			if i == q {
				continue
			}
			e[s[i]] = append(e[s[i]], t[i])
		}
		ans[q] = bfs2(e)
	}
	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func PrintVertically(x []int) {
	defer out.Flush()
	for _, v := range x {
		fmt.Fprintln(out, v)
	}
}
