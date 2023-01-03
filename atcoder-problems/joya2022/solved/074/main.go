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
	ans := solve(s)
	PrintInt(ans)
}

func solve(s string) int {
	const t = "atcoder"
	n := len(s)
	p := make([]int, n)
	for i := range t {
		for j := 0; j < n; j++ {
			if s[j] == t[i] {
				p[i] = j
			}
		}
	}
	var ans int
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if p[i] > p[j] {
				p[i], p[j] = p[j], p[i]
				ans++
			}
		}
	}
	return ans
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
