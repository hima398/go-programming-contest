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
	PrintVertically(ans)
}

func solve(n int, a []int) []int {
	m := make(map[int]struct{})
	for _, ai := range a {
		m[ai] = struct{}{}
	}
	var b []int
	for k := range m {
		b = append(b, k)
	}
	sort.Ints(b)
	s := make(map[int]int)
	for i := len(b) - 1; i >= 0; i-- {
		s[b[i]] = len(b) - 1 - i
	}
	ans := make([]int, n)
	for _, ai := range a {
		ans[s[ai]]++
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

func PrintVertically(x []int) {
	defer out.Flush()
	for _, v := range x {
		fmt.Fprintln(out, v)
	}
}
