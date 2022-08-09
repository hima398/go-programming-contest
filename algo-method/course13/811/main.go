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

	p := nextInt()
	ans := make([]float64, 10)
	q := float64(100-p) / 100.0
	for i := 0; i < 10; i++ {
		ans[i] = float64(p) * math.Pow(q, float64(i)) / 100.0
	}
	PrintVertically(ans)
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

func PrintVertically(x []float64) {
	defer out.Flush()
	for _, v := range x {
		fmt.Fprintln(out, v)
	}
}
