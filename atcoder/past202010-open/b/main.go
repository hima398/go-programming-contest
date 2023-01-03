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

	defer out.Flush()
	x, y := nextInt(), nextInt()
	if y == 0 {
		fmt.Fprintln(out, "ERROR")
		return
	}
	ans := float64(x) / float64(y)
	ans = math.Floor(100.0*ans) / 100.0
	fmt.Fprintf(out, "%.2f\n", ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
