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
	for i := 0; i < len(s)-1; i += 2 {
		if s[i] != s[i+1] {
			Print("No")
			return
		}
	}
	m := make(map[rune]int)
	for _, si := range s {
		m[si]++
	}
	for _, v := range m {
		if v == 1 || v > 2 {
			Print("No")
			return
		}
	}
	Print("Yes")
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func Print(x any) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
