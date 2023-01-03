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

	n := nextInt()
	a := make([][]int, n)
	a[0] = make([]int, 1)
	a[0][0] = 1
	for i := 1; i < n; i++ {
		a[i] = make([]int, i+1)
		for j := range a[i-1] {
			a[i][j] += a[i-1][j]
			a[i][j+1] += a[i-1][j]
		}
	}
	PrintVertically(a)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func PrintHorizonaly(x []int) {
	defer out.Flush()
	fmt.Fprintf(out, "%d", x[0])
	for i := 1; i < len(x); i++ {
		fmt.Fprintf(out, " %d", x[i])
	}
	fmt.Fprintln(out)
}

func PrintVertically(x [][]int) {
	defer out.Flush()
	for _, v := range x {
		//fmt.Fprintln(out, v)
		PrintHorizonaly(v)
	}
}
