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
		Print("Yes")
	} else {
		Print("No")
	}
}

type node struct {
	c byte
	v int
}

func runLength(s string) []node {
	n := len(s)
	res := []node{{s[0], 1}}
	for i := 1; i < n; i++ {
		m := len(res)
		if res[m-1].c == s[i] {
			res[m-1].v++
		} else {
			res = append(res, node{s[i], 1})
		}
	}
	return res
}

func solve(s, t string) bool {
	compressedS := runLength(s)
	compressedT := runLength(t)
	if len(compressedS) != len(compressedT) {
		return false
	}
	//len(compressedS)==len(compressedT)
	for i := range compressedS {
		if compressedS[i].c != compressedT[i].c {
			return false
		}
		if compressedS[i].v == 1 && compressedT[i].v > 1 {
			return false
		}
		if compressedS[i].v > compressedT[i].v {
			return false
		}
	}
	return true
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func Print(x any) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
