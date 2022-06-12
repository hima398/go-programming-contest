package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func solve(n int, a []string) string {
	var ls, rs string
	for i := 0; i < n/2; i++ {
		l := make(map[byte]struct{})
		r := make(map[byte]struct{})
		for j := 0; j < n; j++ {
			l[a[i][j]] = struct{}{}
			r[a[n-i-1][j]] = struct{}{}
		}

		var found bool
		for k := range l {
			if _, ok := r[k]; ok {
				ls += string(k)
				rs = string(k) + rs
				found = true
			}
		}
		if !found {
			return "-1"
		}
	}
	if n%2 == 1 {
		ls += string(a[n/2][0])
	}
	return ls + rs
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	a := make([]string, n)
	for i := 0; i < n; i++ {
		a[i] = nextString()
	}
	ans := solve(n, a)

	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}
