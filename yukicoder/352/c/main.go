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

func solve(n int, a, b []int) float64 {
	type Problem struct {
		consider, implment int
	}
	var ps []Problem
	for i := 0; i < n; i++ {
		if i == 0 || a[i-1]*b[i] <= a[i]*b[i-1] {
			ps = append(ps, Problem{a[i], b[i]})
		} else {
			ps[i-1].consider += a[i]
			ps[i-1].implment += b[i]
		}
	}
	var ans float64
	for _, p := range ps {
		if p.consider <= p.implment {
			ans += 2.0 * math.Sqrt(float64(p.consider)/float64(p.implment))
		} else {
			ans += 1.0
		}

	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	a, b := nextIntSlice(n), nextIntSlice(n)

	ans := solve(n, a, b)
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
