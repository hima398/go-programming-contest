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
	s := make([]string, n)
	for i := 0; i < n; i++ {
		s[i] = nextString()
	}
	ans := solve(n, m, s)
	PrintVertically(ans)
}

func solve(n, m int, s []string) [][]int {
	ans := make([][]int, n)
	for i := 0; i < n; i++ {
		ans[i] = make([]int, m)
	}
	di := []int{-1, -1, -1, 0, 0, 0, 1, 1, 1}
	dj := []int{-1, 0, 1, -1, 0, 1, -1, 0, 1}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			v := 0
			for k := 0; k < 9; k++ {
				ni, nj := i+di[k], j+dj[k]
				if ni < 0 || ni >= n || nj < 0 || nj >= m {
					continue
				}
				if s[ni][nj] == '#' {
					v++
				}
			}
			ans[i][j] = v
		}
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

func PrintVertically(x [][]int) {
	defer out.Flush()
	for _, y := range x {
		for _, v := range y {
			fmt.Fprintf(out, "%d", v)
		}
		fmt.Fprintln(out)
	}
}
