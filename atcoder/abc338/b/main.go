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
	m := make([]int, 26)
	for _, si := range s {
		m[int(si-'a')]++
	}
	mx := 0
	ans := ""
	for c := 0; c < 26; c++ {
		if mx < m[c] {
			mx = m[c]
			ans = string(c + 'a')
		}
	}
	Print(ans)
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func Print(x any) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
