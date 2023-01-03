package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	s := nextString()
	var ss []string
	r := 1
	for len(s) > 0 {
		if r == len(s) || 'A' <= s[r] && s[r] <= 'Z' {
			ss = append(ss, s[:r+1])
			s = s[r+1:]
			r = 0
		}
		r++
	}
	//fmt.Println(ss)
	sort.Slice(ss, func(i, j int) bool {

		return strings.ToLower(ss[i]) < strings.ToLower(ss[j])
	})
	ans := strings.Join(ss, "")
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
