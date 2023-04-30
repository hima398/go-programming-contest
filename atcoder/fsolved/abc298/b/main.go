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

	n := nextInt()
	var a [][]int
	for i := 0; i < n; i++ {
		a = append(a, nextIntSlice(n))
	}
	var b [][]int
	for i := 0; i < n; i++ {
		b = append(b, nextIntSlice(n))
	}
	for k := 0; k < 4; k++ {
		t := make([][]int, n)
		for i := 0; i < n; i++ {
			t[i] = make([]int, n)
		}
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				t[i][j] = a[n-j-1][i]
			}
		}
		ok := true
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if t[i][j] == 1 {
					ok = ok && b[i][j] == 1
				}
			}
		}
		if ok {
			PrintString("Yes")
			return
		}
		a = t
	}
	PrintString("No")
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
