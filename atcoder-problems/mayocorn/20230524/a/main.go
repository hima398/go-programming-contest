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

	s := []byte(nextString())
	n := len(s)
	isPalindrome := func(s []byte) bool {
		ok := true
		n := len(s)
		for i := 0; i < n/2; i++ {
			ok = ok && s[i] == s[n-1-i]
		}
		return ok
	}
	//fmt.Println(string(s[:n/2]), string(s[n/2+1:]))
	if isPalindrome(s) && isPalindrome(s[:n/2]) && isPalindrome(s[n/2+1:]) {
		PrintString("Yes")
	} else {
		PrintString("No")
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
