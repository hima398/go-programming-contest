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

	d, n := nextInt(), nextInt()
	var l, r, h []int
	for i := 0; i < n; i++ {
		l = append(l, nextInt())
		r = append(r, nextInt())
		h = append(h, nextInt())
	}
	ds := make([]int, d+1)
	for i := range ds {
		ds[i] = 24
	}
	for i := 0; i < n; i++ {
		for j := l[i]; j <= r[i]; j++ {
			ds[j] = Min(ds[j], h[i])
		}
	}
	var ans int
	for i := 1; i <= d; i++ {
		ans += ds[i]
	}
	PrintInt(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
