package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func MulMatrix(n int, a, b [][]int, p int) [][]int {
	ret := make([][]int, n)
	for i := 0; i < n; i++ {
		ret[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				ret[i][j] += a[i][k] * b[k][j]
				ret[i][j] %= p
			}
		}
	}
	return ret
}

func PowMatrix(n int, a [][]int, x, p int) [][]int {
	ret := make([][]int, n)
	for i := 0; i < n; i++ {
		ret[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		ret[i][i] = 1
	}
	for x > 0 {
		if x%2 == 1 {
			ret = MulMatrix(n, ret, a, p)
		}
		x >>= 1
		a = MulMatrix(n, a, a, p)
	}
	return ret
}
func solve(n, m, lt int, s, t []int) int {
	const p = 998244353

	w := make([][]int, n)
	for i := 0; i < n; i++ {
		w[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		w[s[i]][t[i]] = 1
		w[t[i]][s[i]] = 1
	}
	w = PowMatrix(n, w, lt, p)

	return w[0][0]
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, m, lt := nextInt(), nextInt(), nextInt()
	var s, t []int
	for i := 0; i < m; i++ {
		s = append(s, nextInt())
		t = append(t, nextInt())
	}

	ans := solve(n, m, lt, s, t)

	PrintInt(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
