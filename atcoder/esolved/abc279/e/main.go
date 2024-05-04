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

	n, m := nextInt(), nextInt()
	a := nextIntSlice(m)
	ans := solve(n, m, a)
	PrintVertically(ans)
}

func solve(n, m int, a []int) []int {

	var b []int
	for i := 0; i < n; i++ {
		b = append(b, i)
	}
	for i := 0; i < m; i++ {
		idx := a[i] - 1
		b[idx], b[idx+1] = b[idx+1], b[idx]
	}
	pos := make([]int, n)
	for i := 0; i < n; i++ {
		pos[b[i]] = i
	}

	for i := 0; i < n; i++ {
		b[i] = i
	}
	var ans []int
	for i := 0; i < m; i++ {
		idx := a[i] - 1
		if b[idx] == 0 {
			ans = append(ans, pos[b[idx+1]]+1)
		} else if b[idx+1] == 0 {
			ans = append(ans, pos[b[idx]]+1)
		} else {
			ans = append(ans, pos[0]+1)
		}
		b[idx], b[idx+1] = b[idx+1], b[idx]
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

func PrintVertically(x []int) {
	defer out.Flush()
	for _, v := range x {
		fmt.Fprintln(out, v)
	}
}
