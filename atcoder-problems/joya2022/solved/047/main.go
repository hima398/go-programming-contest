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

	x, y, z := nextInt(), nextInt(), nextInt()
	if y < 0 {
		x *= -1
		y *= -1
		z *= -1
	}
	var ans int
	if x < y {
		ans = Abs(x)
	} else {
		if y < z {
			ans = -1
		} else {
			ans = Abs(z) + Abs(x-z)
		}
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

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
