package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func reverse(s []int) []int {
	var r []int
	for i := len(s) - 1; i >= 0; i-- {
		r = append(r, s[i])
	}
	return r
}

func solve(n, x int) []int {
	var l, r []int
	m := n / 2
	for i := 1; i <= m; i++ {
		if i != x {
			l = append(l, i)
		}
		if n-i+1 != x {
			r = append(r, n-i+1)
		}
	}
	if n%2 == 1 {
		if x < m+1 {
			l = append(l, m+1)
		} else if x > m+1 {
			r = append(r, m+1)
		}
	}
	//fmt.Println(l, r)
	if len(l) > len(r) {
		l, r = r, l
	}
	l = reverse(l)
	r = reverse(r)
	ans := []int{x}
	for len(l) > 0 {
		ans = append(ans, r[0])
		ans = append(ans, l[0])
		l = l[1:]
		r = r[1:]
	}
	if len(r) > 0 {
		ans = append(ans, r[0])
	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, x := nextInt(), nextInt()
	ans := solve(n, x)

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
