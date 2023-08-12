package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	s := strings.Split(nextString(), "")
	n := len(s)
	var ans int
	for i := n - 1; i >= 0; i-- {
		s = s[:i]
		if len(s)%2 == 1 {
			continue
		}
		ok := true
		for i := 0; i < len(s)/2; i++ {
			ok = ok && s[i] == s[len(s)/2+i]
		}
		if ok {
			ans = Max(ans, len(s))
		}
	}
	PrintInt(ans)
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
