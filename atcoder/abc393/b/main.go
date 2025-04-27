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

	var ans int
	for d := 1; d <= 100; d++ {
		for i := 0; i < len(s)-2*d; i++ {
			if s[i] == 'A' && s[i+d] == 'B' && s[i+2*d] == 'C' {
				ans++
			}
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
