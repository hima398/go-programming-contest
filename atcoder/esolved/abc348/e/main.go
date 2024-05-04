package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
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
	for i := 0; i < n-1; i++ {
		a = append(a, nextInt()-1)
		b = append(b, nextInt()-1)
	}
	c := nextIntSlice(n)
	ans := solve(n, a, b, c)
	Print(ans)
}

func solve(n int, a, b, c []int) int {
	e := make([][]int, n)
	m := n - 1
	for i := 0; i < m; i++ {
		e[a[i]] = append(e[a[i]], b[i])
		e[b[i]] = append(e[b[i]], a[i])
	}

	//問題中のf(x)の値
	f := make([]int, n)
	//頂点0(0-indexed)を根にした木において、
	//頂点iを根にする部分木の中にあるciの合計
	sc := make([]int, n)

	var dfs func(cur, par, d int) int
	dfs = func(cur, par, d int) int {
		for _, next := range e[cur] {
			if next == par {
				continue
			}
			sc[cur] += dfs(next, cur, d+1)
		}
		sc[cur] += c[cur]
		f[0] += d * c[cur]
		return sc[cur]
	}
	dfs(0, -1, 0)

	//fmt.Println("c = ", c)
	//fmt.Println("sc = ", sc)

	var dfs2 func(cur, par int)
	dfs2 = func(cur, par int) {
		//f[0]は計算済みなので親が0以上の場合下記の計算をする
		if par >= 0 {
			f[cur] = f[par] - 2*sc[cur] + sc[0]
		}
		for _, next := range e[cur] {
			if next == par {
				continue
			}
			dfs2(next, cur)
		}
	}
	dfs2(0, -1)

	//fmt.Println("f(x) = ", f)
	ans := math.MaxInt
	for _, v := range f {
		ans = Min(ans, v)
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

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
