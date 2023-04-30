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

	n := nextInt()
	a := nextIntSlice(n)
	ans := solveCommentary(n, a)
	PrintInt(ans)
}

func solveCommentary(n int, a []int) int {
	cnt := make([]int, n+1)
	var cur int
	//nC2
	c2 := func(n int) int {
		return n * (n - 1) / 2
	}
	add := func(x, c int) {
		cur -= c2(cnt[x])
		cnt[x] += c
		cur += c2(cnt[x])
	}
	for _, ai := range a {
		add(ai, 1)
	}
	var ans int
	l, r := 0, n-1
	for l < r {
		ans += c2(r-l+1) - cur
		add(a[l], -1)
		add(a[r], -1)
		l++
		r--
	}
	return ans
}

func solve(n int, a []int) int {
	var ans int
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if a[i] != a[j] {
				ans++
			}
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

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
