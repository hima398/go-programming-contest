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
	a := nextIntSlice(m)
	var x [][]int
	for i := 0; i < n; i++ {
		x = append(x, nextIntSlice(m))
	}
	actual := make([]int, m)
	for i := range x {
		for j := range x[i] {
			actual[j] += x[i][j]
		}
	}
	for i := 0; i < m; i++ {
		if actual[i] < a[i] {
			Print("No")
			return
		}
	}
	Print("Yes")
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

func Print(x any) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
