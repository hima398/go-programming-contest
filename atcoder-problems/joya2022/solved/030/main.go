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

	n := nextInt()
	var x, y []int
	for i := 0; i < n; i++ {
		x = append(x, nextInt())
		y = append(y, nextInt())
	}
	d := 0
	computeDist2 := func(x1, y1, x2, y2 int) int {
		xx := x2 - x1
		yy := y2 - y1
		return xx*xx + yy*yy
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			d = Max(d, computeDist2(x[i], y[i], x[j], y[j]))
		}
	}
	ans := math.Sqrt(float64(d))
	PrintFloat64(ans)
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

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
