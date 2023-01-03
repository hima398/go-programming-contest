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

	h, w, n, r, c := nextInt(), nextInt(), nextInt(), nextInt(), nextInt()
	a := make([][]int, h)
	for i := 0; i < h; i++ {
		a[i] = nextIntSlice(w)
	}
	ans := solve(h, w, n, r, c, a)

	PrintVertically(ans)
}

func solve(h, w, n, r, c int, a [][]int) [][]int {
	ans := make([][]int, h-r+1)
	for i := 0; i <= h-r; i++ {
		ans[i] = make([]int, w-c+1)
	}
	for k := 0; k <= h-r; k++ {
		for l := 0; l <= w-c; l++ {
			d := make(map[int]int)
			for i := 0; i < h; i++ {
				for j := 0; j < w; j++ {
					if k <= i && i < k+r && l <= j && j < l+c {
						continue
					}
					d[a[i][j]]++
				}
			}
			ans[k][l] = len(d)
		}
	}
	return ans
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

func PrintHorizonaly(x []int) {
	defer out.Flush()
	fmt.Fprintf(out, "%d", x[0])
	for i := 1; i < len(x); i++ {
		fmt.Fprintf(out, " %d", x[i])
	}
	fmt.Fprintln(out)
}

func PrintVertically(x [][]int) {
	defer out.Flush()
	for _, v := range x {
		PrintHorizonaly(v)
	}
}
