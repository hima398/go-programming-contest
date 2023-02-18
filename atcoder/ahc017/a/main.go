package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

//var startTime int64

func main() {
	//開始時間を記録しておく
	//startTime = time.Now().UnixNano()

	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, m, d, k := nextInt(), nextInt(), nextInt(), nextInt()
	var u, v, w []int
	for i := 0; i < m; i++ {
		u = append(u, nextInt())
		v = append(v, nextInt())
		w = append(w, nextInt())
	}
	var x, y []int

	for i := 0; i < n; i++ {
		x = append(x, nextInt())
		y = append(y, nextInt())
	}
	//ans, _ := solve04(n, m, d, k, u, v, w, x, y)
	//ans, _ := solve03(n, m, d, k, u, v, w, x, y)
	ans, _ := solve02(n, m, d, k, u, v, w, x, y)
	//ans := solve01(n, m, d, k, u, v, w, x, y)
	//ans := solveHonestly(n, m, d, k, u, v, w, x, y)
	//ans := firstSolve(n, m, d, k, u, v, w, x, y)
	PrintHorizonaly(ans)
}

//ランダムに工事の計画を立てる
func planRandomly(m, d, k int) []int {
	canRepair := make([]int, d)
	for i := range canRepair {
		canRepair[i] = k
	}
	res := make([]int, m)
	for i := 0; i < m; i++ {
		for {
			target := rand.Intn(d)
			if canRepair[target] > 0 {
				res[i] = target + 1
				canRepair[target]--
				break
			}
		}
	}
	return res
}

//工事の計画を変更する
func changePlan(d, i int, plan []int) []int {
	res := make([]int, len(plan))
	copy(res, plan)
	const pattern = 1
	t := i % pattern
	switch t {
	case 1: //2本の道をランダムに選んで工事の予定を入れ替える
		j, k := rand.Intn(len(plan)), rand.Intn(len(plan))
		res[j], res[k] = res[k], res[j]
	case 2: //ランダムで選んだ境界で入れ替える
		days := make([][]int, d)
		for j, v := range plan {
			days[v-1] = append(days[v-1], j)
		}
		d1, d2 := -1, -1
		for d1 == d2 {
			d1 = rand.Intn(d)
			d2 = rand.Intn(d)
			if len(days[d1]) <= 1 || len(days[d2]) <= 1 {
				d1, d2 = -1, -1
				continue
			}
		}
		r1, r2 := rand.Intn(len(days[d1])-1), rand.Intn(len(days[d2])-1)
		div := Min(r1, r2) + 1
		var t []int
		t = append(t, days[d1][div:]...)
		days[d1] = append(days[d1][:div], days[d2][div:]...)
		days[d2] = append(days[d2][:div], t...)
		for j, v := range days {
			for _, v2 := range v {
				res[v2] = j + 1
			}
		}
	}
	return res
}

//ワーシャルフロイドによるスコアの計算(O(N**3)のためスコア計算には活用できなさそう)
func computeScore(n, m int, u, v, w []int) int {
	const INF = int(1e9)
	dist := make([][]int, n)
	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			if i == j {
				continue
			}
			dist[i][j] = INF
		}
	}
	for i := 0; i < m; i++ {
		ui, vi := u[i]-1, v[i]-1
		dist[ui][vi] = w[i]
		dist[vi][ui] = w[i]
	}
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if i == j || j == k || k == i {
					continue
				}
				dist[i][j] = Min(dist[i][j], dist[i][k]+dist[k][j])
			}
		}
	}
	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			res += dist[i][j]
		}
	}
	return res
}

func computeScoreByUnionFind(n, m, d int, u, v, ans []int) int {
	var res int
	for i := 1; i <= d; i++ {
		uf := NewUnionFind(n)
		for j := 0; j < m; j++ {
			//i日目に工事中
			if ans[j] == i {
				continue
			}
			//すでに同じ連結成分
			if uf.ExistSameUnion(u[j]-1, v[j]-1) {
				continue
			}
			uf.Unite(u[j]-1, v[j]-1)
		}
		//fmt.Println(uf)
		roots := make(map[int]struct{})
		for j := 0; j < n; j++ {
			roots[uf.Find(j)] = struct{}{}
		}
		res += len(roots)
	}
	return res
}

