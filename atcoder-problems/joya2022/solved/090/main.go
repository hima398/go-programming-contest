package main

import (
	"bufio"
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
	c := make([]string, h)
	for i := 0; i < h; i++ {
		c[i] = nextString()
	}
	ans := solve(h, w, c)
	PrintString(ans)
}

func solve(h, w int, c []string) string {
	di := []int{-1, 0, 1, 0}
	dj := []int{0, -1, 0, 1}
	oi, oj := -1, -1
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if c[i][j] == 'S' {
				oi, oj = i, j
			}
		}
	}
	type cell struct {
		i, j int
	}
	isInField := func(i, j int) bool {
		if i < 0 || i >= h || j < 0 || j >= w {
			return false
		}
		return true
	}
	for dir := 0; dir < 4; dir++ {
		visited := make([][]bool, h)
		for i := 0; i < h; i++ {
			visited[i] = make([]bool, w)
		}
		si, sj := oi+di[dir], oj+dj[dir]
		if !isInField(si, sj) {
			continue
		}
		if c[si][sj] == '#' {
			continue
		}
		var q []cell
		q = append(q, cell{si, sj})
		visited[si][sj] = true
		for len(q) > 0 {
			cur := q[0]
			q = q[1:]
			for k := 0; k < 4; k++ {
				if k == dir {
					continue
				}
				gi, gj := oi+di[k], oj+dj[k]
				if cur.i == gi && cur.j == gj {
					return "Yes"
				}
			}
			for k := 0; k < 4; k++ {
				ni, nj := cur.i+di[k], cur.j+dj[k]
				if !isInField(ni, nj) {
					continue
				}
				if c[ni][nj] == 'S' || c[ni][nj] == '#' || visited[ni][nj] {
					continue
				}
				q = append(q, cell{ni, nj})
				visited[ni][nj] = true
			}
		}
	}
	return "No"
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
