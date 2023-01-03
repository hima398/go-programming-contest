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

	n, x, y := nextInt(), nextInt(), nextInt()
	a := nextIntSlice(n)
	ans := solve(n, x, y, a)
	PrintString(ans)
}

func solve(n, x, y int, a []int) string {
	dpx := make(map[int]struct{})
	dpy := make(map[int]struct{})
	dpx[a[0]] = struct{}{}
	dpy[0] = struct{}{}
	for i := 1; i < n; i++ {
		if i%2 == 0 {
			next := make(map[int]struct{})
			for k := range dpx {
				next[k+a[i]] = struct{}{}
				next[k-a[i]] = struct{}{}
			}
			dpx = next
		} else {
			next := make(map[int]struct{})
			for k := range dpy {
				next[k+a[i]] = struct{}{}
				next[k-a[i]] = struct{}{}
			}
			dpy = next
		}
	}
	if _, foundX := dpx[x]; foundX {
		if _, foundY := dpy[y]; foundY {
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
