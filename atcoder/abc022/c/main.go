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
	var u, v, l []int
	for i := 0; i < m; i++ {
		u = append(u, nextInt()-1)
		v = append(v, nextInt()-1)
		l = append(l, nextInt())
	}
	ans := solve(n, m, u, v, l)
	PrintInt(ans)
}

func solve(n, m int, u, v, l []int) int {
	const INF = 1 << 60
	dist := make([][]int, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			dist[i][j] = INF
		}
	}
	for i := 0; i < m; i++ {
		if u[i] == 0 || v[i] == 0 {
			continue
		}
		dist[u[i]][v[i]] = l[i]
		dist[v[i]][u[i]] = l[i]
	}
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if i == j || j == k || k == i {
					continue
				}
				dist[i][j] = Min(dist[i][j], dist[i][k]+dist[j][k])
			}
		}
	}
	for i := 0; i < m; i++ {
		if u[i] == 0 || v[i] == 0 {
			dist[u[i]][v[i]] = l[i]
			dist[v[i]][u[i]] = l[i]
		}
	}
	ans := INF
	for i := 1; i < n; i++ {
		for j := i + 1; j < n; j++ {
			ans = Min(ans, dist[0][i]+dist[i][j]+dist[j][0])
		}
	}
	if ans == INF {
		return -1
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

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
