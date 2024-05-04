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

	s, t := nextString(), nextString()

	ok := solve(s, t)
	if ok {
		PrintString("Yes")
	} else {
		PrintString("No")
	}
}

func solve(s, t string) bool {
	const u = "atcoder"
	ms := make(map[rune]int)
	for _, si := range s {
		ms[si]++
	}
	mt := make(map[rune]int)
	for _, ti := range t {
		mt[ti]++
	}

	for _, ui := range u {
		max := Max(ms[ui], mt[ui])
		ms['@'] -= max - ms[ui]
		ms[ui] = max
		mt['@'] -= max - mt[ui]
		mt[ui] = max
	}
	if ms['@'] < 0 || mt['@'] < 0 {
		return false
	}

	for b := 'a'; b <= 'z'; b++ {
		if ms[b] != mt[b] {
			return false
		}
	}
	return true
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func PrintString(x string) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
