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
	p := nextIntSlice(n)
	q := nextInt()
	var a, b []int
	for i := 0; i < q; i++ {
		a = append(a, nextInt())
		b = append(b, nextInt())
	}
	//人iが何番目にいるか
	m := make([]int, n+1)
	for i, pi := range p {
		m[pi] = i
	}
	for i := 0; i < q; i++ {
		if m[a[i]] < m[b[i]] {
			Print(a[i])
		} else {
			Print(b[i])
		}
	}
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
