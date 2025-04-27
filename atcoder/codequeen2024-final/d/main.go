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

	n, m, k := nextInt(), nextInt(), nextInt()
	var a, b, d []int
	for i := 0; i < m; i++ {
		a = append(a, nextInt()-1)
		b = append(b, nextInt()-1)
		d = append(d, nextInt())
	}
	var c, t []int
	for i := 0; i < k; i++ {
		c = append(c, nextInt()-1)
		t = append(t, nextInt())
	}

	ans := solve(n, m, k, a, b, d, c, t)

	Print(ans)
}

type edge struct {
	to, weight int
}

func Dijkstra(n, s int, e [][]edge) []int {
	const INF = 1 << 60

	dist := make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = INF
	}

	//Dijkstra
	q := priorityqueue.New[edge](func(a, b edge) int {
		if a.weight == b.weight {
			return 0
		}
		if a.weight < b.weight {
			return -1
		}
		return 1
	})
	push := func(to, cost int) {
		if dist[to] <= cost {
			return
		}
		dist[to] = cost
		q.Push(edge{to, cost})
	}
	push(s, 0)
	for !q.Empty() {
		cur := q.Pop()
		if dist[cur.to] < cur.weight {
			continue
		}
		for _, next := range e[cur.to] {
			push(next.to, cur.weight+next.weight)
		}
	}

	return dist

}

func solve(n, m, k int, a, b, d, c, t []int) int {
	var ans int
	return ans
}

func first(n, m, k int, a, b, d, c, t []int) int {
	e := make([][]edge, n)
	for i := 0; i < m; i++ {
		e[a[i]] = append(e[a[i]], edge{b[i], d[i]})
		e[b[i]] = append(e[b[i]], edge{a[i], d[i]})
	}

	dist := make([][]int, n)
	for i := 0; i < n; i++ {
		dist[i] = Dijkstra(n, i, e)
	}
	//for _, v := range dist {
	//	fmt.Println(v)
	//}

	type node struct {
		i   int //現在地
		t   int
		cnt int //イベントに参加できた回数
	}
	q := priorityqueue.New[node](func(a, b node) int {
		if a.t == b.t {
			return 0
		}
		if a.t < b.t {
			return -1
		}
		return 1
	})

	var ans int
	for i := 0; i < k; i++ {
		nextT := Max(dist[0][c[i]], t[i])
		//時間を過ぎてしまうので到達できない
		if nextT > t[i] {
			continue
		}
		q.Push(node{c[i], nextT, 1})
		ans = 1
	}

	for !q.Empty() {
		//fmt.Println("len = ", q.Size())
		cur := q.Pop()
		//fmt.Println(cur)
		//次に向かうイベント
		for j := 0; j < k; j++ {
			//fmt.Println(cur.i, c[j], t[j])
			nextT := Max(cur.t+dist[cur.i][c[j]], t[j])
			//時間を過ぎてしまうので到達できない
			if nextT > t[j] {
				continue
			}
			if cur.i == c[j] && cur.t == nextT {
				continue
			}
			q.Push(node{c[j], nextT, cur.cnt + 1})
			ans = Max(ans, cur.cnt+1)
		}
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

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
