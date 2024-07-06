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
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	a, m, l, r := nextInt(), nextInt(), nextInt(), nextInt()
	ans := solve(a, m, l, r)
	Print(ans)
}

func solve(a, m, l, r int) int {
	l -= a
	r -= a
	return Floor(r, m) - Floor(l-1, m)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func Print(x any) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Floor(x, y int) int {
	r := (x%y + y) % y
	return (x - r) / y
}
