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

	n, q := nextInt(), nextInt()
	t, x, y := make([]int, q), make([]int, q), make([]int, q)
	for i := 0; i < q; i++ {
		t[i] = nextInt()
		switch t[i] {
		case 1:
			x[i] = nextInt() - 1
			y[i] = nextInt()
		case 3:
			x[i] = nextInt() - 1
		}
	}
	ans := solve(n, q, t, x, y)
	PrintVertically(ans)
}

func solve(n, q int, t, x, y []int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = i + 1
	}
	isReversed := 0
	var ans []int
	for i := 0; i < q; i++ {
		switch t[i] {
		case 1:
			var idx int
			if isReversed == 1 {
				idx = n - 1 - x[i]
			} else {
				idx = x[i]
			}
			a[idx] = y[i]
		case 2:
			isReversed ^= 1
		case 3:
			var idx int
			if isReversed == 1 {
				idx = n - 1 - x[i]
			} else {
				idx = x[i]
			}
			ans = append(ans, a[idx])
		}
	}
	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func PrintVertically(x []int) {
	defer out.Flush()
	for _, v := range x {
		fmt.Fprintln(out, v)
	}
}
