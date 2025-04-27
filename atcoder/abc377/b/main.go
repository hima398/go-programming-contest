package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	const n = 8
	var s []string
	for i := 0; i < n; i++ {
		s = append(s, nextString())
	}
	field := make([][]bool, n)
	for i := 0; i < n; i++ {
		field[i] = make([]bool, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if s[i][j] == '.' {
				continue
			}
			for ii := 0; ii < n; ii++ {
				field[ii][j] = true
			}
			for jj := 0; jj < n; jj++ {
				field[i][jj] = true
			}
		}
	}

	var ans int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if !field[i][j] {
				ans++
			}
		}
	}
	Print(ans)
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func Print(x any) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