func computeMaxConnectedComponents(n, m, d int, u, v, ans []int) int {
	var res int
	for i := 1; i <= d; i++ {
		uf := NewUnionFind(n)
		for j := 0; j < m; j++ {
			//i日目に工事中
			if ans[j] == i {
				continue
			}
			//すでに同じ連結成分
			if uf.ExistSameUnion(u[j]-1, v[j]-1) {
				continue
			}
			uf.Unite(u[j]-1, v[j]-1)
		}
		//fmt.Println(uf)
		roots := make(map[int]struct{})
		for j := 0; j < n; j++ {
			roots[uf.Find(j)] = struct{}{}
		}
		res = Max(res, len(roots))
	}
	return res
}

func pickUpPoints(n int, x, y []int) []int {
	//0:x最小、1:y最小、2:x最大、3:y最大
	points := make([]int, 4)
	xMin, yMin, xMax, yMax := INF, INF, -1, -1
	for i := 0; i < n; i++ {
		if xMin > x[i] {
			xMin = x[i]
			points[0] = i
		}
		if yMin > y[i] {
			yMin = y[i]
			points[1] = i
		}
		if xMax < x[i] {
			xMax = x[i]
			points[2] = i
		}
		if yMax < y[i] {
			yMax = y[i]
			points[3] = i
		}
	}
	return points
}

func computeDefaultScoreByDijkstra(n, m, d int, u, v, w []int) [][]int {
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}
	g := NewGraph(n, m)
	for i := 0; i < m; i++ {
		g.Add(u[i]-1, v[i]-1, w[i])
	}
	for i := 0; i < n; i++ {
		dist := g.Dijkstra(i)
	}
}

func computeScoreByDijkstra(n, m, d int, u, v, w, ps, ans []int) int {
	//var res int
	res := int(1e9)
	for k := 1; k <= d; k++ {
		g := NewGraph(n, m)
		for i := 0; i < m; i++ {
			//k日目に工事中の道はグラフの辺として結ばない
			if ans[i] == k {
				continue
			}
			g.Add(u[i]-1, v[i]-1, w[i])
		}
		for _, p := range ps {
			dist := g.Dijkstra(p)
			var max int
			for _, v := range dist {
				if v == INF {
					max = int(1e9)
				} else {
					//	res += v
					//}
					max = Max(max, v)
				}
			}
			res = Min(res, max)
		}
	}
	return res
}

func computeScoreByDegree(n, m, d int, u, v, ans []int) int {
	degree := make([][]int, d)
	for k := 0; k < d; k++ {
		degree[k] = make([]int, n)
	}
	for k := 1; k <= d; k++ {
		for i := 0; i < m; i++ {
			//i番目の道がk日目に工事
			if ans[i] == k {
				degree[k-1][u[i]-1]++
				degree[k-1][v[i]-1]++
			}
		}
	}
	res := 0
	for k := 0; k < d; k++ {
		s := 0
		for i := 0; i < n; i++ {
			s = Max(s, degree[k][i])
		}
		res += s
	}
	return res
}

