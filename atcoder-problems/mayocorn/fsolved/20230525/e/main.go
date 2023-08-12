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

	h, w := nextInt(), nextInt()
	ch, cw := nextInt()-1, nextInt()-1
	dh, dw := nextInt()-1, nextInt()-1
	var s []string
	for i := 0; i < h; i++ {
		s = append(s, nextString())
	}
	ans := solve(h, w, ch, cw, dh, dw, s)
	PrintInt(ans)
}

func solve(h, w, ch, cw, dh, dw int, s []string) int {
	const INF = 1 << 60

	q := &PriorityQueue{}
	heap.Init(q)

	dist := make([][]int, h)
	for i := range dist {
		dist[i] = make([]int, w)
		for j := range dist[i] {
			dist[i][j] = INF
		}
	}
	push := func(ni, nj, cost int) {
		if dist[ni][nj] <= cost {
			return
		}
		dist[ni][nj] = cost
		heap.Push(q, cell{ni, nj, cost})
	}
	push(ch, cw, 0)

	di := []int{-1, 0, 1, 0}
	dj := []int{0, -1, 0, 1}
	for q.Len() > 0 {
		cur := heap.Pop(q).(cell)
		if dist[cur.i][cur.j] < cur.w {
			continue
		}
		for k := 0; k < 4; k++ {
			ni, nj := cur.i+di[k], cur.j+dj[k]
			if ni < 0 || ni >= h || nj < 0 || nj >= w {
				continue
			}
			if s[ni][nj] == '#' {
				continue
			}
			push(ni, nj, cur.w)
		}
		for i := -2; i <= 2; i++ {
			for j := -2; j <= 2; j++ {
				ni, nj := cur.i+i, cur.j+j
				if ni < 0 || ni >= h || nj < 0 || nj >= w {
					continue
				}
				if s[ni][nj] == '#' {
					continue
				}
				push(ni, nj, cur.w+1)
			}
		}
	}
	ans := dist[dh][dw]
	if ans == INF {
		ans = -1
	}
	return ans
}

type cell struct {
	i, j, w int
}

type PriorityQueue []cell

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
	*pq = append(*pq, item.(cell))
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

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
