package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	m := make(map[int]struct{})
	m[0] = struct{}{}
	cur := 0
	for _, v := range a {
		cur += v
		cur %= 360
		m[cur] = struct{}{}
	}

	var b []int
	for k := range m {
		b = append(b, k)
	}
	b = append(b, 360)
	sort.Ints(b)
	ans := 0
	for i := 1; i < len(b); i++ {
		ans = Max(ans, b[i]-b[i-1])
	}
	PrintInt(ans)
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

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
