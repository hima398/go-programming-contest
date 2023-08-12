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
	var p, c []int
	var f [][]int
	for i := 0; i < n; i++ {
		p = append(p, nextInt())
		c = append(c, nextInt())
		f = append(f, nextIntSlice(c[i]))
	}
	f2 := make([][]bool, n)
	for i := range f2 {
		f2[i] = make([]bool, m)
	}
	for i := range c {
		for _, fj := range f[i] {
			f2[i][fj-1] = true
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			ok := true
			for k := 0; k < m; k++ {
				if !f2[i][k] {
					continue
				}
				ok = ok && f2[i][k] && f2[j][k]
			}
			if ok && p[i] >= p[j] {
				if p[i] > p[j] || len(f[i]) < len(f[j]) {
					PrintString("Yes")
					return
				}
			}
		}
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
