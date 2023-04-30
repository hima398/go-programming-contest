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

	n, k := nextInt(), nextInt()
	var a [][]int
	for i := 0; i < n; i++ {
		a = append(a, nextIntSlice(n))
	}
	q := nextInt()
	var s, t []int
	for i := 0; i < q; i++ {
		s = append(s, nextInt()-1)
		t = append(t, nextInt()-1)
	}

	ans := solve(n, k, a, q, s, t)

	PrintVertically(ans)
}

func solve(n, k int, a [][]int, q int, s, t []int) []int {
	const INF = 1 << 60
	dist := make([][]int, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == j {
				dist[i][j] = 0
			} else {
				if a[i][j] == 1 {
					dist[i][j] = 1
				} else {
					dist[i][j] = INF
				}
			}
		}
	}
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if i == j || j == k || k == i {
					continue
				}
				dist[i][j] = Min(dist[i][j], dist[i][k]+dist[k][j])
			}
		}
	}
	var ans []int
	for l := 0; l < q; l++ {
		//i, j := s[l]%k, t[l]%k
		v := INF
		i := s[l] % n
		for j := 0; j < n; j++ {
			if a[i][j] == 1 {
				v = Min(v, dist[j][t[l]%n]+1)
			}
		}
		if v == INF {
			ans = append(ans, -1)
		} else {
			ans = append(ans, v)
		}
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

func PrintVertically(x []int) {
	defer out.Flush()
	for _, v := range x {
		fmt.Fprintln(out, v)
	}
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
