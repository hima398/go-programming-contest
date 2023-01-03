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
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	a := nextIntSlice(n)
	ans := solve(n, a)
	PrintInt(ans)
}

func dfs(cur int, a []int) int {
	if cur < 0 {
		return 0
	}
	var s, t []int
	for i := 0; i < len(a); i++ {
		if a[i]>>cur&1 == 0 {
			s = append(s, a[i])
		} else {
			t = append(t, a[i])
		}
	}
	if len(s) == 0 {
		return dfs(cur-1, t)
	}
	if len(t) == 0 {
		return dfs(cur-1, s)
	}
	return Min(dfs(cur-1, s), dfs(cur-1, t)) | (1 << cur)
}

func solve(n int, a []int) int {
	return dfs(29, a)
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

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
