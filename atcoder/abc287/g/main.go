package main

import (
	"bufio"
	"fmt"
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

	n := nextInt()
	var a, b []int
	for i := 0; i < n; i++ {
		a = append(a, nextInt())
		b = append(b, nextInt())
	}
	q := nextInt()
	t, x, y := make([]int, q), make([]int, q), make([]int, q)
	for i := 0; i < q; i++ {
		t[i] = nextInt()
		x[i] = nextInt()
		if t[i] == 1 || t[i] == 2 {
			y[i] = nextInt()
		}
	}
	ans := solveHonestly(n, a, b, q, t, x, y)
	//ans := solve(n, a, b, q, t, x, y)
	PrintVertically(ans)
}

func solve(n int, a, b []int, q int, t, x, y []int) []int {
	mi := make(map[int]struct{})
	for _, ai := range a {
		mi[ai] = struct{}{}
	}
	for i := range t {
		if t[i] == 1 {
			mi[y[i]] = struct{}{}
		}
	}
	var idxs []int
	for k := range mi {
		idxs = append(idxs, k)
	}
	comp := NewCompress()
	comp.Init(idxs)

	ft1 := NewFenwickTree(comp.Size())
	ft2 := NewFenwickTree(comp.Size())
	for i := range a {
		idx := comp.GetIndex(a[i])
		ft1.Update(idx, b[i])
		ft2.Update(idx, a[i]*b[i])
	}

	var ans []int
	for i := 0; i < q; i++ {
		switch t[i] {
		case 1:
			idx := comp.GetIndex(x[i] - 1)
			ft1.Update(idx, -b[x[i]])
			ft1.Update(idx, -a[x[i]]*b[x[i]])
			a[x[i]] = y[i]

		case 2:
		case 3:
			ok, ng := -1, comp.Size()
			for ng-ok > 1 {
				mid := (ok + ng) / 2
				if ft1.Sum(mid, comp.Size()) >= x[i] {
					ok = mid
				} else {
					ng = mid
				}
			}
			if ok == -1 {
				ans = append(ans, -1)
			} else {
				v := ft2.Sum(ok+1, comp.Size())
				v += x[i] - ft1.Sum(ok+1, comp.Size())*comp.x[ok]
				ans = append(ans, v)
			}
		}
	}
	return ans
}

type Compress struct {
	//重複除去済みの圧縮元
	x []int
}

func NewCompress() *Compress {
	return new(Compress)
}

func (c *Compress) Init(x []int) {
	m := make(map[int]struct{})
	for _, v := range x {
		m[v] = struct{}{}
	}
	for k := range m {
		c.x = append(c.x, k)
	}
	sort.Ints(c.x)
}

func (c *Compress) GetIndex(x int) int {
	return sort.Search(len(c.x), func(i int) bool {
		return c.x[i] >= x
	})
}

func (c *Compress) Size() int {
	return len(c.x)
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

//i(0-indexed)をvに更新する
func (fen *FenwickTree) Update(i, v int) {
	//内部では1-indexedなのでここでインクリメントする
	//i++
	for i < fen.n {
		//fen.eval(fen.nodes[i], v)
		fen.nodes[i] = fen.nodes[i] + v
		i += i & -i
	}
}

//i(0-indexed)の値を取得する
func (fen *FenwickTree) Query(i int) int {
	//i++
	res := 0
	for i > 0 {
		//fen.eval(fen.nodes[i], res)
		res = fen.nodes[i] + res
		i -= i & -i
	}
	return res
}

func (fen *FenwickTree) Sum(l, r int) int {
	return fen.Query(r) - fen.Query(l)
}

func solveHonestly(n int, a, b []int, q int, t, x, y []int) []int {
	type card struct {
		i, a, b int
	}
	var s int
	var deck []card
	for i := 0; i < n; i++ {
		deck = append(deck, card{i, a[i], b[i]})
		s += b[i]
	}
	sort.Slice(deck, func(i, j int) bool {
		return deck[i].a > deck[j].a
	})
	var ans []int
	for i := 0; i < q; i++ {
		switch t[i] {
		case 1:
			idx := x[i] - 1
			for j := range deck {
				if deck[j].i == idx {
					deck[j].a = y[i]
				}
			}
			sort.Slice(deck, func(i, j int) bool {
				return deck[i].a > deck[j].a
			})
		case 2:
			idx := x[i] - 1
			var diff int
			for j := range deck {
				if deck[j].i == idx {
					diff = y[i] - deck[j].b
					deck[j].b = y[i]

				}
			}
			s += diff
		case 3:
			if x[i] > s {
				ans = append(ans, -1)
			} else {
				sd := x[i]
				var v int
				for j := range deck {
					v += deck[j].a * Min(sd, deck[j].b)
					sd -= deck[j].b
					if sd <= 0 {
						break
					}
				}
				ans = append(ans, v)
			}
		}
		//fmt.Println(deck)
	}
	return ans
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
