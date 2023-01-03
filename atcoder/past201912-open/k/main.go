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
	p := nextIntSlice(n)
	q := nextInt()
	var a, b []int
	for i := 0; i < q; i++ {
		a = append(a, nextInt()-1)
		b = append(b, nextInt()-1)
	}
	ans := solve(n, p, q, a, b)
	PrintVertically(ans)
}

var depth []int
var parents [][30]int

func Lca(u, v int) int {
	if depth[u] > depth[v] {
		u, v = v, u
	}
	for i := 0; i < 21; i++ {
		if ((depth[v]-depth[u])>>i)&1 == 1 {
			v = parents[v][i]
		}
	}
	if u == v {
		return u
	}
	for i := 20; i >= 0; i-- {
		if parents[u][i] != parents[v][i] {
			u = parents[u][i]
			v = parents[v][i]
		}
	}
	return parents[u][0]
}

func solve(n int, p []int, q int, a, b []int) []string {
	e := make([][]int, 2*n)
	//fmt.Println(p)
	root := -1
	for i, pi := range p {
		pi--
		if pi < 0 {
			root = i
			continue
		}
		e[i] = append(e[i], pi)
		e[pi] = append(e[pi], i)
	}
	//fmt.Println(e)
	depth = make([]int, 2*n)
	parents = make([][30]int, 2*n)
	var Dfs func(i, p, v int)
	Dfs = func(i, p, v int) {
		depth[i] = v
		parents[i][0] = p
		for _, chi := range e[i] {
			if chi == p {
				continue
			}
			Dfs(chi, i, v+1)
		}
	}
	Dfs(root, -1, 0)
	//fmt.Println(parents)
	for i := 0; i < 20; i++ {
		for j := 0; j < n; j++ {
			if parents[j][i] < 0 {
				parents[j][i+1] = -1
			} else {
				parents[j][i+1] = parents[parents[j][i]][i]
			}
		}
	}
	//fmt.Println(parents)

	var ans []string
	for k := 0; k < q; k++ {
		par := Lca(a[k], b[k])
		//fmt.Println(par)
		if par == b[k] {
			ans = append(ans, "Yes")
		} else {
			ans = append(ans, "No")
		}
	}
	return ans
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

func PrintVertically(x []string) {
	defer out.Flush()
	for _, v := range x {
		fmt.Fprintln(out, v)
	}
}
