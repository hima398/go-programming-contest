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
	PrintInt(ans)
}

func solve(s string) int {
	n := len(s)
	m := make(map[int]int)
	var v int
	m[0] = 1
	for i := 0; i < n; i++ {
		v = v ^ (1 << int(s[i]-'0'))
		m[v]++
	}
	//for k, v := range m {
	//	fmt.Printf("%010b, %d\n", k, v)
	//}

	var ans int
	for _, v := range m {
		ans += v * (v - 1) / 2
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
