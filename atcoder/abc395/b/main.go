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

	n := nextInt()
	ans := make([][]string, n)
	for i := range ans {
		ans[i] = make([]string, n)
		for j := range ans[i] {
			ans[i][j] = "?"
		}
	}
	for i := 0; i < n; i++ {
		j := n - i
		for ii := i; ii < j; ii++ {
			for jj := i; jj < j; jj++ {
				if i%2 == 0 {
					ans[ii][jj] = "#"
				} else {
					ans[ii][jj] = "."
				}
			}
		}
	}
	for _, row := range ans {
		Print(strings.Join(row, ""))
	}
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
