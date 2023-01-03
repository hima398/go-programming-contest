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
	a := nextIntSlice(n)
	ans := solve(n, m, a)
	PrintInt(ans)
}

func solve(n, m int, a []int) int {
	var b int
	var s int
	for i := 0; i < m; i++ {
		b += (i + 1) * a[i]
		s += a[i]
	}
	ans := b
	for i := m; i < n; i++ {
		b -= s
		b += m * a[i]
		s += a[i]
		s -= a[i-m]
		ans = Max(ans, b)
	}
	return ans
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

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
