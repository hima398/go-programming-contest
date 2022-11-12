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
	var a, b, c []int
	for i := 0; i < m; i++ {
		a = append(a, nextInt()-1)
		b = append(b, nextInt()-1)
		c = append(c, nextInt())
	}
	ans := solve(n, m, a, b, c)
	PrintVertically(ans)
}

func solve(n, m int, a, b, c []int) []int {
	const INF = 1 << 60
	dist := make([]int, n)
	for i := range dist {
		dist[i] = INF
	}

	e := make([][]Edge, n)
	for i := 0; i < m; i++ {
		e[a[i]] = append(e[a[i]], Edge{b[i], c[i]})
		e[b[i]] = append(e[b[i]], Edge{a[i], c[i]})
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

func PrintVertically(x []int) {
	defer out.Flush()
	for _, v := range x {
		fmt.Fprintln(out, v)
	}
}
