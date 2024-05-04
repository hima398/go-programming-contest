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

	n, q := nextInt(), nextInt()
	r := nextIntSlice(n)
	x := nextIntSlice(q)

	ans := solve(n, q, r, x)

	PrintVertically(ans)
}

func solve(n, q int, r, x []int) []int {
	sort.Ints(r)
	s := make([]int, n+1)
	for i := range r {
		s[i+1] += s[i] + r[i]
	}
	var ans []int
	for _, xi := range x {
		idx := sort.Search(n+1, func(i int) bool {
			return s[i] > xi
		})
		ans = append(ans, idx-1)
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
