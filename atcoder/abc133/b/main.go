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

	n, d := nextInt(), nextInt()
	x := make([][]int, n)
	for i := 0; i < n; i++ {
		x[i] = nextIntSlice(d)
	}
	var ans int
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			var dist int
			for k := 0; k < d; k++ {
				t := x[i][k] - x[j][k]
				dist += t * t
			}
			for k := 1; k*k <= dist; k++ {
				if dist%k == 0 && k == dist/k {
					ans++
				}
			}
		}
	}
	PrintInt(ans)
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

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
