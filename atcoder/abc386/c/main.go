package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	k := nextInt()
	s := nextString()
	t := nextString()

	if solve(k, s, t) {
		Print("Yes")
	} else {
		Print("No")
	}
}

func solve(k int, s, t string) bool {
	if len(s) == len(t) {
		n := len(s)
		var cnt int
		for i := 0; i < n; i++ {
			if s[i] != t[i] {
				cnt++
			}
		}
		return cnt <= k
	} else {
		if len(s) > len(t) {
			s, t = t, s
		}
		if len(s)+1 != len(t) {
			return false
		}
		n := len(t)
		var cnt int
		for i := 0; i < n; i++ {
			if cnt >= 2 {
				return false
			}
			if i-cnt >= len(s) {
				return true
			}
			if s[i-cnt] != t[i] {
				cnt++
			}
		}
		return true
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func Print(x any) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
