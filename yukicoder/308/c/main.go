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

var p int
var dc []int
var dir = [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}

type Edge struct {
	i, j, c int
}

type PriorityQueue []Edge

func (pq PriorityQueue) Len() int {
	return len(pq)
}
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].c < pq[j].c
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

type Field struct {
	h, w int
	//e [][]Edge
	f []string
	c [][]int
}

func NewField(h, w int, f []string) *Field {
	const INF = 1 << 60
	c := make([][]int, h)
	for i := 0; i < h; i++ {
		c[i] = make([]int, w)
		for j := 0; j < w; j++ {
			c[i][j] = INF
		}
	}
	return &Field{h, w, f, c}
}

func (f *Field) Len() int {
	return f.h * f.w //len(this.E)
}

func (f *Field) Dijkstra(xs, ys, xt, yt int) {
	//n := this.Len()
	q := &PriorityQueue{}

	//TODO
	init := func(xs, ys int) {
		heap.Init(q)
		heap.Push(q, Edge{xs, ys, 0})
	}
	push := func(i, j, c, d int) {
		ni, nj := i+dir[d][0], j+dir[d][1]
		if ni < 0 || ni >= f.h || nj < 0 || nj >= f.w {
			return
		}
		if f.f[ni][nj] == '#' {
			return
		}
		nc := c
		if f.f[ni][nj] == '@' {
			nc += p
		}
		nc += dc[d]

		if f.c[ni][nj] > nc {
			f.c[ni][nj] = nc
			heap.Push(q, Edge{ni, nj, nc})
		}
	}

	init(xs, ys)
	for q.Len() > 0 {
		//fmt.Println(q)
		cur := heap.Pop(q).(Edge)
		//fmt.Println("cur", cur)
		if f.c[cur.i][cur.j] < cur.c {
			continue
		}
		//fmt.Println("Edge", this.E[cur.B])
		for k := 0; k < 4; k++ {
			push(cur.i, cur.j, cur.c, k)
		}
	}
	//f.PrintField()
}

func (f *Field) PrintField() {
	for i := 0; i < f.h; i++ {
		fmt.Println(f.c[i])
	}
}

func solve(h, w, u, d, r, l, k, p, xs, ys, xt, yt int, c []string) string {
	xs--
	ys--
	xt--
	yt--
	dc = append(dc, u)
	dc = append(dc, d)
	dc = append(dc, r)
	dc = append(dc, l)

	field := NewField(h, w, c)
	field.Dijkstra(xs, ys, xt, yt)

	if field.c[xt][yt] <= k {
		return "Yes"
	} else {
		return "No"
	}
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	h, w := nextInt(), nextInt()
	u, d, r, l, k := nextInt(), nextInt(), nextInt(), nextInt(), nextInt()
	p = nextInt()
	xs, ys, xt, yt := nextInt(), nextInt(), nextInt(), nextInt()
	c := make([]string, h)
	for i := 0; i < h; i++ {
		c[i] = nextString()
	}
	ans := solve(h, w, u, d, r, l, k, p, xs, ys, xt, yt, c)
	PrintString(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func PrintString(x string) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
