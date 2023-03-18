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

	n, a, b, c, d := nextInt(), nextInt(), nextInt(), nextInt(), nextInt()
	ok := solve(n, a, b, c, d)
	if ok {
		PrintString("Yes")
	} else {
		PrintString("No")
	}
}

func solve(n, a, b, c, d int) bool {
	if Abs(b-c) == 1 {
		return true
	} else if Abs(b-c) == 0 {
		if (a > 0 && d > 0) && (b == 0 && c == 0) {
			return false
		}
		return true
	} else {
		return false
	}
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

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
