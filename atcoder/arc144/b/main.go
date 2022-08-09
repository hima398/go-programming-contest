package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func solve(n, a, b int, as []int) int {
	ok, ng := as[0], as[0]
	for _, ai := range as {
		ok = Min(ok, ai)
		ng = Max(ng, ai)
	}
	ng++
	check := func(x int) bool {
		na, nb := 0, 0
		for _, ai := range as {
			if ai < x {
				na += Ceil(x-ai, a)
			} else {
				nb += (ai - x) / b
			}
		}
		return na <= nb
	}
	for ng-ok > 1 {
		mid := (ng + ok) / 2
		if check(mid) {
			ok = mid
		} else {
			ng = mid
		}
	}
	return ok
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, a, b := nextInt(), nextInt(), nextInt()
	as := nextIntSlice(n)
	ans := solve(n, a, b, as)
	PrintInt(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
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

func Floor(x, y int) int {
	return x / y
}

func Ceil(x, y int) int {
	return (x + y - 1) / y
}
