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

	n, m, x := nextInt(), nextInt(), nextInt()
	var u, v []int
	for i := 0; i < m; i++ {
		u, v = append(u, nextInt()-1), append(v, nextInt()-1)
	}

	ans := solve(n, m, x, u, v)

	Print(ans)
}

func solve(n, m, x int, u, v []int) int {
	const INF = 1 << 60

	e := make([][]int, n)
	re := make([][]int, n)
	for i := 0; i < m; i++ {
		e[u[i]] = append(e[u[i]], v[i])
		re[v[i]] = append(re[v[i]], u[i])
	}
	type node struct {
		i, cost    int
		isReversed int
	}

	q := priorityqueue.New[node](func(a, b node) int {
		if a.cost == b.cost {
			return 0
		}
		if a.cost < b.cost {
			return -1
		}
		return 1
	})

	//Dijkstra
	dist := make([][2]int, n)
	for i := range dist {
		for j := range dist[i] {
			dist[i][j] = INF
		}
	}
	push := func(to, cost, isReversed int) {
		if dist[to][isReversed] <= cost {
			return
		}
		dist[to][isReversed] = cost
		q.Push(node{to, cost, isReversed})
	}
	push(0, 0, 0)
	push(0, x, 1)
	for !q.Empty() {
		cur := q.Pop()
		if dist[cur.i][cur.isReversed] < cur.cost {
			continue
		}
		switch cur.isReversed {
		case 0:
			for _, next := range e[cur.i] {
				push(next, cur.cost+1, 0)
			}
			for _, next := range re[cur.i] {
				push(next, cur.cost+x+1, 1)
			}
		case 1:
			for _, next := range e[cur.i] {
				push(next, cur.cost+x+1, 0)
			}
			for _, next := range re[cur.i] {
				push(next, cur.cost+1, 1)
			}
		}
	}
	return Min(dist[n-1][0], dist[n-1][1])
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
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
