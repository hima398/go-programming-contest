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

	s := nextString()
	m := make(map[rune]int)
	for _, si := range s {
		m[si]++
	}
	ans := ""
	vote := 0
	for k, v := range m {
		if vote < v {
			ans = string(k)
			vote = v
		}
	}
	PrintString(ans)
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func PrintString(x string) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
