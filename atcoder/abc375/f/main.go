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

	n, m, q := nextInt(), nextInt(), nextInt()
	var a, b, c []int
	for i := 0; i < m; i++ {
		a, b = append(a, nextInt()-1), append(b, nextInt()-1)
		c = append(c, nextInt())
	}
	t := make([]int, q)
	idx := make([]int, q)
	x, y := make([]int, q), make([]int, q)
	for i := 0; i < q; i++ {
		t[i] = nextInt()
		switch t[i] {
		case 1:
			idx[i] = nextInt() - 1
		case 2:
			x[i], y[i] = nextInt()-1, nextInt()-1
		}
	}

	ans := solve(n, m, q, a, b, c, t, idx, x, y)

	for _, v := range ans {
		Print(v)
	}
}

func solve(n, m, q int, a, b, c, t, idx, x, y []int) []int {
	const INF = 1 << 60

	stop := make(map[int]struct{})
	for i := 0; i < q; i++ {
		if t[i] == 1 {
			stop[idx[i]] = struct{}{}
		}
	}
	dist := make([][]int, n)
	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			if i == j {
				continue
			}
			dist[i][j] = INF
		}
	}
	for i := 0; i < m; i++ {
		if _, found := stop[i]; found {
			continue
		}
		dist[a[i]][b[i]] = c[i]
		dist[b[i]][a[i]] = c[i]
	}
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				dist[i][j] = Min(dist[i][j], dist[i][k]+dist[k][j])
			}
		}
	}

	var ans []int
	for i := q - 1; i >= 0; i-- {
		switch t[i] {
		case 1:
			dist[a[idx[i]]][b[idx[i]]] = Min(dist[a[idx[i]]][b[idx[i]]], c[idx[i]])
			dist[b[idx[i]]][a[idx[i]]] = Min(dist[b[idx[i]]][a[idx[i]]], c[idx[i]])
			for _, k := range []int{a[idx[i]], b[idx[i]]} {
				for i := 0; i < n; i++ {
					for j := 0; j < n; j++ {
						dist[i][j] = Min(dist[i][j], dist[i][k]+dist[k][j])
					}
				}
			}
		case 2:
			if dist[x[i]][y[i]] == INF {
				ans = append(ans, -1)
			} else {
				ans = append(ans, dist[x[i]][y[i]])
			}
		}
	}
	for i := 0; i < len(ans)/2; i++ {
		j := len(ans) - 1 - i
		ans[i], ans[j] = ans[j], ans[i]
	}

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

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
