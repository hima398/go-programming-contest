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

	n := nextInt()
	var x, y []int
	for i := 0; i < n; i++ {
		x = append(x, nextInt())
		y = append(y, nextInt())
	}
	ans := solve(n, x, y)
	PrintInt(ans)
}

func solve(n int, x, y []int) int {
	visited := make([]bool, n)
	type cell struct {
		i, x, y int
	}
	var c []cell
	f := make(map[int]map[int]cell)
	for i := 0; i < n; i++ {
		c = append(c, cell{i, x[i], y[i]})
		if f[x[i]] == nil {
			f[x[i]] = make(map[int]cell)
		}
		f[x[i]][y[i]] = cell{i, x[i], y[i]}
	}
	dx := [6]int{-1, -1, 0, 0, 1, 1}
	dy := [6]int{-1, 0, -1, 1, 0, 1}
	bfs := func(idx int) {
		var q []int
		q = append(q, idx)
		visited[idx] = true
		for len(q) > 0 {
			cur := q[0]
			q = q[1:]
			for k := 0; k < 6; k++ {
				nx, ny := c[cur].x+dx[k], c[cur].y+dy[k]
				if _, found := f[nx][ny]; found {
					ni := f[nx][ny].i
					if visited[ni] {
						continue
					}
					q = append(q, ni)
					visited[ni] = true
				}
			}
		}
	}
	var ans int
	for i := 0; i < n; i++ {
		if !visited[i] {
			bfs(i)
			ans++
		}
	}
	return ans
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
