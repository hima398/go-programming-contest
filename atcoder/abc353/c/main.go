package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	const m = int(1e8)
	//int(1e8)で割ったあまりでない場合
	var s int
	for _, ai := range a {
		s += (n - 1) * ai
	}

	sort.Ints(a)
	r := n
	var w int
	for i := 0; i < n; i++ {
		r = Max(r, i+1)
		for r-1 > i && a[r-1]+a[i] >= m {
			r--
		}
		w += n - r
	}
	ans := s
	ans -= w * m
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
