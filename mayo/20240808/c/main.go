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
	var l, r []int
	for i := 0; i < q; i++ {
		l = append(l, nextInt())
		r = append(r, nextInt())
	}
	ans := solve(n, a, q, l, r)
	for _, v := range ans {
		Print(v)
	}
}

func solve(n int, a []int, q int, l, r []int) []int {
	a = append(a, int(1e9)+1)
	s := make([]int, n+1)
	for i := 0; i < n; i++ {
		if i%2 == 1 {
			s[i+1] = s[i] + a[i+1] - a[i]
		} else {
			s[i+1] = s[i]
		}
	}
	var ans []int
	for i := 0; i < q; i++ {
		idxL := sort.Search(len(a), func(j int) bool {
			return l[i] < a[j+1]
		})
		idxR := sort.Search(len(a), func(j int) bool {
			return r[i] < a[j+1]
		})
		v := s[idxR] - s[idxL]
		if idxL%2 == 1 {
			v -= Abs(l[i] - a[idxL])
		}
		if idxR%2 == 1 {
			v += Abs(r[i] - a[idxR])
		}
		ans = append(ans, v)
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
