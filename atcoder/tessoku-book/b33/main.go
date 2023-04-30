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

	n, h, w := nextInt(), nextInt(), nextInt()
	var a, b []int
	for i := 0; i < n; i++ {
		a = append(a, nextInt()-1)
		b = append(b, nextInt()-1)
	}

	firstCanWin := solve(n, h, w, a, b)

	if firstCanWin {
		PrintString("First")
	} else {
		PrintString("Second")
	}
}

func solve(n, h, w int, a, b []int) bool {
	var nim int
	for i := 0; i < n; i++ {
		nim ^= a[i]
		nim ^= b[i]
	}

	return !(nim == 0)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func PrintString(x string) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
