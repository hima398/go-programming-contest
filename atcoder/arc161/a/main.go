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

	if solve(n, a) {
		Print("Yes")
	} else {
		Print("No")
	}
}

func solve(n int, a []int) bool {
	sort.Ints(a)
	var low, high []int
	for i := 0; i < Ceil(n, 2); i++ {
		low = append(low, a[i])
	}
	for i := Ceil(n, 2); i < n; i++ {
		high = append(high, a[i])
	}
	for i := 0; i < len(high); i++ {
		if low[i+1] >= high[i] {
			return false
		}
	}
	return true
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

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
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
	/*
		return x / y
	*/
	r := (x%y + y) % y
	return (x - r) / y
}

func Ceil(x, y int) int {
	return (x + y - 1) / y
}
