package main

import (
	"bufio"
	"fmt"
	"math"
	"math/cmplx"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func digToRad(dig float64) float64 {
	return math.Pi * dig / 180.0
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	a, b, d := nextFloat64(), nextFloat64(), nextFloat64()
	c1 := complex(a, b)
	c2 := cmplx.Rect(1, digToRad(d))

	ans := c1 * c2
	fmt.Println(real(ans), imag(ans))
}

func nextFloat64() float64 {
	sc.Scan()
	f, _ := strconv.ParseFloat(sc.Text(), 64)
	return f
}
