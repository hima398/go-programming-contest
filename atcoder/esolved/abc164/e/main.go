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

	n, m, s := nextInt(), nextInt(), nextInt()
	var u, v, a, b []int
	for i := 0; i < m; i++ {
		u = append(u, nextInt()-1)
		v = append(v, nextInt()-1)
		a = append(a, nextInt())
		b = append(b, nextInt())
	}
	var c, d []int
	for i := 0; i < n; i++ {
		c = append(c, nextInt())
		d = append(d, nextInt())
	}
	ans := solve(n, m, s, u, v, a, b, c, d)
	PrintVertically(ans)
}

func solve(n, m, s int, u, v, a, b, c, d []int) []int {
	const INF = 1 << 60
	const maxA = 2501

	convertToIndex := func(i, j int) int {
		return i*maxA + j
	}

	l := (n + 1) * maxA
	e := make([][]edge, l)
	for i := 0; i < n; i++ {
		for j := 0; j < maxA-1; j++ {
			nj := Min(j+c[i], maxA-1)
			e[convertToIndex(i, j)] = append(e[convertToIndex(i, j)], edge{convertToIndex(i, j), convertToIndex(i, nj), c[i], d[i]})
		}
	}
	for i := 0; i < m; i++ {
		for j := maxA - 1; j >= 0; j-- {
			if j-a[i] < 0 {
				break
			}
			nj := j - a[i]
			e[convertToIndex(u[i], j)] = append(e[convertToIndex(u[i], j)], edge{convertToIndex(u[i], j), convertToIndex(v[i], nj), a[i], b[i]})
			e[convertToIndex(v[i], j)] = append(e[convertToIndex(v[i], j)], edge{convertToIndex(v[i], j), convertToIndex(u[i], nj), a[i], b[i]})
		}
	}

	q := &PriorityQueue{}
	heap.Init(q)

	times := make([]int, l)
	for i := 0; i < l; i++ {
		times[i] = INF
	}
	push := func(to, s, cost int) {
		if times[to] <= cost {
			return
		}
		times[to] = cost
		heap.Push(q, node{to, s, cost})
	}
	push(Min(s, maxA-1), s, 0)

	for q.Len() > 0 {
		cur := heap.Pop(q).(node)
		if times[cur.to] < cur.time {
			continue
		}
		for _, next := range e[cur.to] {
			_, j := next.to/maxA, next.to%maxA

			if j >= next.fare {
				//コストを払って次の頂点にいく
				//nj := j - next.fare
				ns := cur.s - next.fare
				push(next.to, ns, cur.time+next.cost)
			} else {
				push(next.to, cur.s+next.fare, cur.time+next.cost)
			}
		}
	}
	//fmt.Println(times)

	ans := make([]int, n-1)
	for i := 1; i < n; i++ {
		ans[i-1] = INF
		for j := 0; j < maxA; j++ {
			ans[i-1] = Min(ans[i-1], times[convertToIndex(i, j)])
		}
	}
	return ans

}

type edge struct {
	from, to, fare, cost int
}

type node struct {
	to, s, time int
}
type PriorityQueue []node

func (q PriorityQueue) Len() int {
	return len(q)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].time < pq[j].time
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(item interface{}) {
	*pq = append(*pq, item.(node))
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

func PrintVertically(x []int) {
	defer out.Flush()
	for _, v := range x {
		fmt.Fprintln(out, v)
	}
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
