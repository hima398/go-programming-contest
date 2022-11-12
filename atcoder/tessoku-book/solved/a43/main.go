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

	n, l := nextInt(), nextInt()
	var a []int
	var b []string
	for i := 0; i < n; i++ {
		a = append(a, nextInt())
		b = append(b, nextString())
	}
	ans := solve(n, l, a, b)
	PrintInt(ans)
}

func solve(n, l int, a []int, b []string) int {
	var ans int
	for i := 0; i < n; i++ {
		switch b[i] {
		case "E":
			ans = Max(ans, l-a[i])
		case "W":
			ans = Max(ans, a[i])
		}
	}
	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextString() string {
	sc.Scan()
	return sc.Text()
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
