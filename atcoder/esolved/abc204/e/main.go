package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
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
	var a, b, c, d []int
	for i := 0; i < m; i++ {
		a = append(a, nextInt()-1)
		b = append(b, nextInt()-1)
		c = append(c, nextInt())
		d = append(d, nextInt())
	}
	ans := solve(n, m, a, b, c, d)
	PrintInt(ans)
}

func solve(n, m int, a, b, c, d []int) int {
	const INF = 1 << 60
	dist := make([]int, n)
	for i := range dist {
		dist[i] = INF
	}

	e := make([][]Edge, n)
	for i := 0; i < m; i++ {
		if a[i] == b[i] {
			continue
		}
		e[a[i]] = append(e[a[i]], Edge{b[i], c[i], d[i]})
		e[b[i]] = append(e[b[i]], Edge{a[i], c[i], d[i]})
	}

	//Dijkstra
	q := &PriorityQueue{}
	heap.Init(q)

	computeCost := func(t, c, d int) int {
		f := func(x, c, d int) int {
			if x < 0 {
				x = 0
			}
			//fmt.Println(x, c, d)
			return x + c + Floor(d, x+1)
		}
		t1, t2 := Sqrt(d)-1, Sqrt(d)
		if t > t2 {
			return f(t, c, d)
		} else {
			return Min(f(t1, c, d), f(t2, c, d))
		}
	}

	push := func(to, cost int) {
		if dist[to] <= cost {
			return
		}
		dist[to] = cost
		heap.Push(q, Node{to, cost})
	}
	push(0, 0)
	for q.Len() > 0 {
		//fmt.Println(q)
		cur := heap.Pop(q).(Node)
		if dist[cur.t] < cur.c {
			continue
		}
		for _, next := range e[cur.t] {
			cost := computeCost(cur.c, next.c, next.d)
			push(next.t, cost)
		}
	}
	for i := range dist {
		if dist[i] == INF {
			dist[i] = -1
		}
	}
	//fmt.Println(dist)
	return dist[n-1]
}

type Edge struct {
	t, c, d int
}

type Node struct {
	t, c int
}

type PriorityQueue []Node

func (pq PriorityQueue) Len() int {
	return len(pq)
}
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].c < pq[j].c
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(item interface{}) {
	*pq = append(*pq, item.(Node))
}

func (pq *PriorityQueue) Pop() interface{} {
	es := *pq // NodeのSlice
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

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Floor(x, y int) int {
	return x / y
}

func Sqrt(x int) int {
	x2 := int(math.Sqrt(float64(x))) - 1
	for (x2+1)*(x2+1) <= x {
		x2++
	}
	return x2
}
