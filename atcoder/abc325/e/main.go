package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/liyue201/gostl/ds/priorityqueue"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, a, b, c := nextInt(), nextInt(), nextInt(), nextInt()
	var d [][]int
	for i := 0; i < n; i++ {
		d = append(d, nextIntSlice(n))
	}
	ans := solve(n, a, b, c, d)
	Print(ans)
}

func Dijkstra(n, start int, cost [][]int) []int {
	const INF = 1 << 60

	//都市nに電車で着くまでの最小時間
	dist := make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = INF
	}
	type Edge struct {
		t, cost int
	}
	//Dijkstra
	q := priorityqueue.New[Edge](func(a, b Edge) int {
		if a.cost == b.cost {
			return 0
		}
		if a.cost < b.cost {
			return -1
		}
		return 1
	})
	push := func(to, cost int) {
		if dist[to] <= cost {
			return
		}
		dist[to] = cost
		q.Push(Edge{to, cost})
	}
	push(start, 0)
	for !q.Empty() {
		cur := q.Pop()
		if dist[cur.t] < cur.cost {
			continue
		}
		for next := 0; next < n; next++ {
			push(next, cur.cost+cost[cur.t][next])
		}
	}

	return dist
}

func solve(n, a, b, c int, d [][]int) int {
	const INF = 1 << 60

	c1 := make([][]int, n)
	for i := range c1 {
		c1[i] = make([]int, n)
		for j := range c1[i] {
			c1[i][j] = d[i][j] * a
		}
	}
	//都市1から都市nに社用車で行くまでの最小時間
	d1 := Dijkstra(n, 0, c1)

	c2 := make([][]int, n)
	for i := range c2 {
		c2[i] = make([]int, n)
		for j := range c2[i] {
			c2[i][j] = d[i][j]*b + c
		}
	}
	//都市nに電車で行くまでの最小時間
	d2 := Dijkstra(n, n-1, c2)

	ans := INF
	for i := 0; i < n; i++ {
		ans = Min(ans, d1[i]+d2[i])
	}
	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func nextIntSlice(n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = nextInt()
	}
	return s
}

func Print(x any) {
	defer out.Flush()
	fmt.Fprintln(out, x)
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
