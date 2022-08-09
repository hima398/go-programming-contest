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

	n, m := nextInt(), nextInt()
	var a, b, p, q []int
	for i := 0; i < n; i++ {
		a = append(a, nextInt())
		p = append(p, nextInt())
	}
	for i := 0; i < m; i++ {
		b = append(b, nextInt())
		q = append(q, nextInt())
	}
	var ex, ey float64
	for i := 0; i < n; i++ {
		ex += float64(a[i] * p[i])
	}
	for i := 0; i < m; i++ {
		ey += float64(b[i] * q[i])
	}
	ans := ex * ey / 10000.0
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
