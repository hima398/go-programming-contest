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
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, m := nextInt(), nextInt()
	var t, x []int
	for i := 0; i < n; i++ {
		t = append(t, nextInt())
		x = append(x, nextInt())
	}
	ans := solve(n, m, t, x)
	PrintInt(ans)
}

func solve(n, m int, t, x []int) int {
	q := &PriorityQueue{}
	heap.Init(q)
	var cs []int
	var openers []int
	var point int
	for i := 0; i < n; i++ {
		switch t[i] {
		case 0:
			heap.Push(q, x[i])
			point += x[i]
			if q.Len() > m {
				v := heap.Pop(q).(int)
				point -= v
			}
		case 1:
			cs = append(cs, x[i])
		case 2:
			openers = append(openers, x[i])
		}
	}
	sort.Ints(cs)
	sort.Ints(openers)
	var numOpen int
	ans := point
	for len(cs) > 0 {
		//fmt.Println(point, q, cs, openers)
		if numOpen == 0 {
			if len(openers) == 0 {
				break
			}
			numOpen += openers[len(openers)-1]
			openers = openers[:len(openers)-1]
			m--
			for q.Len() > m {
				v := heap.Pop(q).(int)
				point -= v
			}
		} else {
			next := cs[len(cs)-1]
			cs = cs[:len(cs)-1]
			heap.Push(q, next)
			point += next
			for q.Len() > m {
				v := heap.Pop(q).(int)
				point -= v
			}
			ans = Max(ans, point)
			numOpen--
		}
	}
	return ans
}

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

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
