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

	n := nextInt()
	var t, x, y []int
	for i := 0; i < n; i++ {
		t = append(t, nextInt())
		x = append(x, nextInt())
		y = append(y, nextInt())
	}
	ok := solve(n, t, x, y)
	if ok {
		PrintString("Yes")
	} else {
		PrintString("No")
	}
}

func computeManhattanDist(x1, y1, x2, y2 int) int {
	return Abs(x2-x1) + Abs(y2-y1)
}

func solve(n int, t, x, y []int) bool {
	curT, curX, curY := 0, 0, 0
	for i := 0; i < n; i++ {
		dt := t[i] - curT
		dist := computeManhattanDist(x[i], y[i], curX, curY)
		if dt < dist {
			return false
		}
		if dt%2 != dist%2 {
			return false
		}
		curT, curX, curY = t[i], x[i], y[i]
	}
	return true
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func PrintString(x string) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
