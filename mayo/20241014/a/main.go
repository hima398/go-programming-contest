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
	a := nextIntSlice(n)
	x := nextIntSlice(q)

	ans := solve(n, q, a, x)

	for _, v := range ans {
		Print(v)
	}
}

func solve(n, q int, a, x []int) []int {
	sort.Ints(a)

	var ans []int
	for _, xi := range x {
		idx := sort.Search(n, func(i int) bool {
			return a[i] >= xi
		})
		ans = append(ans, n-idx)
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
