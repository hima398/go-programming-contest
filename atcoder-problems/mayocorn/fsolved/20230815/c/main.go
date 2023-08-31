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

	w, h, x, y := nextInt(), nextInt(), nextInt(), nextInt()
	ans := float64(w*h) / 2.0
	var canDivide int
	if 2*x == w && 2*y == h {
		canDivide = 1
	}
	fmt.Println(ans, canDivide)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}
