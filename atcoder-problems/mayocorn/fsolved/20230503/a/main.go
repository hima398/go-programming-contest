package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	a, b := nextInt()
	if b >= 500 {
		a++
	}
	PrintInt(a)
}

func nextInt() (int, int) {
	sc.Scan()
	s := sc.Text()
	ss := strings.Split(s, ".")
	a, _ := strconv.Atoi(ss[0])
	b, _ := strconv.Atoi(ss[1])
	return int(a), int(b)
}
func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
