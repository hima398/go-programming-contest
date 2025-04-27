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
	var u, v []int
	for i := 0; i < n-1; i++ {
		u, v = append(u, nextInt()-1), append(v, nextInt()-1)
	}

	ans := solve(n, u, v)

	Print(ans)
}

func solve(n int, u, v []int) int {
	type edge struct {
		t, d int
	}
	m := n - 1
	d := make([]int, n)
	for i := 0; i < m; i++ {
		d[u[i]]++
		d[v[i]]++
	}
	e := make([][]edge, n)
	for i := 0; i < m; i++ {
		e[u[i]] = append(e[u[i]], edge{v[i], d[v[i]]})
		e[v[i]] = append(e[v[i]], edge{u[i], d[u[i]]})
	}
	ans := n
	for d0 := 0; d0 < n; d0++ {
		//fmt.Println(d0, e[d0])
		sort.Slice(e[d0], func(i, j int) bool {
			return e[d0][i].d > e[d0][j].d
		})
		//fmt.Println(d0, e[d0])
		var x int
		for _, d1 := range e[d0] {
			x++
			y := d[d1.t] - 1
			ans = Min(ans, n-(1+x+x*y))
		}
	}
	return ans
}

func solve01(n int, u, v []int) int {
	type edge struct {
		t, d int
	}
	m := n - 1
	d := make([]int, n)
	for i := 0; i < m; i++ {
		d[u[i]]++
		d[v[i]]++
	}
	e := make([][]edge, n)
	for i := 0; i < m; i++ {
		e[u[i]] = append(e[u[i]], edge{v[i], d[v[i]]})
		e[v[i]] = append(e[v[i]], edge{u[i], d[u[i]]})
	}
	ans := n
	for d0 := 0; d0 < n; d0++ {
		//fmt.Println(d0, e[d0])
		sort.Slice(e[d0], func(i, j int) bool {
			return e[d0][i].d > e[d0][j].d
		})
		//fmt.Println(d0, e[d0])
		var x int
		for _, d1 := range e[d0] {
			x++
			y := d[d1.t] - 1
			ans = Min(ans, n-(1+x+x*y))
		}
	}
	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
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
