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

	m := nextInt()
	var a, c []int
	for i := 0; i < m; i++ {
		a = append(a, nextInt())
		c = append(c, nextInt())
	}
	var ans float64
	var s float64
	for i := 0; i < m; i++ {
		ans += float64(a[i] * c[i])
		s += float64(c[i])
	}
	PrintFloat64(ans / s)
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
