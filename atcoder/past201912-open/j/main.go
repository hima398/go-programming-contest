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

	h, w := nextInt(), nextInt()
	a := make([][]int, h)
	for i := 0; i < h; i++ {
		a[i] = nextIntSlice(w)
	}
	ans := solve(h, w, a)
	PrintInt(ans)
}

func solve(h, w int, a [][]int) int {
	field := new(Field)
	field.h = h
	field.w = w
	field.f = a
	ans := 50 * 50 * int(1e5)
	from := field.Dijkstra(h-1, 0)
	via := field.Dijkstra(h-1, w-1)
	to := field.Dijkstra(0, w-1)
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			s := from[i][j] + via[i][j] + to[i][j]
			s -= 2 * field.f[i][j]
			//if ans > s {
			//fmt.Println("i, j = ", i, j, s)
			//ans = s
			//}
			ans = Min(ans, s)
		}
	}
	return ans
}

type Field struct {
	h, w int
	f    [][]int
}

func (f *Field) Dijkstra(si, sj int) [][]int {
	const INF = 1 << 60

	dist := make([][]int, f.h)
	for i := 0; i < f.h; i++ {
		dist[i] = make([]int, f.w)
		for j := 0; j < f.w; j++ {
			dist[i][j] = INF
		}
	}

	q := &PriorityQueue{}
	heap.Init(q)

	push := func(i, j, w int) {
		if dist[i][j] <= w {
			return
		}
		dist[i][j] = w
		heap.Push(q, Cell{i, j, w})
	}
	push(si, sj, f.f[si][sj])
	di := []int{-1, 0, 1, 0}
	dj := []int{0, -1, 0, 1}
	for q.Len() > 0 {
		cur := heap.Pop(q).(Cell)
		if dist[cur.i][cur.j] < cur.w {
			continue
		}
		for k := 0; k < 4; k++ {
			ni, nj := cur.i+di[k], cur.j+dj[k]
			if ni < 0 || ni >= f.h || nj < 0 || nj >= f.w {
				continue
			}
			push(ni, nj, cur.w+f.f[ni][nj])
		}
	}

	return dist
}

type Cell struct {
	i, j, w int
}

type PriorityQueue []Cell

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
	*pq = append(*pq, item.(Cell))
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
