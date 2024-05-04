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

	n := nextInt()
	var a, b, x []int
	for i := 0; i < n-1; i++ {
		a = append(a, nextInt())
		b = append(b, nextInt())
		x = append(x, nextInt())
	}
	ans := solve(n, a, b, x)
	Print(ans)
}

func solve(n int, a, b, x []int) int {
	const INF = 1 << 60
	e := make([][]Edge, n)
	for i := 0; i < n-1; i++ {
		e[i] = append(e[i], Edge{i + 1, a[i]})
		e[i] = append(e[i], Edge{x[i] - 1, b[i]})
	}
	dist := make([]int, n)
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
	return dist[n-1]

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
	return int(i)
}

func Print(x any) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