//
// ソルバー
//
func solve04(n, m, d, k int, u, v, w, x, y []int) ([]int, int) {
	//開始時間を記録しておく
	startTime := time.Now().UnixNano()

	ans := planRandomly(m, d, k)

	score1 := computeScoreByUnionFind(n, m, d, u, v, ans)
	i := 0
	for time.Now().UnixNano()-startTime < 25*int64(1e8) {
		next := changePlan(d, i, ans)
		//next := make([]int, len(ans))
		//copy(next, ans)
		//j, k := rand.Intn(len(ans)), rand.Intn(len(ans))
		//next[j], next[k] = next[k], next[j]
		nextScore := computeScoreByUnionFind(n, m, d, u, v, next)
		if nextScore < score1 {
			ans = next
			score1 = nextScore
		}
		i++
	}

	type node struct {
		d, c int
		is   []int
	}
	for time.Now().UnixNano()-startTime < 55*int64(1e8) {
		var ns []node
		for l := 1; l <= d; l++ {
			uf := NewUnionFind(n)
			var is []int
			for i := 0; i < m; i++ {
				if ans[i] == l {
					is = append(is, i)
					continue
				}
				if uf.ExistSameUnion(u[i]-1, v[i]-1) {
					continue
				}
				uf.Unite(u[i]-1, v[i]-1)
			}
			var c int
			for i := 0; i < n; i++ {
				if i == uf.Find(i) {
					c++
				}
			}
			ns = append(ns, node{l, c, is})
		}
		sort.Slice(ns, func(i, j int) bool {
			if ns[i].c == ns[j].c {
				return len(ns[i].is) > len(ns[j].is)
			}
			return ns[i].c > ns[j].c
		})
		//すべての工事日で連結になっている
		if ns[0].c == 1 {
			break
		}

		idx1 := rand.Intn(len(ns[0].is))

		day := ns[len(ns)-1].d
		ans[ns[0].is[idx1]] = day
	}

	cc := computeMaxConnectedComponents(n, m, d, u, v, ans)

	return ans, cc
}

//連結成分を構築しながら工事の計画を立てていく
//連結成分が完全に構成しきれない。
func solve03(n, m, d, k int, u, v, w, x, y []int) ([]int, int) {
	//未修理の道
	var e []Edge
	for i := 0; i < m; i++ {
		e = append(e, Edge{i, u[i], v[i], w[i]})
	}
	sort.Slice(e, func(i, j int) bool {
		return e[i].w < e[j].w
	})
	//修理済みの道
	var repaired []Edge
	ans := make([]int, m)
	mid := k
	for l := 1; l <= d; l++ {
		uf := NewUnionFind(n)
		for _, r := range repaired {
			if uf.ExistSameUnion(r.s, r.t) {
				continue
			}
			uf.Unite(r.s, r.t)
		}
		s := 0
		var next []Edge
		for i := range e {
			if uf.ExistSameUnion(u[i], v[i]) {
				if s < mid {
					//l日目の連結成分に影響せず、k未満なので工事の対象にする
					ans[e[i].i] = l
					s++
					repaired = append(repaired, Edge{e[i].i, e[i].s, e[i].t, e[i].w})
				} else {
					//l日目の連結成分には影響しないが上限のkを超えているので翌日以降に工事をする
					next = append(next, Edge{e[i].i, e[i].s, e[i].t, e[i].w})
				}
			} else {
				//l日目の連結成分に関わるので翌日以降に工事をする
				uf.Unite(u[i], v[i])
				next = append(next, Edge{e[i].i, e[i].s, e[i].t, e[i].w})
			}
		}
		//fmt.Fprintf(out, "day = %d, repaired = %d, rem = %d\n", l, len(repaired), len(next))
		e = next
		sort.Slice(e, func(i, j int) bool {
			return e[i].w < e[j].w
		})
		//for i := 0; i < k-s; i++ {
		//	if len(e) > 0 {
		//		cur := e[0]
		//		e = e[1:]
		//		ans[cur.i] = l
		//		repaired = append(repaired, cur)
		//	}
		//}
		sort.Slice(repaired, func(i, j int) bool {
			return repaired[i].w < repaired[j].w
		})
	}
	cc := computeMaxConnectedComponents(n, m, d, u, v, ans)
	return ans, cc
}

