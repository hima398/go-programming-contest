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
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	c := nextIntSlice(n)
	// 1-indexed
	c = append([]int{0}, c...)

	q := nextInt()
	t := make([]int, q)
	x := make([]int, q)
	a := make([]int, q)
	for i := 0; i < q; i++ {
		t[i] = nextInt()
		switch t[i] {
		case 1:
			x[i], a[i] = nextInt(), nextInt()
		case 2, 3:
			a[i] = nextInt()
		}
	}
	ans := solve(n, c, q, t, x, a)
	PrintInt(ans)
}

func solve(n int, c []int, q int, t, x, a []int) int {
	//奇数のカード、偶数のカードがそれぞれ何枚売れたか
	soldOdd, soldEven := 0, 0
	fenOdd, fenEven := NewFenwickTree(n), NewFenwickTree(n)
	for i := 1; i <= n; i++ {
		if i%2 == 1 {
			fenOdd.Update(i+1, c[i])
		} else {
			fenEven.Update(i+1, c[i])
		}
	}
	var ans int
	for i := 0; i < q; i++ {
		switch t[i] {
		case 1:
			cc := c[x[i]]
			if x[i]%2 == 1 {
				cc -= soldOdd
			} else {
				cc -= soldEven
			}
			if cc >= a[i] {
				ans += a[i]
				c[x[i]] -= a[i]
				if x[i]%2 == 1 {
					fenOdd.Update(x[i], c[x[i]])
				} else {
					fenEven.Update(x[i], c[x[i]])
				}
			}
		case 2:
			c := fenOdd.Query(n)
			if c-soldOdd >= a[i] {
				ans += a[i] * Ceil(n, 2)
				soldOdd += a[i]
			}
		case 3:
			cOdd, cEven := fenOdd.Query(n), fenEven.Query(n)
			if Min(cOdd-soldOdd, cEven-soldEven) >= a[i] {
				ans += a[i] * n
				soldOdd += a[i]
				soldEven += a[i]
			}
		}
		//fmt.Println(i, ans)
		//fmt.Println(soldOdd, soldEven)
		//fmt.Println(fenOdd.nodes)
		//fmt.Println(fenEven.nodes)
	}
	return ans
}

const INF = 1 << 60

type FenwickTree struct {
	n     int
	nodes []int
	//eval  func(x1, x2 int) int
}

//
func New(n int) *FenwickTree {
	return NewFenwickTree(n)
}

func NewFenwickTree(n int) *FenwickTree {
	fen := new(FenwickTree)
	// 1-indexed
	fen.n = n + 1
	fen.nodes = make([]int, fen.n)
	for i := 0; i < fen.n; i++ {
		fen.nodes[i] = INF
	}
	//bt.eval = f
	return fen
}

//i(0-indexed)をvに更新する
func (fen *FenwickTree) Update(i, v int) {
	for i < fen.n {
		fen.nodes[i] = Min(fen.nodes[i], v) //fen.eval(fen.nodes[i], v)
		i += i & -i
	}
}

//i(0-indexed)の値を取得する
func (fen *FenwickTree) Query(i int) int {
	res := INF
	for i > 0 {
		res = Min(res, fen.nodes[i]) //fen.eval(fen.nodes[i], res)
		i -= i & -i
	}
	return res
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

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Ceil(x, y int) int {
	return (x + y - 1) / y
}
