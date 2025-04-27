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

	s, t := nextString(), nextString()
	if s == t {
		Print(0)
		return
	}
	if len(s) > len(t) {
		s, t = t, s
	}
	for i := range s {
		if s[i] != t[i] {
			Print(i + 1)
			return
		}
	}
	Print(len(s) + 1)
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func Print(x any) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
