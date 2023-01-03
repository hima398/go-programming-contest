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

	n := nextInt()
	w := nextIntSlice(n)
	b := nextIntSlice(n)
	ans := solve(n, w, b)
	PrintString(ans)
}

func solve(n int, w, b []int) string {
	var maxW, maxB int
	for i := 0; i < n; i++ {
		maxW = Max(maxW, w[i])
		maxB = Max(maxB, b[i])
	}
	grundy := make([][]int, maxW+1)
	m := maxB + maxW*(maxW+1)/2
	for i := 0; i <= maxW; i++ {
		grundy[i] = make([]int, m+1)
	}
	for i := 0; i <= maxW; i++ {
		for j := 0; j <= m; j++ {
			mex := make([]int, m+1)
			if i >= 1 {
				if j+i > m {
					break
				}
				mex[grundy[i-1][j+i]] = 1
			}
			if j >= 2 {
				for k := 1; k <= j/2; k++ {
					mex[grundy[i][j-k]] = 1
				}
			}
			for k := 0; k <= m; k++ {
				if mex[k] == 0 {
					grundy[i][j] = k
					break
				}
			}
		}
	}
	ans := 0
	for i := 0; i < n; i++ {
		ans ^= grundy[w[i]][b[i]]
	}
	if ans == 0 {
		return "Second"
	} else {
		return "First"
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
