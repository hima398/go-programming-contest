package main

import (
	"bufio"
	"fmt"
	"math/cmplx"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	a, b := nextFloat64(), nextFloat64()
	c1 := complex(a, b)
	r, t := cmplx.Polar(c1)
	c2 := cmplx.Rect(r/cmplx.Abs(c1), t)
	fmt.Println(real(c2), imag(c2))
}

func nextFloat64() float64 {
	sc.Scan()
	f, _ := strconv.ParseFloat(sc.Text(), 64)
	return f
}
