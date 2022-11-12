package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	s := nextString()
	ans := solve(s)
	PrintString(ans)
}

func solve(s string) string {
	r := regexp.MustCompile("^A?KIHA?BA?RA?$")
	if r.Match([]byte(s)) {
		return "YES"
	} else {
		return "NO"
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
