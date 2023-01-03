package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func computeDist2(x1, y1, x2, y2 int) int {
	xx := x2 - x1
	yy := y2 - y1
	return xx*xx + yy*yy
}
func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, k := nextInt(), nextInt()
	a := nextIntSlice(k)
	var x, y []int
	for i := 0; i < n; i++ {
		x = append(x, nextInt())
		y = append(y, nextInt())
	}
	m := make(map[int]struct{})
	for _, ai := range a {
		m[ai-1] = struct{}{}
	}
	ans := 0
	for i := 0; i < n; i++ {
		if _, found := m[i]; found {
			continue
		}
		dist := 1 << 60
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			if _, found := m[j]; found {
				dist = Min(dist, computeDist2(x[i], y[i], x[j], y[j]))
			}
		}
		ans = Max(ans, dist)
	}
	PrintFloat64(math.Sqrt(float64(ans)))
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
