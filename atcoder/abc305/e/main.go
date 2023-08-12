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

	n, m, k := nextInt(), nextInt(), nextInt()
	var a, b []int
	for i := 0; i < m; i++ {
		a = append(a, nextInt()-1)
		b = append(b, nextInt()-1)
	}
	var p, h []int
	for i := 0; i < k; i++ {
		p = append(p, nextInt()-1)
		h = append(h, nextInt())
	}
	ans := solve(n, m, k, a, b, p, h)
	PrintInt(len(ans))
	PrintHorizonaly(ans)
}

func solve(n, m, k int, a, b, p, h []int) []int {
	e := make([][]int, n)
	for i := 0; i < m; i++ {
		e[a[i]] = append(e[a[i]], b[i])
		e[b[i]] = append(e[b[i]], a[i])
	}
	//点iを守る警備員の最大体力
	maxStaminas := make([]int, n)
	for i := range maxStaminas {
		maxStaminas[i] = -1
	}
	q := &PriorityQueue{}
	heap.Init(q)

	push := func(to, cost int) {
		if maxStaminas[to] >= cost {
			return
		}
		maxStaminas[to] = cost
		heap.Push(q, Edge{to, cost})
	}
	for i := 0; i < k; i++ {
		push(p[i], h[i])
	}
	for q.Len() > 0 {
		cur := heap.Pop(q).(Edge)
		if maxStaminas[cur.t] > cur.w {
			continue
		}
		for _, next := range e[cur.t] {
			push(next, maxStaminas[cur.t]-1)
		}
	}
	var ans []int
	for i, v := range maxStaminas {
		if v >= 0 {
			ans = append(ans, i+1)
		}
	}
	return ans
}

type Edge struct {
	t, w int
}

type PriorityQueue []Edge

func (pq PriorityQueue) Len() int {
	return len(pq)
}
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].w > pq[j].w
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

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}

func PrintHorizonaly(x []int) {
	defer out.Flush()
	fmt.Fprintf(out, "%d", x[0])
	for i := 1; i < len(x); i++ {
		fmt.Fprintf(out, " %d", x[i])
	}
	fmt.Fprintln(out)
}
