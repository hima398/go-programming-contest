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

	n, m := nextInt(), nextInt()
	var u, v []int
	for i := 0; i < m; i++ {
		u = append(u, nextInt()-1)
		v = append(v, nextInt()-1)
	}
	w := nextIntSlice(n)
	a := nextIntSlice(n)

	ans := solve(n, m, u, v, w, a)

	Print(ans)
}

func solve(n, m int, u, v, w, a []int) int {
	dag := NewDirectedAcyclicGraph()
	dag.Init(n)
	for i := 0; i < m; i++ {
		if w[u[i]] < w[v[i]] {
			dag.AddEdge(u[i], v[i], w[i])
		} else if w[u[i]] > w[v[i]] {
			dag.AddEdge(v[i], u[i], w[i])
		}
	}

	return dag.computeMaxOperation(w, a)
}

type WeightedEdge struct {
	t, w int
}

type DirectedAcyclicGraph struct {
	n int              //頂点数
	e [][]WeightedEdge //辺
	d []int            //入次数
}

func NewDirectedAcyclicGraph() *DirectedAcyclicGraph {
	return new(DirectedAcyclicGraph)
}

// 頂点数を代入して初期化する
func (dag *DirectedAcyclicGraph) Init(n int) {
	dag.n = n
	dag.e = make([][]WeightedEdge, n)
	dag.d = make([]int, n)
}

func (dag *DirectedAcyclicGraph) AddEdge(s, t, w int) {
	dag.e[s] = append(dag.e[s], WeightedEdge{t, w})
	dag.d[t]++
}

func (dag *DirectedAcyclicGraph) computeMaxOperation(w, a []int) int {
	var mx int
	for _, wi := range w {
		mx = Max(mx, wi)
	}
	//頂点iにwyの総和がjになるように集合Sを選んだ時のコマの最大値
	dp := make([][]int, dag.n)
	dp2 := make([][]bool, dag.n)
	for i := 0; i < dag.n; i++ {
		dp[i] = make([]int, mx+1)
		dp2[i] = make([]bool, mx+1)
		dp2[i][0] = true
	}
	q := queue.New[int]()
	for i := 0; i < dag.n; i++ {
		if dag.d[i] == 0 {
			q.Push(i)
		}
	}
	for !q.Empty() {
		cur := q.Pop()
		for _, par := range dag.e[cur] {
			dag.d[par.t]--
			if dag.d[par.t] == 0 {
				q.Push(par.t)
			}
			for j := 0; j <= w[par.t]-w[cur]; j++ {
				if !dp2[cur][j] {
					continue
				}
				nj := j + w[cur]
				dp[par.t][nj] += dp[cur][j] + a[cur]
				dp2[par.t][nj] = true
			}
		}
	}
	var ans int
	for i := 0; i < dag.n; i++ {
		fmt.Println(dp[i])
		fmt.Println(dp2[i])
		for j := 0; j <= mx; j++ {
			if !dp2[i][j] {
				continue
			}
			ans = Max(ans, dp[i][j]+a[i])
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
