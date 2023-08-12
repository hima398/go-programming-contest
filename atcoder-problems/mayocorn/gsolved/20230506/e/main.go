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

	n, m := nextInt(), nextInt()
	var a, b []int
	for i := 0; i < n; i++ {
		a = append(a, nextInt()-1)
		b = append(b, nextInt())
	}
	ans := solve(n, m, a, b)
	PrintInt(ans)
}

func solve(n, m int, a, b []int) int {
	//k日後(0-indexed)に得られる報酬の集合
	tasks := make(map[int][]int)
	for i := 0; i < n; i++ {
		tasks[a[i]] = append(tasks[a[i]], b[i])
	}

	q := &PriorityQueue{}
	heap.Init(q)

	var ans int
	//残りi日
	for i := 0; i < m; i++ {
		for _, b := range tasks[i] {
			heap.Push(q, b)
		}
		if q.Len() > 0 {
			v := heap.Pop(q).(int)
			ans += v
		}
	}
	return ans
}

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
	es := *pq
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
