package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func solve(n, q int, a, b, p, x []int) []int {
	// 0-index化は省略

	edge := make([][]int, n+1)
	for i := range a {
		edge[a[i]] = append(edge[a[i]], b[i])
		edge[b[i]] = append(edge[b[i]], a[i])
	}
	size := make([]int, n+1)
	var dfs func(cur, par int) int
	dfs = func(cur, par int) int {
		for _, next := range edge[cur] {
			if next == par {
				continue
			}
			size[cur] += dfs(next, cur)
		}
		size[cur]++
		return size[cur]
	}
	dfs(1, 0)
	sum := 0
	var ans []int
	for i := 0; i < q; i++ {
		sum += size[p[i]] * x[i]
		ans = append(ans, sum)
	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, q := nextInt(), nextInt()
	var a, b []int
	for i := 0; i < n-1; i++ {
		a = append(a, nextInt())
		b = append(b, nextInt())
	}
	var p, x []int
	for i := 0; i < q; i++ {
		p = append(p, nextInt())
		x = append(x, nextInt())
	}
	ans := solve(n, q, a, b, p, x)
	PrintVertically(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func PrintVertically(x []int) {
	defer out.Flush()
	for _, v := range x {
		fmt.Fprintln(out, v)
	}
}
