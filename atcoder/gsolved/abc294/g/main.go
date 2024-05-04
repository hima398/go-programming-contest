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
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	var u, v, w []int
	for i := 0; i < n-1; i++ {
		u = append(u, nextInt()-1)
		v = append(v, nextInt()-1)
		w = append(w, nextInt())
	}
	q := nextInt()
	var t, x, y []int
	for i := 0; i < q; i++ {
		t = append(t, nextInt())
		x = append(x, nextInt())
		y = append(y, nextInt())
		x[i]--
		if t[i] == 2 {
			y[i]--
		}
	}
	ans := solveCommentary(n, u, v, w, q, t, x, y)
	PrintVertically(ans)
}

type Edge struct {
	i, u, v, w int
}

func solveCommentary(n int, u, v, w []int, q int, t, x, y []int) []int {
	e = make([][]Edge, n)
	for i := 0; i < n-1; i++ {
		e[u[i]] = append(e[u[i]], Edge{i, u[i], v[i], w[i]})
		e[v[i]] = append(e[v[i]], Edge{i, v[i], u[i], w[i]})
	}
	//[i][0]:i番目に使う辺、[i][1]:i番目に訪れる頂点
	var eulerTour [][2]int
	var dfs func(cur, par int)
	dfs = func(cur, par int) {
		for _, next := range e[cur] {
			if next.v == par {
				continue
			}
			eulerTour = append(eulerTour, [2]int{next.i, cur})
			dfs(next.v, cur)
			eulerTour = append(eulerTour, [2]int{next.i, next.v})
		}
	}
	dfs(0, 0)
	//fmt.Println("euler tour = ", eulerTour)
	var depth [][2]int
	depth = append(depth, [2]int{0, 0})
	var dist []int
	edgeToIndex := make([][2]int, n-1)
	for i := 0; i < n-1; i++ {
		edgeToIndex[i][0] = 3 * n
		edgeToIndex[i][1] = 3 * n
	}
	vertexToIndex := make([]int, n)
	for i := 0; i < n; i++ {
		vertexToIndex[i] = 3 * n
	}
	vertexToIndex[0] = 0
	for i := 0; i < 2*(n-1); i++ {
		//for i := 0; i < len(eulerTour); i++ {
		edge, vertex := eulerTour[i][0], eulerTour[i][1]
		nextVertex := u[edge] ^ v[edge] ^ vertex
		vertexToIndex[nextVertex] = Min(vertexToIndex[nextVertex], i+1)
		if edgeToIndex[edge][0] > 2*n {
			dist = append(dist, w[edge])
			edgeToIndex[edge][0] = i
			m := len(depth) - 1
			depth = append(depth, [2]int{depth[m][0] + 1, depth[m][1] + 1})
		} else {
			dist = append(dist, -w[edge])
			edgeToIndex[edge][1] = i
			m := len(depth) - 1
			depth = append(depth, [2]int{depth[m][0] - 1, depth[m][1] + 1})
		}
	}
	//fmt.Println(dist)
	//fmt.Println(edgeToIndex)

	distFromRoot := NewFenwickTree(2 * (n - 1))
	for i, v := range dist {
		distFromRoot.Update(i, v)
	}
	//fmt.Println(distFromRoot.nodes)
	BinarySearchLCAInit(n)
	lcaDepth := NewSegmentTree(len(depth), 1<<60, func(x1, x2 [2]int) [2]int {
		if x1[0] == x2[0] {
			if x1[1] <= x2[1] {
				return x1
			} else {
				return x2
			}
		} else {
			if x1[0] <= x2[0] {
				return x1
			} else {
				return x2
			}
		}
	})
	for i, v := range depth {
		lcaDepth.Update(i, v)
	}

	var ans []int
	for k := 0; k < q; k++ {
		switch t[k] {
		case 1:
			first, last := edgeToIndex[x[k]][0], edgeToIndex[x[k]][1]
			//fmt.Println("x, first, last = ", x[k], first, last)
			beforeCost := distFromRoot.Sum(first, first+1)
			distFromRoot.Update(first, y[k]-beforeCost)
			distFromRoot.Update(last, -y[k]+beforeCost)
		case 2:
			//lca := vertexToIndex[BinarySearchLCA(x[k], y[k])]
			//fmt.Println("x, y, lca = ", x[k], y[k], lca)
			u, v := vertexToIndex[x[k]], vertexToIndex[y[k]]
			lca := lcaDepth.Query(Min(u, v), Max(u, v)+1)[1]
			//fmt.Println("u, v, lca = ", u, v, lca)
			a := distFromRoot.Sum(0, u) + distFromRoot.Sum(0, v) - 2*distFromRoot.Sum(0, lca)
			ans = append(ans, a)
		}
	}
	return ans
}

