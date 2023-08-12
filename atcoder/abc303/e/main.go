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
	PrintHorizonaly(ans)
}

func solve(n int, u, v []int) []int {
	e := make([][]int, n)
	m := n - 1
	d := make([]int, n)
	for i := 0; i < m; i++ {
		d[u[i]]++
		d[v[i]]++
		e[u[i]] = append(e[u[i]], v[i])
		e[v[i]] = append(e[v[i]], u[i])
	}
	var ans []int
	var dfs func(cur, par int) int
	dfs = func(cur, par int) int {
		var numChildren int
		for _, next := range e[cur] {
			if next == par {
				continue
			}
			numChildren += dfs(next, cur)
		}
		if d[cur] == 2 && numChildren > 1 {
			ans = append(ans, numChildren)
			return 0
		}
		return numChildren + 1
	}
	root := dfs(0, -1)
	if root-1 >= 2 {
		ans = append(ans, root-1)
	}

	sort.Ints(ans)
	return ans
}
func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func PrintHorizonaly(x []int) {
	defer out.Flush()
	fmt.Fprintf(out, "%d", x[0])
	for i := 1; i < len(x); i++ {
		fmt.Fprintf(out, " %d", x[i])
	}
	fmt.Fprintln(out)
}
