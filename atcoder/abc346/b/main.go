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

	const t = "wbwbwwbwbwbw"

	w, b := nextInt(), nextInt()

	s := strings.Repeat(t, Ceil((w+b), len(t))+1)
	for i := 0; i < len(t); i++ {
		ss := s[i : i+w+b]
		//Print(ss)
		var cw, cb int
		for _, ssi := range ss {
			if ssi == 'w' {
				cw++
			} else if ssi == 'b' {
				cb++
			}
		}
		if cw == w && cb == b {
			Print("Yes")
			return
		}
	}
	Print("No")
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

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Ceil(x, y int) int {
	return (x + y - 1) / y
}
