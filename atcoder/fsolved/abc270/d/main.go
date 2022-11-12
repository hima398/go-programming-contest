package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func solve(n, k int, a []int) int {
	m := make(map[int]struct{})
	for _, ai := range a {
		m[ai] = struct{}{}
	}
	memo := make([]int, n+1)
	visited := make([]bool, n+1)
	var f func(x int) int
	f = func(x int) int {
		if visited[x] {
			return memo[x]
		}

		mx := 0
		for i := 0; i < k && a[i] <= x; i++ {
			mx = Max(mx, a[i]+(x-a[i])-f(x-a[i]))
		}
		memo[x] = mx
		visited[x] = true
		return memo[x]
	}
	return f(n)
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, k := nextInt(), nextInt()
	a := nextIntSlice(k)
	ans := solve(n, k, a)
	PrintInt(ans)
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

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
