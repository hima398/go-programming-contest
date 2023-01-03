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
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, m := nextInt(), nextInt()
	var k []int
	r := make([][]int, m)
	for i := 0; i < m; i++ {
		k = append(k, nextInt())
		r[i] = nextIntSlice(k[i])
		//0-indexed
		for j := range r[i] {
			r[i][j]--
		}
	}
	ans := solve(n, m, k, r)
	PrintVertically(ans)
}

func solve(n, m int, k []int, r [][]int) []int {
	const INF = 1 << 60

	e := make([][]Edge, n+m)
	for i := 0; i < m; i++ {
		for _, v := range r[i] {
			e[n+i] = append(e[n+i], Edge{v, 1})
			e[v] = append(e[v], Edge{n + i, 1})
		}
	}

	dist := make([]int, n+m)
	for i := range dist {
		dist[i] = INF
	}

	//Dijkstra
	q := &PriorityQueue{}
	heap.Init(q)

	push := func(to, cost int) {
		if dist[to] <= cost {
			return
		}
		dist[to] = cost
		heap.Push(q, Edge{to, cost})
	}
	push(0, 0)
	for q.Len() > 0 {
		cur := heap.Pop(q).(Edge)
		if dist[cur.t] < cur.w {
			continue
		}
		for _, next := range e[cur.t] {
			push(next.t, cur.w+next.w)
		}
	}
	for i := range dist {
		if dist[i] == INF {
			dist[i] = -1
		}
	}
	dist = dist[:n]
	for i := range dist {
		if dist[i] >= 0 {
			dist[i] /= 2
		}
	}
	return dist
}

type Edge struct {
	t, w int
}

type PriorityQueue []Edge

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
	*pq = append(*pq, item.(Edge))
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
	return i
}

func nextIntSlice(n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = nextInt()
	}
	return s
}

func PrintVertically(x []int) {
	defer out.Flush()
	for _, v := range x {
		fmt.Fprintln(out, v)
	}
}
