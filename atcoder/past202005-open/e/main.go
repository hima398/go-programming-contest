package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func solve(n, m, q int, u, v, c, t, x, y []int) (ans []int) {
	for i := 0; i < m; i++ {
		u[i]--
		v[i]--
	}
	for i := 0; i < q; i++ {
		x[i]--
	}
	e := make([][]int, n)
	for i := 0; i < m; i++ {
		e[u[i]] = append(e[u[i]], v[i])
		e[v[i]] = append(e[v[i]], u[i])
	}
	for k := 0; k < q; k++ {
		ans = append(ans, c[x[k]])
		switch t[k] {
		case 1:
			for _, next := range e[x[k]] {
				c[next] = c[x[k]]
			}
		case 2:
			c[x[k]] = y[k]
		}
	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, m, q := nextInt(), nextInt(), nextInt()
	var u, v []int
	for i := 0; i < m; i++ {
		u = append(u, nextInt())
		v = append(v, nextInt())
	}
	c := nextIntSlice(n)
	var t, x, y []int
	for i := 0; i < q; i++ {
		ti := nextInt()
		t = append(t, ti)
		switch ti {
		case 1:
			x = append(x, nextInt())
			y = append(y, 0)
		case 2:
			x = append(x, nextInt())
			y = append(y, nextInt())
		}
	}

	ans := solve(n, m, q, u, v, c, t, x, y)

	PrintSlice(ans)
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

func PrintSlice(a []int) {
	defer out.Flush()
	for _, v := range a {
		fmt.Fprintln(out, v)
	}
}
