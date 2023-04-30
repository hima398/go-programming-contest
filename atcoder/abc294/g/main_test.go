package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

type args struct {
	n int
	u []int
	v []int
	w []int
	q int
	t []int
	x []int
	y []int
}

type testCase struct {
	name string
	args args
}

func Test_solveCommentary(t *testing.T) {
	var tests []testCase
	seed := time.Now().UnixNano()
	rand.Seed(seed)
	for i := 0; i < 1000; i++ {
		var args args
		n := rand.Intn(2*int(1e5)) + 1
		args.n = n
		uf := NewUnionFind(n)
		for uf.Size(0) < n {
			var u, v int
			for u == v || uf.ExistSameUnion(u, v) {
				u = rand.Intn(n)
				v = rand.Intn(n)
				//fmt.Println("u, v = ", u, v)
			}
			uf.Unite(u, v)
			args.u = append(args.u, u)
			args.v = append(args.v, v)
			w := rand.Intn(int(1e9)) + 1
			args.w = append(args.w, w)
		}

		q := rand.Intn(2*int(1e5)) + 1
		args.q = q
		for j := 0; j < q; j++ {
			t := rand.Intn(2) + 1
			var x, y int
			if t == 1 {
				if n == 1 {
					t = 2
					x = rand.Intn(n)
					y = rand.Intn(n)
				} else {
					x = rand.Intn(n - 1)
					y = rand.Intn(int(1e9)) + 1
				}
			} else {
				x = rand.Intn(n)
				y = rand.Intn(n)
			}
			args.t = append(args.t, t)
			args.x = append(args.x, x)
			args.y = append(args.y, y)
		}
		tests = append(tests, testCase{fmt.Sprintf("Random %04d", i), args})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//fmt.Println(tt.args)
			solveCommentary(tt.args.n, tt.args.u, tt.args.v, tt.args.w, tt.args.q, tt.args.t, tt.args.x, tt.args.y)
		})
	}
	fmt.Println("seed = ", seed)
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
