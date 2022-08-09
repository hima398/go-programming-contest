package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const Mod = 1000000007

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	t := nextInt()
	var n, m []int
	a := make([][]int, t)
	for i := 0; i < t; i++ {
		n = append(n, nextInt())
		m = append(m, nextInt())
		a[i] = append(a[i], nextIntSlice(n[i])...)
	}
	var ans []string
	for i := 0; i < t; i++ {
		ok := true
		for j := 1; j < n[i] && ok; j++ {
			ok = ok && a[i][j-1]+a[i][j] >= m[i]
		}
		if ok {
			ans = append(ans, "Yes")
		} else {
			ans = append(ans, "No")
		}
	}
	PrintVertically(ans)
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

func PrintVertically(x []string) {
	defer out.Flush()
	for _, v := range x {
		fmt.Fprintln(out, v)
	}
}
