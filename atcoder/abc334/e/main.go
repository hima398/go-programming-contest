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

	h, w := nextInt(), nextInt()
	var s []string
	for i := 0; i < h; i++ {
		s = append(s, nextString())
	}

	ans := solve(h, w, s)

	Print(ans)
}

func IsOutGrid(i, j, h, w int) bool {
	return i < 0 || i >= h || j < 0 || j >= w
}

func convertToUnionFindIndex(i, j, w int) int {
	return w*i + j
}

func solve(h, w int, s []string) int {
	const p = 998244353
	di := []int{-1, 0, 1, 0}
	dj := []int{0, -1, 0, 1}

	/*
		visited := make([][]bool, h)
		for i := range visited {
			visited[i] = make([]bool, w)
		}
	*/
	grid := make([][]int, h)
	for i := range grid {
		grid[i] = make([]int, w)
	}
	type cell struct {
		i, j int
	}
	bfs := func(i, j, v int) {
		q := queue.New[cell]()
		q.Push(cell{i, j})
		grid[i][j] = v

		for !q.Empty() {
			cur := q.Pop()
			for k := 0; k < 4; k++ {
				ni, nj := cur.i+di[k], cur.j+dj[k]
				if IsOutGrid(ni, nj, h, w) {
					continue
				}
				if s[ni][nj] == '.' {
					continue
				}
				if grid[ni][nj] > 0 {
					continue
				}
				q.Push(cell{ni, nj})
				grid[ni][nj] = v
			}
		}
	}
	var connectedComponentsNumber int
	//連結成分の数を数える
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if s[i][j] == '.' || grid[i][j] > 0 {
				continue
			}
			connectedComponentsNumber++
			bfs(i, j, connectedComponentsNumber)
		}
	}
	var redCellNumber int
	m := make(map[int]int)
	//赤色の近傍を調べる
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if s[i][j] == '#' {
				continue
			}
			redCellNumber++
			mm := make(map[int]struct{})
			for k := 0; k < 4; k++ {
				ni, nj := i+di[k], j+dj[k]
				if IsOutGrid(ni, nj, h, w) {
					continue
				}
				if grid[ni][nj] == 0 {
					continue
				}
				mm[grid[ni][nj]] = struct{}{}
			}
			k := connectedComponentsNumber - len(mm) + 1
			m[k]++
		}
	}
	//for _, v := range grid {
	//	fmt.Println(v)
	//}
	//テスト出力
	//fmt.Println(m)

	var ans int
	//期待値を計算する
	for k, v := range m {
		ans += ((k * v % p) * Inv(redCellNumber, p)) % p
		ans %= p
	}
	return ans
}

func solve01(h, w int, s []string) int {
	const p = 998244353
	di := []int{-1, 0, 1, 0}
	dj := []int{0, -1, 0, 1}

	uf := NewUnionFind(h * w)
	//連結成分の数を数える
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if s[i][j] == '.' {
				continue
			}
			for k := 0; k < 4; k++ {
				ni, nj := i+di[k], j+dj[k]
				if IsOutGrid(ni, nj, h, w) {
					continue
				}
				if s[ni][nj] == '.' {
					continue
				}
				u, v := convertToUnionFindIndex(i, j, w), convertToUnionFindIndex(ni, nj, w)
				if uf.ExistSameUnion(u, v) {
					continue
				}
				uf.Unite(u, v)
			}
		}
	}

	var connectedComponentsNumber int
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			idx := convertToUnionFindIndex(i, j, w)
			if idx == uf.Find(idx) {
				connectedComponentsNumber++
			}
		}
	}

	//赤色の近傍を調べる

	//期待値を計算する

	var ans int
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

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func Print(x any) {
	defer out.Flush()
	fmt.Fprintln(out, x)
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
