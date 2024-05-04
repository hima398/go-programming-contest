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
	m := nextInt()
	b := nextIntSlice(m)
	l := nextInt()
	c := nextIntSlice(l)
	q := nextInt()
	x := nextIntSlice(q)

	oks := solve(n, a, m, b, l, c, q, x)

	for _, ok := range oks {
		if ok {
			Print("Yes")
		} else {
			Print("No")
		}
	}
}

func solve(n int, a []int, m int, b []int, l int, c []int, q int, x []int) []bool {
	y := make(map[int]struct{})
	for _, ai := range a {
		for _, bi := range b {
			for _, ci := range c {
				y[ai+bi+ci] = struct{}{}
			}
		}
	}

	var ans []bool
	for _, xi := range x {
		if _, found := y[xi]; found {
			ans = append(ans, true)
		} else {
			ans = append(ans, false)
		}
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
