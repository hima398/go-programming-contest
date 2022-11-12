package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	x := nextString()
	var dfs func(x string) bool
	dfs = func(x string) bool {
		if len(x) == 0 {
			return true
		}
		switch x[len(x)-1] {
		case 'h':
			if len(x) > 1 && x[len(x)-2] == 'c' {
				return dfs(x[:len(x)-2]) && true
			} else {
				return false
			}
		case 'o', 'k', 'u':
			return dfs(x[:len(x)-1]) && true
		default:
			return false
		}
	}
	if dfs(x) {
		PrintString("YES")
	} else {
		PrintString("NO")
	}
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func PrintString(x string) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
