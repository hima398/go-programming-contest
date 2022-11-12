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
	var l, r, s []int
	for i := 0; i < n; i++ {
		l = append(l, nextInt()-1)
		r = append(r, nextInt()-1)
		s = append(s, nextInt())
	}
	ans := solve(n, m, l, r, s)
	PrintInt(ans)
}

func solve(n, m int, l, r, s []int) int {
	f := make([]int, m+1)
	ans := 0
	for i := 0; i < n; i++ {
		f[l[i]] += s[i]
		f[r[i]+1] -= s[i]
		ans += s[i]
	}
	for i := 0; i < m; i++ {
		f[i+1] += f[i]
	}
	//fmt.Println(f)
	min := ans
	for i := 0; i < m; i++ {
		min = Min(min, f[i])
	}
	ans -= min
	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
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
