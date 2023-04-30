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

	PrintInt(ans)
}

func solve(n int, a []int) int {
	type point struct {
		i, d int
	}
	for i := range a {
		a[i]--
	}
	d := make([]int, n)
	for _, ai := range a {
		d[ai]++
	}
	var q []int
	for i, di := range d {
		if di == 0 {
			q = append(q, i)
		}
	}
	rem := 0
	for len(q) > 0 {
		rem++
		x := q[0]
		q = q[1:]
		d[a[x]]--
		if d[a[x]] == 0 {
			q = append(q, a[x])
		}
	}
	ans := n - rem
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

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
