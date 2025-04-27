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

	ans := solve(n, a)

	Print(ans)
}

func solve(n int, a []int) int {
	for i := range a {
		a[i]--
	}
	p := make([][]int, n)
	for i, ai := range a {
		if len(p[ai]) == 0 {
			p[ai] = append(p[ai], 0)
		}
		p[ai] = append(p[ai], i+1)
	}
	var ans int
	for i := 0; i < n; i++ {
		if len(p[i]) == 0 {
			continue
		}
		cnt := n * (n + 1) / 2
		p[i] = append(p[i], n+1)
		for j := 1; j < len(p[i]); j++ {
			cnt -= (p[i][j] - p[i][j-1] - 1) * (p[i][j] - p[i][j-1]) / 2
		}
		ans += cnt
	}
	return ans
}

func solveHonestly(n int, a []int) int {
	var ans int
	for i := 0; i < n; i++ {
		m := make(map[int]struct{})
		for j := i; j < n; j++ {
			m[a[j]] = struct{}{}
			ans += len(m)
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
