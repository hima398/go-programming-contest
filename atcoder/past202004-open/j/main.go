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
	ans := solve(s)
	PrintString(ans)
}

func solve(s string) string {
	rev := func(s string) string {
		res := ""
		for i := len(s) - 1; i >= 0; i-- {
			res += string(s[i])
		}
		return res
	}
	stack := []string{""}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, "")
		} else if s[i] == ')' {
			t := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			t = t + rev(t)
			stack[len(stack)-1] += t
		} else {
			stack[len(stack)-1] += string(s[i])
		}
	}
	ans := stack[0]
	return ans
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func PrintString(x string) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
