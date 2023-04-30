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

	n, l, r := nextInt(), nextInt(), nextInt()
	a := nextIntSlice(n)

	firstCanWin := solve(n, l, r, a)

	if firstCanWin {
		PrintString("First")
	} else {
		PrintString("Second")
	}
}

func solve(n, l, r int, a []int) bool {
	var s int
	for _, ai := range a {
		grundy := (ai % (l + r)) / l

		s ^= grundy
	}

	//grundy数のXORが0であれば後手必勝
	//それ以外先手必勝
	return s != 0
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func nextIntSlice(n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = nextInt()
	}
	return s
}

func PrintString(x string) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
