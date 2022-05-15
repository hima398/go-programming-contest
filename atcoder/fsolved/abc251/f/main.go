package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

type edge struct {
	s, t int
}

func constructT1(n, m int, e [][]int) []edge {

	v := make([]bool, n)

	var res []edge
	var f func(par, i int)
	f = func(par, i int) {
		v[i] = true
		for _, next := range e[i] {
			if v[next] {
				continue
			}
			f(i, next)
		}
		if par < 0 {
			return
		}
		res = append(res, edge{par + 1, i + 1})
	}
	f(-1, 0)
	return res
}

func constructT2(n, m int, e [][]int) []edge {
	type item struct {
		par, i int
	}
	var q []item
	v := make([]bool, n)
	q = append(q, item{-1, 0})
	v[0] = true

	var res []edge
	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		for _, next := range e[p.i] {
			if v[next] {
				continue
			}
			v[next] = true
			res = append(res, edge{p.i + 1, next + 1})
			q = append(q, item{p.i, next})
		}
	}
	return res
}

func solve(n, m int, u, v []int) (t1, t2 []edge) {
	for i := 0; i < m; i++ {
		u[i]--
		v[i]--
	}
	e := make([][]int, n)
	for i := 0; i < m; i++ {
		e[u[i]] = append(e[u[i]], v[i])
		e[v[i]] = append(e[v[i]], u[i])
	}

	t1 = constructT1(n, m, e)
	t2 = constructT2(n, m, e)

	return t1, t2
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, m := nextInt(), nextInt()
	var u, v []int
	for i := 0; i < m; i++ {
		u = append(u, nextInt())
		v = append(v, nextInt())
	}
	t1, t2 := solve(n, m, u, v)

	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for _, v := range t1 {
		fmt.Fprintln(out, v.s, v.t)
	}
	for _, v := range t2 {
		fmt.Fprintln(out, v.s, v.t)
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
