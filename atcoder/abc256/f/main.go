package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

var p = 998244353

func solve(n, q int, a, t, x, v []int) (ans []int) {
	for i := 0; i < q; i++ {
		x[i]--
	}

	ft0, ft1, ft2 := NewFenwickTree(n), NewFenwickTree(n), NewFenwickTree(n)
	for i := 0; i < n; i++ {
		ft0.Add(i, a[i])
		ft1.Add(i, i*a[i]%p)
		ft2.Add(i, ((i*i)%p*a[i])%p)
	}

	i2 := Inv(2, p)
	for i := 0; i < q; i++ {
		switch t[i] {
		case 1:
			d := (v[i] - a[x[i]] + p) % p
			a[x[i]] = v[i]
			a0 := d
			a1 := x[i] * d % p
			a2 := (x[i] * x[i] % p) * d % p
			ft0.Add(x[i], a0)
			ft1.Add(x[i], a1)
			ft2.Add(x[i], a2)
		case 2:
			a2 := ft2.Sum(0, x[i]+1)
			a1 := -(2*x[i] + 3) * ft1.Sum(0, x[i]+1) % p
			a1 = (a1 + p) % p
			a0 := (x[i] + 1) * (x[i] + 2) % p
			a0 = a0 * ft0.Sum(0, x[i]+1) % p
			ans = append(ans, (i2*(a2+a1+a0))%p)
		}
	}
	return
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, q := nextInt(), nextInt()
	a := nextIntSlice(n)
	var t, x, v []int
	for i := 0; i < q; i++ {
		t = append(t, nextInt())
		x = append(x, nextInt())
		if t[i] == 1 {
			v = append(v, nextInt())
		} else {
			v = append(v, 0)
		}
	}
	ans := solve(n, q, a, t, x, v)
	PrintVertically(ans)
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

func PrintVertically(x []int) {
	defer out.Flush()
	for _, v := range x {
		fmt.Fprintln(out, v)
	}
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

type FenwickTree struct {
	n     int
	nodes []int
}

func NewFenwickTree(n int) *FenwickTree {
	fen := new(FenwickTree)
	fen.n = n + 1
	fen.nodes = make([]int, fen.n)
	return fen
}

func (fen *FenwickTree) Add(i, v int) {
	i++
	for i <= fen.n {
		fen.nodes[i-1] += v
		fen.nodes[i-1] %= p
		i += i & -i
	}
}

func (fen *FenwickTree) Sum(l, r int) int {
	return (fen.sum(r) - fen.sum(l) + p) % p
}

func (fen *FenwickTree) sum(i int) int {
	res := 0
	for i > 0 {
		res += fen.nodes[i-1]
		res %= p
		i -= i & -i
	}
	return res
}
