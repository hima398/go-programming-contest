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

	m, n, ln := nextInt(), nextInt(), nextInt()
	ans := solve(m, n, ln)
	PrintInt(ans)
}

func solve(m, n, ln int) int {
	ans := ln
	r := 0
	for ln >= m {
		d := ln / m
		r = ln % m
		ln = n * d
		//販売する
		ans += ln

		//あまりを足す
		ln += r
	}
	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
