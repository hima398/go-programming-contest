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
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, m := nextInt(), nextInt()
	a := make([]string, n)
	for i := 0; i < n; i++ {
		a[i] = nextString()
	}
	ans := solve(n, m, a)
	PrintInt(ans)
}

func computeDist(i1, j1, i2, j2 int) int {
	return Abs(i2-i1) + Abs(j2-j1)
}
func solve(n, m int, a []string) int {
	type cell struct {
		i, j int
	}
	cells := make([][]cell, 11)
	const INF = 1 << 60
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
		for j := 0; j < m; j++ {
			dp[i][j] = INF
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if a[i][j] == 'S' {
				cells[0] = append(cells[0], cell{i, j})
				dp[i][j] = 0
			} else if a[i][j] == 'G' {
				cells[10] = append(cells[10], cell{i, j})
			} else {
				cells[int(a[i][j]-'0')] = append(cells[int(a[i][j]-'0')], cell{i, j})
			}
		}
	}
	for k := 1; k <= 10; k++ {
		for _, c1 := range cells[k-1] {
			for _, c2 := range cells[k] {
				dp[c2.i][c2.j] = Min(dp[c2.i][c2.j], dp[c1.i][c1.j]+computeDist(c1.i, c1.j, c2.i, c2.j))
			}
		}
	}
	ans := dp[cells[10][0].i][cells[10][0].j]
	if ans == INF {
		ans = -1
	}
	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func PrintInt(x int) {
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
