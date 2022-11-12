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

type PriorityQueue []int

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i] < pq[j]
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(item interface{}) {
	*pq = append(*pq, item.(int))
}

func (pq *PriorityQueue) Pop() interface{} {
	t := *pq
	n := len(t)
	item := t[n-1]
	*pq = t[0 : n-1]
	return item
}

func (pq *PriorityQueue) Top() interface{} {
	t := *pq
	n := len(t)
	item := t[n-1]
	return item
}

func NewPriorityQueue() *PriorityQueue {
	pq := new(PriorityQueue)
	return pq
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	q := nextInt()
	pq := NewPriorityQueue()
	heap.Init(pq)
	for i := 0; i < q; i++ {
		t := nextInt()
		switch t {
		case 1:
			x := nextInt()
			heap.Push(pq, x)
		case 2:
			res := heap.Pop(pq).(int)
			heap.Push(pq, res)
			PrintInt(res)
		case 3:
			heap.Pop(pq)
		}
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
