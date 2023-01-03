package main

import (
	"bufio"
	"container/heap"
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

	n, m, k := nextInt(), nextInt(), nextInt()
	var u, v, a []int
	for i := 0; i < m; i++ {
		u = append(u, nextInt()-1)
		v = append(v, nextInt()-1)
		a = append(a, nextInt())
	}
	s := nextIntSlice(k)
	ans := solveBfs(n, m, k, u, v, a, s)
	//ans := solve(n, m, k, u, v, a, s)
	PrintInt(ans)
}

func solveBfs(n, m, k int, u, v, a, s []int) int {
	const INF = 1 << 60

	e := make([][]int, 2*n)
	for i := 0; i < m; i++ {
		s, t := n*a[i]+u[i], n*a[i]+v[i]
		e[s] = append(e[s], t)
		e[t] = append(e[t], s)
	}
	for _, si := range s {
		si--
		e[si] = append(e[si], si+n)
		e[si+n] = append(e[si+n], si)
	}
	dist := make([]int, 2*n)
	for i := 0; i < 2*n; i++ {
		dist[i] = INF
	}
	q := new(Queue)
	q.Push(n)
	dist[n] = 0
	for q.Len() > 0 {
		//fmt.Println(q)
		cur := q.Pop()
		for _, next := range e[cur] {
			//訪問済
			if dist[next] < INF {
				continue
			}
			if Abs(next-cur) == n {
				//スイッチを押す移動
				//普通のBFSのようにQueueの後ろに積んではいけなくて
				q.Push(next)
				//本当は優先して前に入れないといけない
				//q.PushFront(next)
				dist[next] = dist[cur]
			} else {
				//
				q.Push(next)
				dist[next] = dist[cur] + 1
			}
		}
	}
	ans := Min(dist[n-1], dist[2*n-1])
	if ans == INF {
		ans = -1
	}
	return ans
}

type Queue struct {
	front, rear []int
}

func (q *Queue) Len() int {
	return len(q.front) + len(q.rear)
}

func (q *Queue) Push(x int) {
	q.rear = append(q.rear, x)
}

func (q *Queue) PushFront(x int) {
	q.front = append(q.front, x)
}

func (q *Queue) Pop() int {
	if len(q.front) > 0 {
		res := q.front[len(q.front)-1]
		q.front = q.front[:len(q.front)-1]
		return res
	} else if len(q.rear) > 0 {
		res := q.rear[0]
		q.rear = q.rear[1:]
		return res
	}
	return -1
}

func solve(n, m, k int, u, v, a, s []int) int {
	const INF = 1 << 60
	dist := make([]int, 2*n)
	for i := range dist {
		dist[i] = INF
	}

	e := make([][]Edge, 2*n)
	for i := 0; i < m; i++ {
		if a[i] == 0 {
			e[u[i]+n] = append(e[u[i]+n], Edge{v[i] + n, 1})
			e[v[i]+n] = append(e[v[i]+n], Edge{u[i] + n, 1})
		} else {
			e[u[i]] = append(e[u[i]], Edge{v[i], 1})
			e[v[i]] = append(e[v[i]], Edge{u[i], 1})
		}
	}
	for _, si := range s {
		si--
		e[si] = append(e[si], Edge{si + n, 0})
		e[si+n] = append(e[si+n], Edge{si, 0})
	}

	//Dijkstra
	q := &PriorityQueue{}
	heap.Init(q)

	push := func(to, cost int) {
		if dist[to] <= cost {
			return
		}
		dist[to] = cost
		heap.Push(q, Edge{to, cost})
	}
	push(0, 0)
	for q.Len() > 0 {
		cur := heap.Pop(q).(Edge)
		if dist[cur.t] < cur.w {
			continue
		}
		for _, next := range e[cur.t] {
			push(next.t, cur.w+next.w)
		}
	}
	ans := Min(dist[n-1], dist[2*n-1])
	//for i := range dist {
	//	if dist[i] == INF {
	//		dist[i] = -1
	//	}
	//}
	if ans == INF {
		ans = -1
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
	return i
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
