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

	n, a, b := nextInt(), nextInt(), nextInt()

	memo := make(map[int]bool)
	var canWin func(n int) bool
	canWin = func(n int) bool {
		if _, found := memo[n]; found {
			return memo[n]
		}
		p1, p2 := n-a, n-b
		if p2 < 0 {
			if p1 < 0 {
				memo[n] = false
				return memo[n]
			} else {
				if canWin(p1) {
					memo[n] = false
					return memo[n]
				} else {
					memo[n] = true
					return memo[n]
				}
			}
		}
		if canWin(p1) && canWin(p2) {
			memo[n] = false
			return memo[n]
		} else {
			memo[n] = true
			return memo[n]
		}
	}

	if canWin(n) {
		PrintString("First")
	} else {
		PrintString("Second")
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func PrintString(x string) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
