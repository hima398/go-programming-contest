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
	a := make([][]int, n)
	for i := 0; i < n-1; i++ {
		a[i] = nextIntSlice(n - i - 1)
	}
	//fmt.Println(math.Pow(3, 10))
	ans := solve(n, a)
	PrintInt(ans)
}

func solve(n int, a [][]int) int {
	score := make([][]int, n)
	for i := 0; i < n; i++ {
		score[i] = make([]int, n)
	}
	for i := 0; i < n-1; i++ {
		for j, ai := range a[i] {
			u, v := i, i+j+1
			score[u][v] = ai
			score[v][u] = ai
		}
	}
	//for _, si := range score {
	//	PrintHorizonaly(si)
	//}
	computeScore := func(pattern []int) int {
		groups := make(map[int][]int)
		for i, g := range pattern {
			groups[g] = append(groups[g], i)
		}
		res := 0
		//fmt.Println(groups)
		for _, members := range groups {
			for i := 0; i < len(members)-1; i++ {
				for j := i + 1; j < len(members); j++ {
					res += score[members[i]][members[j]]
				}
			}
		}
		return res
	}
	ans := -int(1e18)
	var dfs func(k int, pattern []int)
	dfs = func(k int, pattern []int) {
		if len(pattern) == n {
			//fmt.Println(pattern)
			ans = Max(ans, computeScore(pattern))
			return
		}
		for i := 1; i <= k; i++ {
			next := make([]int, len(pattern))
			copy(next, pattern)
			next = append(next, i)
			dfs(k, next)
		}
	}
	for i := 1; i <= 3; i++ {
		dfs(i, []int{})
	}

	//dfs(1, []int{})
	//computeScore([]int{1, 1, 1, 1, 1, 1})
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
