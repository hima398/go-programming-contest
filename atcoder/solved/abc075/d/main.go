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

	n, k := nextInt(), nextInt()
	var x, y []int
	for i := 0; i < n; i++ {
		x = append(x, nextInt())
		y = append(y, nextInt())
	}
	ans := solve(n, k, x, y)
	PrintInt(ans)
}

func solve(n, k int, x, y []int) int {

	ans := 4*int(1e18) + 1
	isIncluded := func(x1, y1, x2, y2, x, y int) bool {
		xn, xx := Min(x1, x2), Max(x1, x2)
		yn, yx := Min(y1, y2), Max(y1, y2)
		return xn <= x && x <= xx && yn <= y && y <= yx
	}
	computeArea := func(x1, y1, x2, y2 int) int {
		return Abs(x2-x1) * Abs(y2-y1)
	}
	for _, x1 := range x {
		for _, x2 := range x {
			for _, y1 := range y {
				for _, y2 := range y {
					if x1 >= x2 || y1 >= y2 {
						continue
					}
					var s int
					for l := 0; l < n; l++ {
						if isIncluded(x1, y1, x2, y2, x[l], y[l]) {
							s++
						}
					}
					if s >= k {
						ans = Min(ans, computeArea(x1, y1, x2, y2))
					}
				}
			}
		}
	}

	return ans
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
