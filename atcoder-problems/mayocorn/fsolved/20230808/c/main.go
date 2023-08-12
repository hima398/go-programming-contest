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
	q := nextInt()
	var l, r, x []int
	for i := 0; i < q; i++ {
		l = append(l, nextInt()-1)
		r = append(r, nextInt()-1)
		x = append(x, nextInt())
	}
	ans := solve(n, a, q, l, r, x)
	PrintVertically(ans)
}

func solve(n int, a []int, q int, l, r, x []int) []int {
	m := make([][]int, n+1)
	for i, v := range a {
		m[v] = append(m[v], i)
	}
	var ans []int
	for i := 0; i < q; i++ {
		li := sort.Search(len(m[x[i]]), func(j int) bool {
			return l[i] <= m[x[i]][j]
		})
		ri := sort.Search(len(m[x[i]]), func(j int) bool {
			return r[i] < m[x[i]][j]
		})
		ans = append(ans, ri-li)
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
