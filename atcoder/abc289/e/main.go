package main

import (
	"bufio"
	"container/heap"
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

	t := nextInt()
	var ans []int
	for i := 0; i < t; i++ {
		n, m := nextInt(), nextInt()
		c := nextIntSlice(n)
		var u, v []int
		for i := 0; i < m; i++ {
			u = append(u, nextInt()-1)
			v = append(v, nextInt()-1)
		}
		ans = append(ans, solveAfterContest(n, m, c, u, v))
		//ans = append(ans, solve(n, m, c, u, v))
	}
	PrintVertically(ans)
}

func solveAfterContest(n, m int, c, u, v []int) int {
	const INF = 1 << 60
	e := make([][]int, n)
	for i := 0; i < m; i++ {
		e[u[i]] = append(e[u[i]], v[i])
		e[v[i]] = append(e[v[i]], u[i])
	}
	dist := make([][]int, n)
	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			dist[i][j] = INF
		}
	}
	//BFS
	var q []Pos
	//push(0, n-1, 0)
	q = append(q, Pos{0, n - 1, 0})
	dist[0][n-1] = 0
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		for _, nextT := range e[cur.t] {
			for _, nextA := range e[cur.a] {
				if dist[nextT][nextA] < INF {
					continue
				}
				if c[nextT] == c[nextA] {
					continue
				}
				q = append(q, Pos{nextT, nextA, cur.w + 1})
				dist[nextT][nextA] = cur.w + 1
			}
		}
	}
	for i := range dist {
		for j := range dist[i] {
			if dist[i][j] == INF {
				dist[i][j] = -1
			}
		}
	}
	return dist[n-1][0]
}

func solve(n, m int, c, u, v []int) int {
	const INF = 1 << 60
	e := make([][]int, n)
	for i := 0; i < m; i++ {
		e[u[i]] = append(e[u[i]], v[i])
		e[v[i]] = append(e[v[i]], u[i])
	}
	dist := make([][]int, n)
	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			dist[i][j] = INF
		}
	}
	type pos struct {
		t, a int
	}
	//Dijkstra
	q := &PriorityQueue{}
	heap.Init(q)

	push := func(t, a, cost int) {
		if dist[t][a] <= cost {
			return
		}
		if c[t] == c[a] {
			return
		}
		dist[t][a] = cost
		heap.Push(q, Pos{t, a, cost})
	}
	push(0, n-1, 0)
	for q.Len() > 0 {
		cur := heap.Pop(q).(Pos)
		if dist[cur.t][cur.a] < cur.w {
			continue
		}
		for _, nextT := range e[cur.t] {
			for _, nextA := range e[cur.a] {
				push(nextT, nextA, cur.w+1)
			}
		}
	}
	for i := range dist {
		for j := range dist[i] {
			if dist[i][j] == INF {
				dist[i][j] = -1
			}
		}
	}
	return dist[n-1][0]
}

type Pos struct {
	t, a, w int
}

type PriorityQueue []Pos

func (pq PriorityQueue) Len() int {
	return len(pq)
}
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].w < pq[j].w
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(item interface{}) {
	*pq = append(*pq, item.(Pos))
}

func (pq *PriorityQueue) Pop() interface{} {
	es := *pq // EdgeのSlice
	n := len(es)
	item := es[n-1]
	*pq = es[0 : n-1]
	return item
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func nextIntSlice(n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = nextInt()
	}
	return s
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
