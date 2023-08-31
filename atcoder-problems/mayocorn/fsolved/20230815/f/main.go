package main

import (
	"bufio"
	"fmt"
	"math"
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

	n, q := nextInt(), nextInt()
	a := nextIntSlice(n)
	var l, r []int
	for i := 0; i < q; i++ {
		l = append(l, nextInt()-1)
		r = append(r, nextInt()-1)
	}

	ans := solve(n, q, a, l, r)

	PrintVertically(ans)
}

func solve(n, q int, a, l, r []int) []int {
	//d := n / (int(math.Sqrt(float64(q)) + 1))
	d := n / (Sqrt(q) + 1)
	d++

	type query struct {
		i, l, r int
	}
	var qs []query
	for i := 0; i < q; i++ {
		qs = append(qs, query{i, l[i], r[i] + 1})
	}

	sort.Slice(qs, func(i, j int) bool {
		li, lj := qs[i].l/d, qs[j].l/d
		if li == lj {
			if li%2 == 0 {
				return qs[i].r < qs[j].r
			} else {
				return qs[i].r > qs[j].r
			}
		}
		return li < lj
	})

	cnt := make([]int, 2*int(1e5)+1)
	// xC3を求める
	combination := func(x int) int {
		return x * (x - 1) * (x - 2) / 6
	}

	var cur int
	add := func(idx int) {
		cur -= combination(cnt[a[idx]])
		cnt[a[idx]]++
		cur += combination(cnt[a[idx]])
	}
	del := func(idx int) {
		cur -= combination(cnt[a[idx]])
		cnt[a[idx]]--
		cur += combination(cnt[a[idx]])
	}

	ans := make([]int, q)
	var curL, curR int
	for _, v := range qs {
		for curR < v.r {
			add(curR)
			curR++
		}
		for curL > v.l {
			curL--
			add(curL)
		}
		for curR > v.r {
			curR--
			del(curR)
		}
		for curL < v.l {
			del(curL)
			curL++
		}
		ans[v.i] = cur
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

func Sqrt(x int) int {
	x2 := int(math.Sqrt(float64(x))) - 1
	for (x2+1)*(x2+1) <= x {
		x2++
	}
	return x2
}
