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
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, m := nextInt(), nextInt()
	a := nextIntSlice(n)

	ans := solve(n, m, a)
	PrintVertically(ans)
}

func solve(n, m int, a []int) []int {
	ans := make([]int, m)

	f := func(i, x int) int {
		return (i+1)*x + a[i]
	}

	ps := make([][]int, n+1)
	for i := 0; i < n; i++ {
		x := sort.Search(m+1, func(x int) bool {
			return f(i, x) >= 0
		})
		for j := x; j <= m && f(i, j) <= n; j++ {
			ps[f(i, j)] = append(ps[f(i, j)], j)
		}
	}
	//fmt.Println(ps)
	for y := 0; y <= n; y++ {
		for _, x := range ps[y] {
			if x > 0 && ans[x-1] == y {
				ans[x-1]++
			}
		}
	}

	return ans
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

func PrintVertically(x []int) {
	defer out.Flush()
	for _, v := range x {
		fmt.Fprintln(out, v)
	}
}
