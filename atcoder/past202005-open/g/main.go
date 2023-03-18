package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

type cell struct {
	x, y int
}

func solve(n, tx, ty int, x, y []int) int {
	const offset = 220
	h := 2*offset + 1
	w := h
	d := make([][]int, h)
	v := make([][]bool, h)
	for i := 0; i < h; i++ {
		d[i] = make([]int, w)
		v[i] = make([]bool, w)
	}
	//壁を訪問済みとして登録する
	for i := 0; i < n; i++ {
		v[x[i]+offset][y[i]+offset] = true
	}

	dx := []int{1, 0, -1, 1, -1, 0}
	dy := []int{1, 1, 1, 0, 0, -1}

	var q []cell
	q = append(q, cell{offset, offset})
	v[offset][offset] = true
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		for k := 0; k < 6; k++ {
			nx, ny := cur.x+dx[k], cur.y+dy[k]
			if nx < 0 || nx >= h || ny < 0 || ny >= w {
				continue
			}
			if v[nx][ny] {
				continue
			}
			q = append(q, cell{nx, ny})
			v[nx][ny] = true
			d[nx][ny] = d[cur.x][cur.y] + 1
		}
	}
	if v[tx+offset][ty+offset] {
		return d[tx+offset][ty+offset]
	} else {
		return -1
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
