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

	n, m := nextInt(), nextInt()
	var u, v, t []int
	for i := 0; i < m; i++ {
		u = append(u, nextInt()-1)
		v = append(v, nextInt()-1)
		t = append(t, nextInt())
	}
	q := nextInt()
	var k []int
	var b [][]int
	for i := 0; i < q; i++ {
		k = append(k, nextInt())
		b = append(b, nextIntSlice(k[i]))
		for j := range b[i] {
			b[i][j]--
		}
	}

	ans := solve(n, m, u, v, t, q, k, b)

	for _, v := range ans {
		Print(v)
	}
}

func solve(n, m int, u, v, t []int, q int, k []int, b [][]int) []int {
	const INF = 1 << 60

	dist := make([][]int, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == j {
				//dist[i][j]=0のこと
				continue
			}
			dist[i][j] = INF
		}
	}
	for i := 0; i < m; i++ {
		dist[u[i]][v[i]] = Min(dist[u[i]][v[i]], t[i])
		dist[v[i]][u[i]] = Min(dist[v[i]][u[i]], t[i])
	}
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				dist[i][j] = Min(dist[i][j], dist[i][k]+dist[k][j])
			}
		}
	}

	var ans []int
	for k := 0; k < q; k++ {
		minDist := INF
		for {
			for j := 0; j < 1<<len(b[k]); j++ {
				var totalDist int
				var cur int
				for i := range b[k] {
					if (j>>i)&1 == 0 {
						totalDist += dist[cur][u[b[k][i]]] + t[b[k][i]]
						cur = v[b[k][i]]
					} else {
						totalDist += dist[cur][v[b[k][i]]] + t[b[k][i]]
						cur = u[b[k][i]]
					}
				}
				totalDist += dist[cur][n-1]
				minDist = Min(minDist, totalDist)
			}

			if !NextPermutation(sort.IntSlice(b[k])) {
				break
			}
		}
		ans = append(ans, minDist)
	}
	return ans
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

func Print(x any) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func NextPermutation(x sort.Interface) bool {
	n := x.Len() - 1
	if n < 1 {
		return false
	}
	j := n - 1
	for ; !x.Less(j, j+1); j-- {
		if j == 0 {
			return false
		}
	}
	l := n
	for !x.Less(j, l) {
		l--
	}
	x.Swap(j, l)
	for k, l := j+1, n; k < l; {
		x.Swap(k, l)
		k++
		l--
	}
	return true
}
