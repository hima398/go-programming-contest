package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/liyue201/gostl/ds/queue"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	h, w := nextInt(), nextInt()
	var a []string
	for i := 0; i < h; i++ {
		a = append(a, nextString())
	}
	n := nextInt()
	var r, c, e []int
	for i := 0; i < n; i++ {
		r = append(r, nextInt()-1)
		c = append(c, nextInt()-1)
		e = append(e, nextInt())
	}

	ok := solve(h, w, a, n, r, c, e)

	if ok {
		Print("Yes")
	} else {
		Print("No")
	}
}

func solve(h, w int, a []string, n int, r, c, e []int) bool {
	var si, sj, ti, tj int
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if a[i][j] == 'S' {
				si, sj = i, j
			}
			if a[i][j] == 'T' {
				ti, tj = i, j
			}
		}
	}
	f := make([][]int, h)
	for i := range f {
		f[i] = make([]int, w)
	}
	for i := 0; i < n; i++ {
		f[r[i]][c[i]] = e[i]
	}
	//スタート地点でエネルギーが確保できない
	if f[si][sj] == 0 {
		return false
	}
	type node struct {
		i, j, e int
	}
	visited := make([][]bool, h)
	for i := range visited {
		visited[i] = make([]bool, w)
	}
	q := queue.New[node]()
	q.Push(node{si, sj, f[si][sj]})
	bfs := func(i, j, e int) {
		q2 := queue.New[node]()
		q2.Push(node{i, j, e})
		v2 := make([][]bool, h)
		for i := range visited {
			v2[i] = make([]bool, w)
		}
		visited[i][j] = true
		v2[i][j] = true
		di := []int{-1, 0, 1, 0}
		dj := []int{0, -1, 0, 1}
		for !q2.Empty() {
			cur := q2.Pop()
			if cur.e == 0 {
				continue
			}
			for k := 0; k < 4; k++ {
				ni, nj := cur.i+di[k], cur.j+dj[k]
				if ni < 0 || ni >= h || nj < 0 || nj >= w {
					continue
				}
				if a[ni][nj] == '#' {
					continue
				}
				if v2[ni][nj] {
					continue
				}
				q2.Push(node{ni, nj, cur.e - 1})
				v2[ni][nj] = true
				visited[ni][nj] = true
				if f[ni][nj] > 0 {
					q.Push(node{ni, nj, f[ni][nj]})
					f[ni][nj] = 0
				}
			}
		}
	}
	for !q.Empty() {
		cur := q.Pop()
		bfs(cur.i, cur.j, cur.e)
	}
	return visited[ti][tj]
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

func Print(x any) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
