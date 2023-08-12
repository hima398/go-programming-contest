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

	n, m, d := nextInt(), nextInt(), nextInt()
	a := nextIntSlice(n)
	b := nextIntSlice(m)

	ans := solve(n, m, d, a, b)

	PrintInt(ans)
}

func solve(n, m, d int, a, b []int) int {
	sort.Ints(a)
	sort.Ints(b)
	ans := -1
	for i := 0; i < n; i++ {
		if a[i]-d < 0 {
			continue
		}
		idx := sort.Search(m, func(j int) bool {
			return b[j] >= a[i]-d
		})
		if idx == m {
			continue
		}
		if b[idx] > a[i] {
			continue
		}
		//fmt.Println(a[i], b[idx])
		ans = Max(ans, a[i]+b[idx])
	}
	for i := 0; i < m; i++ {
		if b[i]-d < 0 {
			continue
		}
		idx := sort.Search(n, func(j int) bool {
			return a[j] >= b[i]-d
		})
		if idx == n {
			continue
		}
		if a[idx] > b[i] {
			continue
		}
		ans = Max(ans, a[idx]+b[i])
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
