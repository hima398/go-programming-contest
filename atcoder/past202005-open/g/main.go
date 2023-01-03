package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

type cell struct {
	x, y, d int
}

func solve(n, tx, ty int, x, y []int) int {
	const INF = 1 << 60
	const mx = 200
	dy := []int{1, 1, 1, 0, 0, -1}
	dx := []int{1, 0, -1, 1, -1, 0}
	h := 2*mx + 1
	w := h
	dist := make([][]int, h)
	v := make([][]bool, h)
	for i := 0; i < h; i++ {
		dist[i] = make([]int, w)
		v[i] = make([]bool, w)
	}
	//fmt.Println("h = ", len(f))

	for k := 0; k < n; k++ {
		v[y[k]+200][x[k]+200] = true
	}
	var q []cell
	q = append(q, cell{200, 200, 0})
	v[200][200] = true
	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		for k := 0; k < 6; k++ {
			ni, nj, nd := p.y+dy[k], p.x+dx[k], p.d+1
			if ni < 0 || ni >= h || nj < 0 || nj >= w {
				continue
			}
			if v[ni][nj] {
				continue
			}
			dist[ni][nj] = nd
			q = append(q, cell{ni, nj, nd})
			v[ni][nj] = true
		}
	}
	if !v[ty+200][tx+200] {
		return -1
	} else {
		return dist[ty+200][tx+200]
	}
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
