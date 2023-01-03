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

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	a, b := nextInt(), nextInt()
	ans := solve(a, b)

	PrintFloat64(ans)
}

func solve(a, b int) float64 {
	l, r := 0.0, float64(a)/float64(b)
	f := func(x float64) float64 {
		return float64(a)/math.Sqrt(x+1) + float64(b)*x
	}
	for r-l > 2 {
		m1 := (2*l + r) / 3
		m2 := (l + 2*r) / 3
		if f(m1) > f(m2) {
			l = m1
		} else {
			r = m2
		}
	}
	ans := f(float64(int(r) + 1))
	for i := int(l) - 1; i < int(r+1); i++ {
		ans = math.Min(ans, f(float64(i)))
	}
	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func PrintFloat64(x float64) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
