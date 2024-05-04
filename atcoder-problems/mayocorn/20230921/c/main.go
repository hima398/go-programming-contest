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

	r1, c1 := nextInt(), nextInt()
	r2, c2 := nextInt(), nextInt()

	ans := solve(r1, c1, r2, c2)

	Print(ans)
}

func solve(r1, c1, r2, c2 int) int {
	if r1 == r2 && c1 == c2 {
		return 0
	}
	if r1+c1 == r2+c2 || r1-c1 == r2-c2 || Abs(r1-r2)+Abs(c1-c2) <= 3 {
		return 1
	}
	if (r1+c1+r2+c2)%2 == 0 {
		return 2
	}
	if Abs(r1-r2)+Abs(c1-c2) <= 6 {
		return 2
	}
	if Abs(r1+c1-r2-c2) <= 3 || Abs(r1-c1-r2+c2) <= 3 {
		return 2
	}
	return 3
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

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
