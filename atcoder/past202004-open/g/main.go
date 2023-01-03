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
	q := nextInt()
	t, x, d := make([]int, q), make([]int, q), make([]int, q)
	c := make([]string, q)
	for i := 0; i < q; i++ {
		t[i] = nextInt()
		switch t[i] {
		case 1:
			c[i], x[i] = nextString(), nextInt()
		case 2:
			d[i] = nextInt()
		}
	}
	ans := solve(q, t, c, x, d)
	PrintVertically(ans)
}

func solve(q int, t []int, c []string, x []int, d []int) []int {
	var ans []int
	type node struct {
		c string
		x int
	}
	var ns []node
	for i := 0; i < q; i++ {
		switch t[i] {
		case 1:
			ns = append(ns, node{c[i], x[i]})
		case 2:
			m := make(map[string]int)
			for len(ns) > 0 && d[i] >= ns[0].x {
				m[ns[0].c] += ns[0].x
				d[i] -= ns[0].x
				ns = ns[1:]
			}
			if len(ns) > 0 && d[i] > 0 {
				m[ns[0].c] += d[i]
				ns[0].x -= d[i]
			}
			s := 0
			for _, v := range m {
				s += v * v
			}
			ans = append(ans, s)
		}
	}
	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func PrintVertically(x []int) {
	defer out.Flush()
	for _, v := range x {
		fmt.Fprintln(out, v)
	}
}
