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

	n, k := nextInt(), nextInt()
	a := nextIntSlice(n)
	b := nextIntSlice(k)
	m := make(map[int]struct{})
	for _, bi := range b {
		m[bi-1] = struct{}{}
	}
	var m1, m2 int
	for i, ai := range a {
		if _, found := m[i]; found {
			m2 = Max(m2, ai)
		} else {
			m1 = Max(m1, ai)
		}
	}
	if m1 == m2 || m2 > m1 {
		PrintString("Yes")
	} else {
		PrintString("No")
	}
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

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
