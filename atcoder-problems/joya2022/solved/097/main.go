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
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, q := nextInt(), nextInt()
	x := nextIntSlice(n)
	var a, b []int
	for i := 0; i < n-1; i++ {
		a = append(a, nextInt()-1)
		b = append(b, nextInt()-1)
	}
	var v, k []int
	for i := 0; i < q; i++ {
		v = append(v, nextInt()-1)
		k = append(k, nextInt()-1)
	}
	ans := solve(n, q, x, a, b, v, k)
	PrintVertically(ans)
}

func solve(n, q int, x, a, b, v, k []int) []int {
	e := make([][]int, n)
	for i := 0; i < n-1; i++ {
		e[a[i]] = append(e[a[i]], b[i])
		e[b[i]] = append(e[b[i]], a[i])
	}
	rank := make([][]int, n)
	var dfs func(cur, par int) []int
	dfs = func(cur, par int) []int {
		var subRank []int
		subRank = append(subRank, x[cur])
		for _, next := range e[cur] {
			if next == par {
				continue
			}
			s := dfs(next, cur)
			subRank = append(subRank, s...)
		}
		sort.Slice(subRank, func(i, j int) bool {
			return subRank[i] > subRank[j]
		})
		if len(subRank) > 20 {
			subRank = subRank[:20]
		}
		rank[cur] = subRank
		return subRank
	}
	dfs(0, -1)
	var ans []int
	for i := 0; i < q; i++ {
		ans = append(ans, rank[v[i]][k[i]])
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

func PrintVertically(x []int) {
	defer out.Flush()
	for _, v := range x {
		fmt.Fprintln(out, v)
	}
}
