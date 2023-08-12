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
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, x, y := nextInt(), nextInt(), nextInt()
	a := nextIntSlice(n)

	ok := solve(n, x, y, a)
	if ok {
		PrintString("Yes")
	} else {
		PrintString("No")
	}
}

func solve(n, x, y int, a []int) bool {
	dpx := make([]map[int]struct{}, n+2)
	dpy := make([]map[int]struct{}, n+2)
	for i := 0; i < n+2; i++ {
		dpx[i] = make(map[int]struct{})
		dpy[i] = make(map[int]struct{})
	}
	dpx[2][a[0]] = struct{}{}
	dpy[2][0] = struct{}{}
	for i := 3; i <= n+1; i++ {
		if i&1 == 0 {
			for k, v := range dpx[i-1] {
				dpx[i][k-a[i-2]] = v
				dpx[i][k+a[i-2]] = v
			}
			for k, v := range dpy[i-1] {
				dpy[i][k] = v
			}
		} else {
			for k, v := range dpx[i-1] {
				dpx[i][k] = v
			}
			for k, v := range dpy[i-1] {
				dpy[i][k-a[i-2]] = v
				dpy[i][k+a[i-2]] = v
			}
		}
	}
	_, foundX := dpx[n+1][x]
	_, foundY := dpy[n+1][y]
	return foundX && foundY
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
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
