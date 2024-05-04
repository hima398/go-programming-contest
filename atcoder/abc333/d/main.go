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
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	var u, v []int
	for i := 0; i < n-1; i++ {
		u = append(u, nextInt()-1)
		v = append(v, nextInt()-1)
	}
	ans := solve(n, u, v)
	Print(ans)
}

func solve(n int, u, v []int) int {
	e := make([][]int, n)
	for i := 0; i < n-1; i++ {
		e[u[i]] = append(e[u[i]], v[i])
		e[v[i]] = append(e[v[i]], u[i])
	}
	var p []int
	var dfs func(cur, par int) int
	dfs = func(cur, par int) int {
		res := 1
		for _, next := range e[cur] {
			if next == par {
				continue
			}
			res += dfs(next, cur)
		}
		return res
	}
	for _, c := range e[0] {
		p = append(p, dfs(c, 0))
	}
	sort.Ints(p)
	var ans int
	for _, v := range p {
		ans += v
	}
	ans = ans - p[len(p)-1] + 1
	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func Print(x any) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
