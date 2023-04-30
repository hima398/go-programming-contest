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

	n, m, d := nextInt(), nextInt(), nextInt()
	a := nextIntSlice(m)
	ans := solve(n, m, d, a)
	PrintVertically(ans)
}

func solve(n, m, d int, a []int) []int {
	for i := range a {
		a[i]--
	}
	var p []int
	for i := 0; i < n; i++ {
		p = append(p, i)
	}

	for _, ai := range a {
		p[ai], p[ai+1] = p[ai+1], p[ai]
	}
	t := make([]int, n)
	for i, v := range p {
		t[v] = i
	}

	memo := make([]map[int]int, n)
	var f func(x, d int) int
	f = func(x, d int) int {
		if v, found := memo[x][d]; found {
			return v
		}
		if memo[x] == nil {
			memo[x] = make(map[int]int)
		}
		if d == 1 {
			return t[x]
		}
		if d%2 == 1 {
			x = t[x]
			d--
		}
		memo[x][d] = f(f(x, d/2), d/2)
		return memo[x][d]
	}
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = f(i, d)
		ans[i]++
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
