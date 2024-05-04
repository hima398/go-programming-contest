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

	ans := solve(s)

	Print(ans)
}

func solve(s string) int {
	n := len(s)
	m := make(map[rune]int)
	for _, si := range s {
		m[si]++
	}
	ans := n * (n - 1) / 2
	var hasSameString bool
	for _, v := range m {
		if v >= 2 {
			hasSameString = true
		}
		ans -= v * (v - 1) / 2
	}
	if hasSameString {
		ans++
	}
	return ans
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func Print(x any) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
