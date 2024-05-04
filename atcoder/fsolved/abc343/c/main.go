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

	n := nextInt()

	ans := solve(n)

	Print(ans)
}

func IsPalindrome(x int) bool {
	s := strconv.Itoa(x)

	for i := 0; i < len(s)/2; i++ {
		j := len(s) - 1 - i
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func solve(n int) int {
	var ans int
	for x := 1; x*x*x <= n; x++ {
		k := x * x * x
		if IsPalindrome(k) {
			ans = k
		}
	}
	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func Print(x any) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
