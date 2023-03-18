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

	n, m := nextInt(), nextInt()
	a := nextIntSlice(m)

	ans := solve(n, m, a)
	PrintVertically(ans)
}

func solve(n, m int, a []int) []int {
	children := make([]int, n)
	ans := make([]int, m)
	for k := 0; k < m; k++ {
		idx := sort.Search(n, func(i int) bool {
			return children[i] > -a[k]
		})
		if idx == n {
			ans[k] = -1
		} else {
			ans[k] = idx + 1
			children[idx] = -a[k]
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

func PrintVertically(x []int) {
	defer out.Flush()
	for _, v := range x {
		fmt.Fprintln(out, v)
	}
}
