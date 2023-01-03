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

	var x, y []int
	for i := 0; i < 4; i++ {
		x = append(x, nextInt())
		y = append(y, nextInt())
	}
	ans := solve(x, y)
	PrintString(ans)
}

func computeAngle(x0, y0, x1, y1, x2, y2 int) float64 {
	ax, ay := float64(x1-x0), float64(y1-y0)
	bx, by := float64(x2-x0), float64(y2-y0)
	c := (ax*bx + ay*by) / (math.Sqrt(ax*ax+ay*ay) * math.Sqrt(bx*bx+by*by))
	return math.Acos(c)
}

func solve(x, y []int) string {
	var sa []float64
	for i := 0; i < 4; i++ {
		i1, i2 := (i+3)%4, (i+1)%4
		a := computeAngle(x[i], y[i], x[i1], y[i1], x[i2], y[i2])
		sa = append(sa, a)
	}
	for i := 0; i < 4; i++ {
		var s1, s2 float64
		for j := 0; j < 4; j++ {
			if i == j {
				s1 += sa[j]
				s2 += 2.0*math.Pi - sa[j]
			} else {
				s1 += sa[j]
				s2 += sa[j]
			}
		}
		if math.Abs(2.0*math.Pi-s1) > math.Abs(2.0*math.Pi-s2) {
			return "No"
		}
	}
	return "Yes"
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func PrintString(x string) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