//最初にある程度連結させたあと、ピックアップした点からの次数を減らしていく
func solve02(n, m, d, k int, u, v, w, x, y []int) ([]int, int) {
	//開始時間を記録しておく
	startTime := time.Now().UnixNano()

	ans := planRandomly(m, d, k)

	score1 := computeScoreByUnionFind(n, m, d, u, v, ans)
	for time.Now().UnixNano()-startTime < 25*int64(1e8) {
		next := make([]int, len(ans))
		copy(next, ans)
		j, k := rand.Intn(len(ans)), rand.Intn(len(ans))
		next[j], next[k] = next[k], next[j]
		nextScore := computeScoreByUnionFind(n, m, d, u, v, next)
		if nextScore < score1 {
			ans = next
			score1 = nextScore
		}
	}

	type node struct {
		d, c int
		is   []int
	}
	for time.Now().UnixNano()-startTime < 55*int64(1e8) {
		var ns []node
		for l := 1; l <= d; l++ {
			uf := NewUnionFind(n)
			var is []int
			for i := 0; i < m; i++ {
				if ans[i] == l {
					is = append(is, i)
					continue
				}
				if uf.ExistSameUnion(u[i]-1, v[i]-1) {
					continue
				}
				uf.Unite(u[i]-1, v[i]-1)
			}
			var c int
			for i := 0; i < n; i++ {
				if i == uf.Find(i) {
					c++
				}
			}
			ns = append(ns, node{l, c, is})
		}
		sort.Slice(ns, func(i, j int) bool {
			if ns[i].c == ns[j].c {
				return len(ns[i].is) > len(ns[j].is)
			}
			return ns[i].c > ns[j].c
		})
		//すべての工事日で連結になっている
		if ns[0].c == 1 {
			break
		}
		//for _, node := range ns {
		//	fmt.Println(node.d, node.c, len(node.is))
		//}
		//if len(ns[0].is) == 0 {
		//	fmt.Println(ns[0].d, ns[0].c, len(ns[0].is))
		//	break
		//}
		idx1 := rand.Intn(len(ns[0].is))
		//交換するansのインデックス

		day := ns[len(ns)-1].d
		ans[ns[0].is[idx1]] = day
	}
	/*
		score2 := computeScoreByDegree(n, m, d, u, v, ans)
		var i int
		for i = 0; time.Now().UnixNano()-startTime < 55*int64(1e8); i++ {
			next := make([]int, len(ans))
			copy(next, ans)
			j, k := rand.Intn(len(ans)), rand.Intn(len(ans))
			next[j], next[k] = next[k], next[j]
			//nextScore := computeScoreByDijkstra(n, m, d, u, v, w, ps, next)
			nextScore := computeScoreByDegree(n, m, d, u, v, next)
			//fmt.Println("score, newScore = ", score2, nextScore)
			if nextScore < score2 {
				//fmt.Println("Updated")
				ans = next
				score2 = nextScore
			}
		}
	*/
	//fmt.Println("i = ", i)
	//PrintHorizonaly(ans)
	cc := computeMaxConnectedComponents(n, m, d, u, v, ans)

	return ans, cc
}

//試行回数は1000回程度が限界そう
func solve01(n, m, d, k int, u, v, w, x, y []int) []int {
	//開始時間を記録しておく
	startTime := time.Now().UnixNano()

	ans := planRandomly(m, d, k)

	score := computeScoreByUnionFind(n, m, d, u, v, ans)
	for time.Now().UnixNano()-startTime < 55*int64(1e8) {
		next := make([]int, len(ans))
		copy(next, ans)
		j, k := rand.Intn(len(ans)), rand.Intn(len(ans))
		next[j], next[k] = next[k], next[j]
		nextScore := computeScoreByUnionFind(n, m, d, u, v, next)
		if nextScore < score {
			ans = next
			score = nextScore
		}
	}
	//PrintHorizonaly(s)
	return ans
}

func solveHonestly(n, m, d, k int, u, v, w, x, y []int) []int {
	ans := planRandomly(m, d, k)

	computeScore(n, m, u, v, w)
	return ans
}

//
// アルゴリズム関連
//

//
// Dijkstra
//

const INF = 1 << 60

type Graph struct {
	n, m int
	d    []int    //distance
	v    []bool   // visited
	e    [][]Edge //edges
}

func NewGraph(n, m int) *Graph {
	g := new(Graph)
	g.n = n
	g.m = m
	g.d = make([]int, n)
	for i := range g.d {
		g.d[i] = INF
	}
	g.v = make([]bool, n)
	g.e = make([][]Edge, n)
	return g
}

