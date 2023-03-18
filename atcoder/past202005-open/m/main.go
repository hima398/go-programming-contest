package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, m := nextInt(), nextInt()
	var u, v []int
	for i := 0; i < m; i++ {
		u = append(u, nextInt()-1)
		v = append(v, nextInt()-1)
	}
	s := nextInt() - 1
	k := nextInt()
	t := nextIntSlice(k)
	for i := range t {
		t[i]--
	}
	ans := solve(n, m, u, v, s, k, t)
	PrintInt(ans)
}

func solve(n, m int, u, v []int, s, k int, t []int) int {
	const INF = 1 << 60
	e := make([][]Edge, n)
	for i := 0; i < m; i++ {
		e[u[i]] = append(e[u[i]], Edge{v[i], 1})
		e[v[i]] = append(e[v[i]], Edge{u[i], 1})
	}
	dist := make(map[int][]int)
	bfs := func(s int) {
		dist[s] = make([]int, n)
		for i := range dist[s] {
			dist[s][i] = INF
		}
		var q []int
		q = append(q, s)
		dist[s][s] = 0
		for len(q) > 0 {
			cur := q[0]
			q = q[1:]
			for _, next := range e[cur] {
				//到達済み
				if dist[s][next.t] < INF {
					continue
				}
				q = append(q, next.t)
				dist[s][next.t] = dist[s][cur] + 1
			}
		}
	}
	/*
		dijkstra := func(s int) {
			dist[s] = make([]int, n)
			for i := range dist[s] {
				dist[s][i] = INF
			}
			q := &PriorityQueue{}
			heap.Init(q)

			push := func(to, cost int) {
				if dist[s][to] <= cost {
					return
				}
				dist[s][to] = cost
				heap.Push(q, Edge{to, cost})
			}
			push(s, 0)
			for q.Len() > 0 {
				cur := heap.Pop(q).(Edge)
				if dist[s][cur.t] < cur.w {
					continue
				}
				for _, next := range e[cur.t] {
					push(next.t, cur.w+next.w)
				}
			}
		}
	*/
	bfs(s)
	//dijkstra(s)
	for _, v := range t {
		bfs(v)
		//dijkstra(v)
	}
	var ps []int
	for i := 0; i < 1<<k; i++ {
		ps = append(ps, i)
	}
	sort.SliceStable(ps, func(i, j int) bool {
		return bits.OnesCount(uint(ps[i])) < bits.OnesCount(uint(ps[j]))
	})
	//すべて0を取り除いておく
	ps = ps[1:]
	dp := make([][]int, 1<<k)
	for i := 0; i < 1<<k; i++ {
		dp[i] = make([]int, k)
		for j := 0; j < k; j++ {
			dp[i][j] = INF
		}
	}
	for i, ti := range t {
		next := 1 << i
		dp[next][i] = dist[s][ti]
	}
	for _, pat := range ps {
		for i, ti := range t {
			for j, tj := range t {
				if i == j {
					continue
				}
				if dp[pat][i] == INF {
					continue
				}
				if pat>>j&1 > 0 {
					continue
				}
				next := pat | (1 << j)
				dp[next][j] = Min(dp[next][j], dp[pat][i]+dist[ti][tj])
			}
		}
	}
	ans := INF
	for i := 0; i < k; i++ {
		ans = Min(ans, dp[1<<k-1][i])
	}
	return ans
}

type Edge struct {
	t, w int
}

type PriorityQueue []Edge

func (pq PriorityQueue) Len() int {
	return len(pq)
}
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].w < pq[j].w
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(item interface{}) {
	*pq = append(*pq, item.(Edge))
}

func (pq *PriorityQueue) Pop() interface{} {
	es := *pq // EdgeのSlice
	n := len(es)
	item := es[n-1]
	*pq = es[0 : n-1]
	return item
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

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
