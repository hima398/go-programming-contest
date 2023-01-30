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
	t, k := make([]int, n), make([]int, n)
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		t[i], k[i] = nextInt(), nextInt()
		a[i] = nextIntSlice(k[i])
		for j := range a[i] {
			a[i][j]--
		}
	}
	ans := solve(n, t, k, a)
	PrintInt(ans)
}

func solve(n int, t, k []int, a [][]int) int {
	e := make([][]int, n)
	for i := range t {
		for _, v := range a[i] {
			e[i] = append(e[i], v)
		}
	}
	memo := make([]int, n)
	for i := range memo {
		memo[i] = -1
	}
	var dfs func(cur int) int
	dfs = func(cur int) int {
		if k[cur] == 0 {
			memo[cur] = t[cur]
			return memo[cur]
		}
		if memo[cur] >= 0 {
			return memo[cur]
		}
		s := 0
		for i := k[cur] - 1; i >= 0; i-- {
			prev := a[cur][i]
			if memo[prev] >= 0 {
				continue
			}
			s += dfs(prev)
		}
		memo[cur] = s + t[cur]
		return memo[cur]
	}
	return dfs(n - 1)
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

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
