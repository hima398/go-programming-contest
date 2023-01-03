package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	var a, b []int
	for i := 0; i < n; i++ {
		a = append(a, nextInt())
		b = append(b, nextInt())
	}
	ans := solve(n, a, b)
	PrintVertically(ans)
}

func solve(n int, a, b []int) []int {
	var ts []Task
	for i := 0; i < n; i++ {
		ts = append(ts, Task{a[i], b[i]})
	}
	sort.Slice(ts, func(i, j int) bool {
		return ts[i].day < ts[j].day
	})
	q := &PriorityQueue{}
	heap.Init(q)
	s := 0
	var ans []int
	for i := 1; i <= n; i++ {
		for len(ts) > 0 && ts[0].day == i {
			heap.Push(q, ts[0])
			ts = ts[1:]
		}
		if q.Len() > 0 {
			node := heap.Pop(q).(Task)
			s += node.point
		}
		ans = append(ans, s)
	}
	return ans
}

type Task struct {
	day, point int
}

type PriorityQueue []Task

func (pq PriorityQueue) Len() int {
	return len(pq)
}
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].point > pq[j].point
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(item interface{}) {
	*pq = append(*pq, item.(Task))
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
