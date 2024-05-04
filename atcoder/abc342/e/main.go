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

type answer struct {
	fk        int
	reachable bool
}

func main() {
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, m := nextInt(), nextInt()
	var l, d, k, c, a, b []int
	for i := 0; i < m; i++ {
		l = append(l, nextInt())
		d = append(d, nextInt())
		k = append(k, nextInt())
		c = append(c, nextInt())
		a = append(a, nextInt()-1)
		b = append(b, nextInt()-1)
	}

	ans := solve(n, m, l, d, k, c, a, b)

	for _, v := range ans {
		if !v.reachable {
			Print("Unreachable")
		} else {
			Print(v.fk)
		}
	}
}

func solve(n, m int, l, d, k, c, a, b []int) []answer {
	const INF = 2 * int(1e18)

	//f(x)を示すスライス
	f := make([]int, n)
	for i := range f {
		f[i] = -INF
	}

	//辺の構造を変換
	type edge struct {
		to         int
		l, d, k, c int
	}
	e := make([][]edge, n)
	//解はNから順に求めていくので有向辺を逆に張る
	for i := 0; i < m; i++ {
		e[b[i]] = append(e[b[i]], edge{a[i], l[i], d[i], k[i], c[i]})
	}

	type node struct {
		fi int //f(i)の値
		e  edge
	}
	//Dijkstra
	q := priorityqueue.New[node](func(a, b node) int {
		if a.fi > b.fi {
			return -1
		} else if a.fi < b.fi {
			return 1
		} else {
			return 0
		}
	})

	push := func(to int, node node) {
		if f[to] > node.fi {
			return
		}
		f[to] = node.fi
		q.Push(node)
	}
	push(n-1, node{INF, edge{to: n - 1}})
	for !q.Empty() {
		cur := q.Pop()
		if f[cur.e.to] > cur.fi {
			continue
		}
		for _, next := range e[cur.e.to] {
			x := Floor(f[cur.e.to]-next.c-next.l, next.d)
			if x < 0 {
				continue
			}
			if x >= next.k {
				x = next.k - 1
			}
			push(next.to, node{next.l + x*next.d, next})
		}
	}
	var ans []answer
	for i := 0; i < n-1; i++ {
		if f[i] == -INF {
			ans = append(ans, answer{-INF, false})
		} else {
			ans = append(ans, answer{f[i], true})
		}
	}
	return ans
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

func Floor(x, y int) int {
	return x / y
}
