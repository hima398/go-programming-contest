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

type PriorityQueue []int

func (pq PriorityQueue) Len() int {
	return len(pq)
}
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i] > pq[j]
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(item interface{}) {
	*pq = append(*pq, item.(int))
}

func (pq *PriorityQueue) Pop() interface{} {
	es := *pq // EdgeのSlice
	n := len(es)
	item := es[n-1]
	*pq = es[0 : n-1]
	return item
}

func firstSolve(t int, n []int, a [][]int) []string {
	var ans []string
	for k := 0; k < t; k++ {
		q := &PriorityQueue{}
		heap.Init(q)
		for _, v := range a[k] {
			heap.Push(q, v)
		}
		for q.Len() >= 3 {
			var b []int
			for j := 0; j < 3; j++ {
				b = append(b, heap.Pop(q).(int))
			}
			sort.Ints(b)
			b[1] -= b[0]
			b[2] -= b[0]
			if b[1] > 0 {
				heap.Push(q, b[1])
			}
			if b[2] > 0 {
				heap.Push(q, b[2])
			}
		}
		if q.Len() == 0 {
			ans = append(ans, "Yes")
		} else {
			ans = append(ans, "No")
		}
	}
	return ans
}

func solve(t int, n []int, a [][]int) []string {
	var ans []string
	for k := 0; k < t; k++ {
		var mx, s int
		for _, v := range a[k] {
			mx = Max(mx, v)
			s += v
		}
		if s%3 == 0 && s/3 >= mx {
			ans = append(ans, "Yes")
		} else {
			ans = append(ans, "No")
		}
	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	t := nextInt()

	var n []int
	var a [][]int
	for i := 0; i < t; i++ {
		ni := nextInt()
		n = append(n, ni)
		a = append(a, nextIntSlice(ni))
	}
	ans := solve(t, n, a)

	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for _, v := range ans {
		fmt.Fprintln(out, v)
	}
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

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
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
