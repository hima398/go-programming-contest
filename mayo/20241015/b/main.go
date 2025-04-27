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

	x1, y1, x2, y2 := nextInt(), nextInt(), nextInt(), nextInt()
	dx, dy := x2-x1, y2-y1
	x3, y3 := x2-dy, y2+dx
	x4, y4 := x3-dx, y3-dy

	fmt.Println(x3, y3, x4, y4)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}
