package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()

	solve(n)
}

func computeDist(i1, j1, i2, j2 int) int {
	return Max(Abs(i2-i1), Abs(j2-j1))
}

func solve(n int) {
	di := []int{0, 1, 0, -1}
	dj := []int{1, 0, -1, 0}
	ans := make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, n)
		for j := range ans[i] {
			ans[i][j] = -1
		}
	}
	m := n / 2
	ans[m][m] = 0

	var odd []int
	for v := 1; v < n; v += 2 {
		odd = append(odd, v)
	}

	var dfs func(i, j, p, d int)
	dfs = func(i, j, p, d int) {
		ans[i][j] = p
		for k := 0; k < 4; k++ {
			ni, nj := i+di[k], j+dj[k]

			if ni < 0 || ni >= n || nj < 0 || nj >= n {
				continue
			}

			dist := computeDist(m, m, ni, nj)
			if dist != d {
				continue
			}

			if ans[ni][nj] >= 0 {
				continue
			}
			dfs(ni, nj, p+1, d)
		}
	}
	//fmt.Println(odd)
	for k, v := range odd {
		var si, sj int
		mx := -1
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if ans[i][j] > mx {
					si, sj = i, j
					mx = ans[i][j]
				}
			}
		}
		dfs(si-1, sj, v*v, k+1)
	}
	for i := range ans {
		for _, v := range ans[i] {
			if v == 0 {
				fmt.Printf("%s ", "T")
			} else {
				fmt.Printf("%d ", v)
			}
		}
		fmt.Println()
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
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
