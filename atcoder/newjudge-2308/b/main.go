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

	x, k := nextInt(), nextInt()
	w := make([]int, 16)
	w[0] = 1
	for i := 0; i < 15; i++ {
		w[i+1] = 10 * w[i]
	}
	for i := 0; i < k; i++ {
		v := x / w[i]
		m := v % 10
		if m >= 5 {
			v = (v - m) + 10
		} else {
			v = (v - m)
		}
		x = v * w[i]
	}
	Print(x)
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
