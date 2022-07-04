package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func solve(n, x, y int, a []int) string {
	c := make([]int, n)
	for i := range a {
		c[i] = a[i] % (x + y)
	}
	// c[i] (c[i] < x + y)を最後に取るのが最善策とする
	canNotTake := true
	for _, ci := range c {
		canNotTake = canNotTake && x > ci
	}
	if canNotTake {
		return "Second"
	}
	if x <= y {
		return "First"
	}
	// x > y
	winSecond := false
	for _, ci := range c {
		//x個取れない山にy個取れる山があると、後手に有利
		winSecond = winSecond || x > ci && y <= ci
	}
	if winSecond {
		return "Second"
	} else {
		return "First"
	}
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, x, y := nextInt(), nextInt(), nextInt()
	a := nextIntSlice(n)
	ans := solve(n, x, y, a)
	PrintString(ans)
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

func PrintString(x string) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
