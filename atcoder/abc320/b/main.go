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
	n := len(s)
	isPalindrome := func(s string, l, r int) bool {
		n := r - l + 1
		ok := true
		for i := 0; i < n/2; i++ {
			ok = ok && s[l+i] == s[r-i]
		}
		return ok
	}
	ans := 1
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if isPalindrome(s, i, j) {
				ans = Max(ans, j-i+1)
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

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
