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

	s := nextString()

	ans := solve(s)

	Print(ans)
}

func solve(s string) string {
	if isPalindrome(s) {
		return s
	}
	for i := 0; i < len(s); i++ {
		t := s[i:]
		//fmt.Println(i, t)
		if isPalindrome(t) {
			return s + reverse(s[:len(s)-len(t)])
		}
	}
	return s
}

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func reverse(s string) string {
	res := strings.Split(s, "")
	n := len(res)
	for i := 0; i < n/2; i++ {
		j := n - 1 - i
		res[i], res[j] = res[j], res[i]
	}
	return strings.Join(res, "")
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func Print(x any) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
