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
	var b [][]int
	for i := 0; i < n; i++ {
		b = append(b, nextIntSlice(m))
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m-1; j++ {
			if b[i][j+1]-b[i][j] != 1 || (b[i][j+1]-1)%7 <= (b[i][j]-1)%7 {
				PrintString("No")
				return
			}
		}
	}
	for j := 0; j < m; j++ {
		for i := 0; i < n-1; i++ {
			if b[i+1][j]-b[i][j] != 7 {
				PrintString("No")
				return
			}
		}
	}
	PrintString("Yes")
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
