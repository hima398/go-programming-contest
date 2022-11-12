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

	n, k := nextInt(), nextInt()
	var a, b []int
	for i := 0; i < n; i++ {
		a = append(a, nextInt())
		b = append(b, nextInt())
	}
	ans := solve(n, k, a, b)
	PrintInt(ans)
}

func solve(n, k int, a, b []int) int {
	const maxPower = 100
	var ans int
	for i := 0; i <= maxPower; i++ {
		for j := 0; j <= maxPower; j++ {
			s := 0
			for ii := 0; ii < n; ii++ {
				if i <= a[ii] && a[ii] <= i+k && j <= b[ii] && b[ii] <= j+k {
					s++
				}
			}
			ans = Max(ans, s)
		}
	}
	return ans
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

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
