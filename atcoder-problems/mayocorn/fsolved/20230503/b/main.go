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

	h, _ := nextInt(), nextInt()
	var c []string
	for i := 0; i < h; i++ {
		c = append(c, nextString())
	}
	ans := make([]string, 2*h)
	for i := 0; i < h; i++ {
		ans[2*i] = c[i]
		ans[2*i+1] = c[i]
	}
	for _, v := range ans {
		PrintString(v)
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func PrintString(x string) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
