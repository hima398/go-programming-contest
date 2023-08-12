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

	h, w := nextInt(), nextInt()
	var a [][]int
	for i := 0; i < h; i++ {
		a = append(a, nextIntSlice(w))
	}
	n, ms := solve(h, w, a)
	PrintInt(n)
	for _, v := range ms {
		PrintHorizonaly(v)
	}
}

func solve(h, w int, a [][]int) (int, [][]int) {
	var n int
	var ms [][]int
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			//奇数
			if a[i][j]%2 == 1 {
				if j < w-1 {
					a[i][j]--
					ni, nj := i, j+1
					a[ni][nj]++
					n++
					ms = append(ms, []int{i + 1, j + 1, ni + 1, nj + 1})
				} else {
					//j==w-1、右端
					if i < h-1 {
						ni, nj := i+1, j
						a[i][j]--
						a[ni][nj]++
						n++
						ms = append(ms, []int{i + 1, j + 1, ni + 1, nj + 1})
					}
				}
			}
		}
	}
	return n, ms
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
