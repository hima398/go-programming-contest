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

	n, p, q, r := nextInt(), nextInt(), nextInt(), nextInt()
	a := nextIntSlice(n)
	ans := solve(n, p, q, r, a)
	PrintString(ans)
}

func solve(n, p, q, r int, a []int) string {
	s := make([]int, n+1)
	for i := 1; i <= n; i++ {
		s[i] = s[i-1] + a[i-1]
	}
	m := make(map[int]struct{})
	for _, si := range s {
		m[si] = struct{}{}
	}
	for i := 0; i <= n; i++ {
		ok := true
		_, yFound := m[s[i]+p]
		ok = ok && yFound
		_, zFound := m[s[i]+p+q]
		ok = ok && zFound
		_, wFound := m[s[i]+p+q+r]
		ok = ok && wFound
		if ok {
			return "Yes"
		}
	}
	return "No"
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

func PrintString(x string) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