func (g *Graph) Add(a, b, w int) {
	g.e[a] = append(g.e[a], Edge{0, a, b, w})
	g.e[b] = append(g.e[b], Edge{0, a, b, w})
}

func (g *Graph) Dfs(cur int) {
	g.v[cur] = true
	for _, next := range g.e[cur] {
		if g.v[next.t] {
			continue
		}
		g.Dfs(next.t)
	}
}

func (g *Graph) Bfs(s int) {
	var q []int
	q = append(q, s)
	g.v[s] = true
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		for _, next := range g.e[cur] {
			if g.v[next.t] {
				continue
			}
			q = append(q, next.t)
			g.v[next.t] = true
		}
	}
}

func (g *Graph) Dijkstra(s int) []int {
	q := &PriorityQueue{}
	heap.Init(q)

	push := func(i, t, c int) {
		if g.d[t] <= c {
			return
		}
		g.d[t] = c
		heap.Push(q, Edge{0, i, t, c})
	}
	push(0, 0, 0)

	for q.Len() > 0 {
		cur := heap.Pop(q).(Edge)
		if g.d[cur.t] < cur.w {
			continue
		}
		for _, next := range g.e[cur.t] {
			push(next.s, next.t, cur.w+next.w)
		}
	}
	return g.d
}

type Edge struct {
	i    int
	s, t int
	w    int
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

//
// DSU
//
type UnionFind struct {
	parent []int // parentent numbers
	rank   []int // height of tree
	size   []int
}

func NewUnionFind(n int) *UnionFind {
	if n <= 0 {
		return nil
	}
	u := new(UnionFind)
	// for accessing index without minus 1
	u.parent = make([]int, n+1)
	u.rank = make([]int, n+1)
	u.size = make([]int, n+1)
	for i := 0; i <= n; i++ {
		u.parent[i] = i
		u.rank[i] = 0
		u.size[i] = 1
	}
	return u
}

func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] == x {
		return x
	} else {
		// compress path
		// ex. Find(4)
		// 1 - 2 - 3 - 4
		// 1 - 2
		//  L-3
		//  L 4
		uf.parent[x] = uf.Find(uf.parent[x])
		return uf.parent[x]
	}
}

func (uf *UnionFind) Size(x int) int {
	return uf.size[uf.Find(x)]
}

func (uf *UnionFind) ExistSameUnion(x, y int) bool {
	return uf.Find(x) == uf.Find(y)
}

func (uf *UnionFind) Unite(x, y int) {
	x = uf.Find(x)
	y = uf.Find(y)
	if x == y {
		return
	}
	// rank
	if uf.rank[x] < uf.rank[y] {
		//yがrootの木にxがrootの木を結合する
		uf.parent[x] = y
		uf.size[y] += uf.size[x]
	} else {
		// uf.rank[x] >= uf.rank[y]
		//xがrootの木にyがrootの木を結合する
		uf.parent[y] = x
		uf.size[x] += uf.size[y]
		if uf.rank[x] == uf.rank[y] {
			uf.rank[x]++
		}
	}
}

func PrintUnionFind(u *UnionFind) {
	// for debuging. not optimize.
	fmt.Println(u.parent)
	fmt.Println(u.rank)
	fmt.Println(u.size)
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

func PrintString(x string) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}

func PrintHorizonaly(x []int) {
	defer out.Flush()
	fmt.Fprintf(out, "%d", x[0])
	for i := 1; i < len(x); i++ {
		fmt.Fprintf(out, " %d", x[i])
	}
	fmt.Fprintln(out)
}

func PrintVertically(x []int) {
	defer out.Flush()
	for _, v := range x {
		fmt.Fprintln(out, v)
	}
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

func Floor(x, y int) int {
	return x / y
}

func Ceil(x, y int) int {
	return (x + y - 1) / y
}

func Sqrt(x int) int {
	x2 := int(math.Sqrt(float64(x))) - 1
	for (x2+1)*(x2+1) <= x {
		x2++
	}
	return x2
}
