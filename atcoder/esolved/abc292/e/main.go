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
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, m := nextInt(), nextInt()
	var u, v []int
	for i := 0; i < m; i++ {
		u = append(u, nextInt()-1)
		v = append(v, nextInt()-1)
	}
	ans := solve(n, m, u, v)
	PrintInt(ans)
}

func solve(n, m int, u, v []int) int {
	e := make([][]int, n)
	for i := 0; i < m; i++ {
		e[u[i]] = append(e[u[i]], v[i])
	}
	bfs := func(x int) int {
		dist := make([]int, n)
		for i := range dist {
			dist[i] = -1
		}
		var q []int
		q = append(q, x)
		dist[x] = 0
		for len(q) > 0 {
			cur := q[0]
			q = q[1:]
			for _, next := range e[cur] {
				if dist[next] >= 0 {
					continue
				}
				q = append(q, next)
				dist[next] = dist[cur] + 1
			}
		}
		var res int
		for _, d := range dist {
			if d >= 2 {
				res++
			}
		}
		return res
	}
	var ans int
	for i := 0; i < n; i++ {
		ans += bfs(i)
	}
	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
