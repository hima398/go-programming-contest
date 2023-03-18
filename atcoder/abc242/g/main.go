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

	n := nextInt()
	a := nextIntSlice(n)
	q := nextInt()
	var l, r []int
	for i := 0; i < q; i++ {
		l = append(l, nextInt()-1)
		r = append(r, nextInt()-1)
	}
	ans := solve(n, a, q, l, r)
	PrintVertically(ans)
}

// solve with Mo's algorithm
func solve(n int, a []int, q int, l, r []int) []int {
	d := n/Sqrt(q+1) + 1

	type query struct {
		i, l, r int
	}
	var qs []query
	//[l, r)(0-indexed)のクエリにする
	for i := 0; i < q; i++ {
		qs = append(qs, query{i, l[i], r[i] + 1})
	}
	sort.Slice(qs, func(i, j int) bool {
		if qs[i].r/d == qs[j].r/d {
			return qs[i].l < qs[j].l
		}
		return qs[i].r < qs[j].r
	})
	//maxA := n
	var cur int
	cnt := make([]int, n+1)
	//x人から2ペアを作る最大数
	f := func(x int) int {
		return x / 2
	}
	add := func(idx int) {
		cur -= f(cnt[a[idx]])
		cnt[a[idx]]++
		cur += f(cnt[a[idx]])
	}
	del := func(idx int) {
		cur -= f(cnt[a[idx]])
		cnt[a[idx]]--
		cur += f(cnt[a[idx]])
	}

	var curL, curR int
	ans := make([]int, q)
	for k := 0; k < q; k++ {
		nl, nr := qs[k].l, qs[k].r
		for curR < nr {
			add(curR)
			curR++
		}
		for nl < curL {
			curL--
			add(curL)
		}
		for nr < curR {
			curR--
			del(curR)
		}
		for curL < nl {
			del(curL)
			curL++
		}
		ans[qs[k].i] = cur
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

func PrintVertically(x []int) {
	defer out.Flush()
	for _, v := range x {
		fmt.Fprintln(out, v)
	}
}

func Sqrt(x int) int {
	x2 := int(math.Sqrt(float64(x))) - 1
	for (x2+1)*(x2+1) <= x {
		x2++
	}
	return x2
}
