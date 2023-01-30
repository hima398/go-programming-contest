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

	h, b := nextFloat64(), nextFloat64()
	h /= 100.0
	ans := b * h * h
	PrintFloat64(ans)
}

func nextFloat64() float64 {
	sc.Scan()
	f, _ := strconv.ParseFloat(sc.Text(), 64)
	return f
}

func PrintFloat64(x float64) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
