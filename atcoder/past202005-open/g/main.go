package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

type point struct {
	x, y, d int
}

type PriorityQueue []point

func (pq PriorityQueue) Len() int {
	return len(pq)
}
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].d > pq[j].d
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(item interface{}) {
	*pq = append(*pq, item.(point))
}

func (pq *PriorityQueue) Pop() interface{} {
	es := *pq // EdgeのSlice
	n := len(es)
	item := es[n-1]
	*pq = es[0 : n-1]
	return item
}

func solve(n, tx, ty int, x, y []int) int {
	const INF = 1 << 60
	const mx = 200
	dy := []int{1, 1, 1, 0, 0, -1}
	dx := []int{1, -1, 0, 1, -1, 0}
	h := 2*mx + 1
	w := h
	f := make([][]int, h)
	v := make([][]bool, h)
	//fmt.Println("h = ", len(f))
	for i := 0; i < h; i++ {
		f[i] = make([]int, w)
		v[i] = make([]bool, w)
		for j := 0; j < w; j++ {
			f[i][j] = INF
		}
		//fmt.Println("w = ", len(f[i]))
	}
	for k := 0; k < n; k++ {
		v[y[k]+200][x[k]+200] = true
	}
	//var q []point
	q := &PriorityQueue{}
	heap.Init(q)
	sx, sy := 200, 200
	//q = append(q, point{sx, sy, 0})
	q.Push()
	f[sy][sx] = 0
	v[sy][sx] = true
	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		for k := 0; k < 6; k++ {
			nx, ny := p.x+dx[k], p.y+dy[k]
			if nx < 0 || nx >= w || ny < 0 || ny >= h {
				continue
			}
			if v[ny][nx] {
				continue
			}
			if nx == tx+200 && ny == ty+200 {
				return p.d + 1
			}
			f[ny][nx] = p.d + 1
			q = append(q, point{nx, ny, p.d + 1})
			v[ny][nx] = true
		}
	}
	return -1
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, tx, ty := nextInt(), nextInt(), nextInt()
	var x, y []int
	for i := 0; i < n; i++ {
		x = append(x, nextInt())
		y = append(y, nextInt())
	}

	ans := solve(n, tx, ty, x, y)

	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
