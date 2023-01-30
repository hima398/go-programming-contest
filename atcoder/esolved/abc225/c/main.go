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

	n, m := nextInt(), nextInt()
	b := make([][]int, n)
	for i := 0; i < n; i++ {
		b[i] = nextIntSlice(m)
	}
	ans := solve(n, m, b)
	PrintString(ans)
}

func solve(n, m int, b [][]int) string {
	c := make([][]int, n)
	for i := 0; i < n; i++ {
		c[i] = make([]int, m)
		for j := 0; j < m; j++ {
			c[i][j] = (b[i][j] + 6) % 7
		}
	}
	for i := 0; i < n; i++ {
		for j := 1; j < m; j++ {
			if b[i][j]-b[i][j-1] != 1 || c[i][j]-c[i][j-1] != 1 {
				return "No"
			}
		}
	}
	for j := 0; j < m; j++ {
		for i := 1; i < n; i++ {
			if b[i][j]-b[i-1][j] != 7 || c[i][j] != c[i-1][j] {
				return "No"
			}
		}
	}
	return "Yes"
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
