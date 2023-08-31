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

	n := nextInt()
	var p, q []int
	for i := 0; i < n-1; i++ {
		p = append(p, nextInt()-1)
		q = append(q, nextInt()-1)
	}

	ans := solve(n, p, q)

	PrintHorizonaly(ans)
}

func solve(n int, p, q []int) []int {
	const mod = 998244353
	type edge struct {
		t, w int
	}
	uf := NewUnionFind(n)
	//UnionFindの親となっているノードが最後に出た試合のインデックス
	mx := make([]int, 2*n)
	for i := 0; i < n; i++ {
		mx[i] = i
	}
	e := make([][]edge, 2*n)
	//試合ごとにチームをマージしながらトーナメント表を構築する
	for i := 0; i < n-1; i++ {
		//試合を表すノードのインデックス
		idx := n + i
		//idxの子ノードのインデックス
		c1, c2 := mx[uf.Find(p[i])], mx[uf.Find(q[i])]
		//各チームの勝率を計算
		w1 := uf.Size(p[i])
		w2 := uf.Size(q[i])
		s := w1 + w2
		w1 = w1 * Inv(s, mod) % mod
		w2 = w2 * Inv(s, mod) % mod
		//チームの統合
		uf.Unite(p[i], q[i])
		//トーナメント表を構築
		e[idx] = append(e[idx], edge{c1, w1})
		e[idx] = append(e[idx], edge{c2, w2})
		mx[uf.Find(p[i])] = idx
		mx[uf.Find(q[i])] = idx
	}

	//トーナメント表を上から辿って解答を作成する
	ans := make([]int, 2*n)

	bfs := func(cur int) {
		q := queue.New[edge]()
		q.Push(edge{cur, 0})
		for !q.Empty() {
			cur := q.Pop()
			for _, next := range e[cur.t] {
				q.Push(edge{next.t, (cur.w + next.w) % mod})
				ans[next.t] = (cur.w + next.w) % mod
			}
		}
	}
	bfs(2 * (n - 1))
	ans = ans[:n]
	return ans
}

type UnionFind struct {
	parent []int // parentent numbers
	rank   []int // height of tree
	size   []int
}

func New(n int) *UnionFind {
	return NewUnionFind(n)
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

func PrintHorizonaly(x []int) {
	defer out.Flush()
	fmt.Fprintf(out, "%d", x[0])
	for i := 1; i < len(x); i++ {
		fmt.Fprintf(out, " %d", x[i])
	}
	fmt.Fprintln(out)
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

func Pow(x, y, p int) int {
	ret := 1
	for y > 0 {
		if y%2 == 1 {
			ret = ret * x % p
		}
		y >>= 1
		x = x * x % p
	}
	return ret
}

func Inv(x, p int) int {
	return Pow(x, p-2, p)
}
