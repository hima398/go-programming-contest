package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	s := nextString()
	t := nextString()

	near := func(x, y byte) bool {
		if x > y {
			x, y = y, x
		}
		return y-x == 1 || y-x == 4
	}
	if near(s[0], s[1]) == near(t[0], t[1]) {
		Print("Yes")
	} else {
		Print("No")
	}
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func Print(x any) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
