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
	ans := solveCommentary(n, q, a, l, r)
	PrintVertically(ans)
}

// solve with Mo's algorithm
func solveCommentary(n, q int, a, l, r []int) []int {
	//分割幅をSqrt(q)くらいに設定
	d := n / (int(math.Sqrt(float64(q)) + 1))
	d++

	type query struct {
		i, l, r int
	}
	var qs []query
	for i := 0; i < q; i++ {
		qs = append(qs, query{i, l[i], r[i] + 1})
	}
	//クエリを並び替え
	sort.Slice(qs, func(i, j int) bool {
		if qs[i].r/d == qs[j].r/d {
			return qs[i].l < qs[j].l
		}
		return qs[i].r < qs[j].r
	})

	//Aiの出現頻度
	const maxA = 2 * int(1e5)
	var cur int
	cnt := make([]int, maxA+1)

	//探索する際に使用するメソッド
	c3 := func(x int) int {
		return x * (x - 1) * (x - 2) / 6
	}
	add := func(idx int) {
		cur -= c3(cnt[a[idx]]) //Combination(cnt[a[idx]], 3)
		cnt[a[idx]]++
		cur += c3(cnt[a[idx]]) //Combination(cnt[a[idx]], 3)
	}
	del := func(idx int) {
		cur -= c3(cnt[a[idx]]) //Combination(cnt[a[idx]], 3)
		cnt[a[idx]]--
		cur += c3(cnt[a[idx]]) //Combination(cnt[a[idx]], 3)
	}

	ans := make([]int, q)
	var curL, curR int
	//[l, r)の区間なのでインクリメントのタイミングが異なる
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

func PrintVertically(x []int) {
	defer out.Flush()
	for _, v := range x {
		fmt.Fprintln(out, v)
	}
}

func Combination(N, K int) int {
	if K == 0 {
		return 1
	}
	if K == 1 {
		return N
	}
	return Combination(N, K-1) * (N + 1 - K) / K
}
