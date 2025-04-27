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
	type node struct {
		s   rune
		cnt int
	}
	var ns []node
	for _, si := range s {
		if len(ns) == 0 || ns[len(ns)-1].s != si {
			ns = append(ns, node{si, 1})
		} else {
			ns[len(ns)-1].cnt++
		}
	}
	var ans int
	for _, v := range ns {
		if v.s == '0' {
			ans += Ceil(v.cnt, 2)
		} else {
			ans += v.cnt
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

func Ceil(x, y int) int {
	return (x + y - 1) / y
}
