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
	x := nextIntSlice(3)
	s := nextString()
	var a, b, c []int
	for i := 0; i < m; i++ {
		a = append(a, nextInt()-1)
		b = append(b, nextInt()-1)
		c = append(c, nextInt())
	}
	ans := solve(n, m, x, s, a, b, c)
	PrintInt(ans)
}

type Edge struct {
	t, w int
}

func solve(n, m int, x []int, s string, a, b, c []int) int {
	const INF = 1 << 60
	e := make([][]Edge, n+6)
	//n  :a->b: xab
	//n+1:a->c: xac
	//n+2:b->c: xbc
	//n+3:b->a: xab
	//n+4:c->a: xac
	//n+5:c->b: xbc
	for i := 0; i < m; i++ {
		e[a[i]] = append(e[a[i]], Edge{b[i], c[i]})
		e[b[i]] = append(e[b[i]], Edge{a[i], c[i]})
	}
	for i, si := range s {
		switch si {
		case 'A':
			e[i] = append(e[i], Edge{n, x[0]})
			e[i] = append(e[i], Edge{n + 1, x[1]})
			e[n+3] = append(e[n+3], Edge{i, 0})
			e[n+4] = append(e[n+4], Edge{i, 0})
		case 'B':
			e[i] = append(e[i], Edge{n + 2, x[2]})
			e[i] = append(e[i], Edge{n + 3, x[0]})
			e[n] = append(e[n], Edge{i, 0})
			e[n+5] = append(e[n+5], Edge{i, 0})
		case 'C':
			e[i] = append(e[i], Edge{n + 4, x[1]})
			e[i] = append(e[i], Edge{n + 5, x[2]})
			e[n+1] = append(e[n+1], Edge{i, 0})
			e[n+2] = append(e[n+2], Edge{i, 0})
		}
	}
	dist := make([]int, n+6)
	for i := 0; i < n+6; i++ {
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

	//fmt.Println(dist)
	return dist[n-1]
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

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}

func PrintString(x string) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
