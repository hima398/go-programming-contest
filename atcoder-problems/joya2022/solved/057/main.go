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

	s, t := nextString(), nextString()
	ans := solve(s, t)
	PrintString(ans)
}

func solve(s, t string) string {
	type subString struct {
		s   byte
		cnt int
	}
	f := func(s string) []subString {
		n := len(s)
		res := []subString{{s[0], 1}}
		prev := s[0]
		for i := 1; i < n; i++ {
			if prev == s[i] {
				res[len(res)-1].cnt++
			} else {
				res = append(res, subString{s[i], 1})
			}
			prev = s[i]
		}
		return res
	}
	ss := f(s)
	tt := f(t)
	if len(ss) != len(tt) {
		return "No"
	}
	for i := range ss {
		if ss[i].s != tt[i].s {
			return "No"
		}
		if ss[i].cnt > tt[i].cnt {
			return "No"
		}
		if ss[i].cnt == 1 && tt[i].cnt > 1 {
			return "No"
		}
	}
	return "Yes"
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func PrintString(x string) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
