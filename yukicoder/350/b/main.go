package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func solve(n int, a []int) float64 {

	v := 1.0
	for _, ai := range a {
		v *= (1000.0 - float64(ai)) / 1000.0
	}
	d := 1000.0 * (1.0 - v)
	return d
}

func solveHonestly(n int, a []int) float64 {
	d := 1000.0
	for _, ai := range a {
		d *= (1000.0 - float64(ai)) / 1000.0
	}
	return 1000.0 - d

}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	a := nextIntSlice(n - 1)

	ans := solveHonestly(n, a)
	//ans := solve(n, a)
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
