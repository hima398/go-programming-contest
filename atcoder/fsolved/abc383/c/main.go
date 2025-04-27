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

	h, w, d := nextInt(), nextInt(), nextInt()
	var s []string
	for i := 0; i < h; i++ {
		s = append(s, nextString())
	}

	ans := solve(h, w, d, s)

	Print(ans)
}

func solve(h, w, d int, s []string) int {
	const INF = 1 << 60
	type node struct {
		i, j, d int
	}
	dist := make([][]int, h)
	for i := range dist {
		dist[i] = make([]int, w)
		for j := range dist[i] {
			dist[i][j] = INF
		}
	}
	q := queue.New[node]()
	/*
		q := priorityqueue.New[node](func(a, b node) int {
			if a.d == b.d {
				return 0
			}
			if a.d < b.d {
				return -1
			}
			return 1
		})
	*/
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if s[i][j] == 'H' {
				q.Push(node{i, j, 0})
				dist[i][j] = 0
			}
		}
	}
	di := []int{-1, 0, 1, 0}
	dj := []int{0, -1, 0, 1}
	for !q.Empty() {
		cur := q.Pop()
		for k := 0; k < 4; k++ {
			ni, nj := cur.i+di[k], cur.j+dj[k]
			if ni < 0 || ni >= h || nj < 0 || nj >= w {
				continue
			}
			if s[ni][nj] == '#' {
				continue
			}
			if cur.d+1 > d {
				continue
			}
			if dist[ni][nj] < cur.d+1 {
				continue
			}
			q.Push(node{ni, nj, cur.d + 1})
			dist[ni][nj] = cur.d + 1
		}
	}
	var ans int
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if dist[i][j] <= d {
				ans++
			}
		}
	}
	return ans
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
