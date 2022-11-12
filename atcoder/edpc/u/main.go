package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func solve(n int, a [][]int) int {
	//集合Sに属するウサギをグループ分けした最高得点
	//memo := make(map[int]int)
	memo := make([]int, 1<<n)
	visited := make([]bool, 1<<n)
	var dfs func(s int) int
	dfs = func(s int) int {
		if visited[s] {
			return memo[s]
		}
		visited[s] = true
		var score int
		for i := 0; i < n-1; i++ {
			for j := i + 1; j < n; j++ {
				if s>>i&1 == 1 && s>>j&1 == 1 {
					score += a[i][j]
				}
			}
		}
		for t := s; t >= 0; t-- {
			t &= s
			if t == s || t == 0 {
				continue
			}
			score = Max(score, dfs(t)+dfs(s^t))
		}
		memo[s] = score
		return memo[s]
	}

	ans := dfs(1<<n - 1)
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = nextIntSlice(n)
	}
	ans := solve(n, a)
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

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
