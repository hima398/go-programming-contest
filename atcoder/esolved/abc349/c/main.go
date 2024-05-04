package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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
	t := nextString()

	//ok := solve(s, t)
	ok := solveByRegexp(s, t)

	if ok {
		Print("Yes")
	} else {
		Print("No")
	}
}

func solveByRegexp(s, t string) bool {
	s += "X"
	s = strings.ToUpper(s)

	pattern := ".*" + strings.Join(strings.Split(t, ""), ".*") + ".*"
	reg, _ := regexp.Compile(pattern)

	return reg.Match([]byte(s))
}

func solve(s, t string) bool {
	s += "x"
	s = strings.ToUpper(s)
	//fmt.Println(s)
	i, j := 0, 0
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			j++
		}
		i++
	}
	return j == len(t)
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func Print(x any) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
