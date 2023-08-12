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

	a := nextIntSlice(7)
	computeDist := func(a, b, c, x int) int {
		t := x / (a + c)
		r := x % (a + c)
		return (t*a + Min(a, r)) * b
	}
	takahashi := computeDist(a[0], a[1], a[2], a[6])
	aoki := computeDist(a[3], a[4], a[5], a[6])
	if takahashi > aoki {
		PrintString("Takahashi")
	} else if takahashi < aoki {
		PrintString("Aoki")
	} else {
		PrintString("Draw")
	}
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

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