const MaxNode = 2 * int(1e5)
const MaxLogNode = 17 //16.61
var parent [MaxLogNode + 1][MaxNode + 1]int
var depth [MaxNode + 1]int

var e [][]Edge

type Node struct {
	i, d int
}

func Dfs(i, par, d int) {
	parent[0][i] = par
	depth[i] = d

	for _, next := range e[i] {
		if next.v == par {
			continue
		}
		Dfs(next.v, i, d+1)
	}
}

func BinarySearchLCAInit(v int) {
	Dfs(0, -1, 0)
	for i := 0; i < MaxLogNode; i++ {
		for j := 0; j < v; j++ {
			if parent[i][j] < 0 {
				parent[i+1][j] = -1
			} else {
				parent[i+1][j] = parent[i][parent[i][j]]
			}
		}
	}
}

func BinarySearchLCA(u, v int) int {
	// depth(u) <= depth(v)になるように調整する
	if depth[u] > depth[v] {
		u, v = v, u
	}
	for i := 0; i < MaxLogNode; i++ {
		if ((depth[v]-depth[u])>>i)&1 == 1 {
			v = parent[i][v]
		}
	}
	if u == v {
		return u
	}
	for i := MaxLogNode - 1; i >= 0; i-- {
		//fmt.Println("i, u, v = ", i, u, v)
		if parent[i][u] != parent[i][v] {
			u = parent[i][u]
			v = parent[i][v]
		}
	}
	return parent[0][u]

}

type FenwickTree struct {
	n     int
	nodes []int
	//eval  func(x1, x2 int) int
}

func NewFenwickTree(n int) *FenwickTree {
	fen := new(FenwickTree)
	// 1-indexed
	fen.n = n + 1
	fen.nodes = make([]int, fen.n)
	//bt.eval = f
	return fen
}

// i(0-indexed)をvに更新する
func (fen *FenwickTree) Update(i, v int) {
	//内部では1-indexedなのでここでインクリメントする
	i++
	for i <= fen.n {
		fen.nodes[i-1] = fen.nodes[i-1] + v //fen.eval(fen.nodes[i], v)
		i += i & -i
	}
}

// i(0-indexed)の値を取得する
func (fen *FenwickTree) Query(i int) int {
	//i++
	res := 0
	for i > 0 {
		res = fen.nodes[i-1] + res //fen.eval(fen.nodes[i], res)
		i -= i & -i
	}
	return res
}

func (fen *FenwickTree) Sum(l, r int) int {
	return fen.Query(r) - fen.Query(l)
}

type SegmentTree struct {
	size  int
	nodes [][2]int
	f     func(x1, x2 [2]int) [2]int
	inf   int
}

func NewSegmentTree(n, inf int, f func(x1, x2 [2]int) [2]int) (st *SegmentTree) {
	st = new(SegmentTree)
	st.size = 1
	for st.size < n {
		st.size *= 2
	}
	st.nodes = make([][2]int, 2*st.size)
	for i := range st.nodes {
		st.nodes[i][0] = inf
		st.nodes[i][1] = inf
	}
	st.inf = inf
	st.f = f
	return st
}

func (this *SegmentTree) QueryRecursively(a, b, k, l, r int) [2]int {
	// [a, b)と[l, r)が交差しない
	if a >= r || b <= l {
		return [2]int{this.inf, this.inf}
	}

	// [a, b)が[l, r)を完全に含んでいる
	if a <= l && b >= r {
		return this.nodes[k]
	}

	vl := this.QueryRecursively(a, b, 2*k, l, (l+r)/2)
	vr := this.QueryRecursively(a, b, 2*k+1, (l+r)/2, r)
	return this.f(vl, vr) //Max(vl, vr)
}

// [l, r)の区間の値をxorした結果を返す
func (this *SegmentTree) Query(l, r int) [2]int {
	return this.QueryRecursively(l, r, 1, 0, this.size)
}

func (this *SegmentTree) Update(k int, x [2]int) {
	k += this.size
	this.nodes[k] = this.f(this.nodes[k], x)
	for k > 1 {
		k /= 2
		this.nodes[k] = this.f(this.nodes[k*2], this.nodes[k*2+1])
	}
	//fmt.Println(this.nodes)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
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
