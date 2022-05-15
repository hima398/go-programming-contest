package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func solve(n, k int, a []int) []int {
	ans := make([]int, n)
	copy(ans, a)
	for kk := 0; kk < k; kk++ {
		b := make([]int, n)
		for i := 0; i < n; i++ {
			l := Max(i-ans[i], 0)
			r := Min(i+ans[i], n-1)
			b[l]++
			if r+1 < n {
				b[r+1]--
			}
		}
		for i := 1; i < n; i++ {
			b[i] += b[i-1]
		}
		ok := true
		for i := 0; i < n; i++ {
			ok = ok && ans[i] == b[i]
		}
		if ok {
			break
		}
		copy(ans, b)
	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, k := nextInt(), nextInt()
	a := nextIntSlice(n)
	ans := solve(n, k, a)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	fmt.Fprintf(out, "%d", ans[0])
	for i := 1; i < n; i++ {
		fmt.Fprintf(out, " %d", ans[i])
	}
	fmt.Fprintln(out)
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
