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

	n, k := nextInt(), nextInt()
	a := nextIntSlice(n)
	ans := solve(n, k, a)
	PrintInt(ans)
}

func solve(n, k int, a []int) int {
	min := a[0]
	for _, v := range a {
		min = Min(min, v)
	}
	check := func(x int) bool {
		offset := 0
		cnt := 0
		for i, v := range a {
			diff := x - (v + offset)
			//すでにa_iは目標x以上
			if diff <= 0 {
				continue
			}
			c := Ceil(diff, i+1)
			cnt += c
			offset += c * (i + 1)
		}
		return cnt <= k
	}
	ok, ng := min, int(1e15)
	for ng-ok > 1 {
		mid := (ok + ng) / 2
		if check(mid) {
			ok = mid
		} else {
			ng = mid
		}
	}
	return ok
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

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Ceil(x, y int) int {
	return (x + y - 1) / y
}
