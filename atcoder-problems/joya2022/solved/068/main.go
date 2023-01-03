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

	n, x, y := nextInt(), nextInt(), nextInt()
	ans := solve(n, x, y)
	PrintInt(ans)
}

func solve(n, x, y int) int {
	const red, blue = 0, 1
	var f func(lv int, color int) int
	f = func(lv int, color int) int {
		if lv == 1 {
			if color == red {
				return 0
			} else {
				return 1
			}
		}
		switch color {
		case red:
			return f(lv-1, red) + f(lv, blue)*x
		case blue:
			return f(lv-1, red) + f(lv-1, blue)*y
		}
		return 0
	}
	ans := f(n, red)
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
