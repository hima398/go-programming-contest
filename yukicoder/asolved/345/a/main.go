package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func solveHonestly(a, b int) int {
	ans := 0
	for x := 0; x <= a; x++ {
		for y := 0; y <= b; y++ {
			mn := Min(x, Min(y, x^y))
			fmt.Println(x, y, x^y, mn)
			ans = Max(ans, mn)
		}
	}
	return ans
}

func solve(a, b int) int {
	if a > b {
		a, b = b, a
	}
	// 以降、a <= b
	p := 1
	for p<<1 <= b {
		p <<= 1
	}

	return Min(a, p-1)
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	a, b := nextInt(), nextInt()
	//ans := solveHonestly(a, b)
	ans := solve(a, b)
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

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
