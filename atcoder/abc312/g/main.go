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
	ans := solve(n, a, b)
	PrintInt(ans)
}

func solve(n int, a, b []int) int {
	e := make([][]int, n)
	for i := 0; i < n-1; i++ {
		e[a[i]] = append(e[a[i]], b[i])
		e[b[i]] = append(e[b[i]], a[i])
	}
	c := make([]int, n)
	//部分木の要素がいくつあるかを数えるDFS
	var dfs func(cur, par int) int
	dfs = func(cur, par int) int {
		res := 1
		for _, next := range e[cur] {
			if next == par {
				continue
			}
			res += dfs(next, cur)
		}
		c[cur] = res
		return c[cur]
	}
	dfs(0, -1)
	//fmt.Println(c)

	//要素curにおける部分木からi, jを、部分木外からkを選ぶ組み合わせの数
	nC2 := func(n int) int {
		return n * (n - 1) / 2
	}
	ans := n * (n - 1) * (n - 2) / 6 // nC3
	var s int
	var dfs2 func(cur, par int)
	dfs2 = func(cur, par int) {
		if par >= 0 {
			s += c[cur] * (n - c[cur])
		}
		for _, next := range e[cur] {
			if next == par {
				continue
			}
			dfs2(next, cur)
		}
	}
	dfs2(0, -1)
	ans -= s - nC2(n)
	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
