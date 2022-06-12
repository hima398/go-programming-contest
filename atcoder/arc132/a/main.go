package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func solve(n, q int, lr, lc, r, c []int) string {
	var ans string
	for i := 0; i < q; i++ {
		if lr[r[i]-1]+lc[c[i]-1] >= n+1 {
			ans += "#"
		} else {
			ans += "."
		}
	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	lr, lc := nextIntSlice(n), nextIntSlice(n)
	q := nextInt()
	var r, c []int
	for i := 0; i < q; i++ {
		r = append(r, nextInt())
		c = append(c, nextInt())
	}

	ans := solve(n, q, lr, lc, r, c)

	fmt.Println(ans)
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
