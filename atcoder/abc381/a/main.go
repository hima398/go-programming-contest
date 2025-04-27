package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	_ = nextInt()
	s := nextString()
	t := strings.Split(s, "/")
	//fmt.Println(t, len(t))
	if len(t) != 2 {
		Print("No")
		return
	}
	for _, ti := range t[0] {
		if ti != '1' {
			Print("No")
			return
		}
	}
	if len(t[0]) != len(t[1]) {
		Print("No")
		return
	}
	for _, ti := range t[1] {
		if ti != '2' {
			Print("No")
			return
		}
	}
	Print("Yes")
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func Print(x any) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
