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
	a := nextIntSlice(n)

	ans := solve(n, a)

	Print(ans)
}

func solve(n int, a []int) int {
	l := make([]int, n)
	cur := 0
	for i := 0; i < n; i++ {
		if cur < a[i] {
			cur++
		} else {
			cur = a[i]
		}
		l[i] = cur
	}
	r := make([]int, n)
	cur = 0
	for i := n - 1; i >= 0; i-- {
		if cur < a[i] {
			cur++
		} else {
			cur = a[i]
		}
		r[i] = cur
	}
	ans := 1
	for i := 0; i < n; i++ {
		ans = Max(ans, Min(l[i], r[i]))
	}
	return ans
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

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
