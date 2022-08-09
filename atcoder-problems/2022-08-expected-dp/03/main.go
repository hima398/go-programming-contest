package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func solve(n int, a []int) float64 {
	var dp [301][301][301]float64
	var visited [301][301][301]bool
	var dfs func(x, y, z int) float64
	dfs = func(x, y, z int) float64 {
		if visited[x][y][z] {
			return dp[x][y][z]
		}
		visited[x][y][z] = true
		if x == 0 && y == 0 && z == 0 {
			return 0.0
		}
		res := float64(n)
		if x > 0 {
			fx := float64(x)
			res += dfs(x-1, y, z) * fx
		}
		if y > 0 {
			fy := float64(y)
			res += dfs(x+1, y-1, z) * fy

		}
		if z > 0 {
			fz := float64(z)
			res += dfs(x, y+1, z-1) * fz
		}
		res /= float64(x + y + z)
		dp[x][y][z] = res
		//fmt.Println(res, x, y, z)
		return dp[x][y][z]
	}

	var m [4]int
	for _, v := range a {
		m[v]++
	}

	ans := dfs(m[1], m[2], m[3])
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	a := nextIntSlice(n)
	ans := solve(n, a)
	PrintFloat64(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextIntSlice(n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = nextInt()
	}
	return s
}

func PrintFloat64(x float64) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
