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

	h := nextIntSlice(3)
	w := nextIntSlice(3)
	ans := solve(h, w)
	PrintInt(ans)
}

func solve(h, w []int) int {
	var ans int
	for a1 := 1; a1 <= 30; a1++ {
		for a3 := 1; a3 <= 30; a3++ {
			for a5 := 1; a5 <= 30; a5++ {
				for a7 := 1; a7 <= 30; a7++ {
					for a9 := 1; a9 <= 30; a9++ {
						a2 := h[0] - (a1 + a3)
						a8 := h[2] - (a7 + a9)
						a4 := w[0] - (a1 + a7)
						a6 := w[2] - (a3 + a9)
						if a2 <= 0 || a4 <= 0 || a6 <= 0 || a8 <= 0 {
							continue
						}
						ok := w[1] == a2+a5+a8
						ok = ok && h[1] == a4+a5+a6
						if ok {
							//fmt.Println(a1, a2, a3, a4, a5, a6, a7, a8, a9)
							ans++
						}
					}
				}
			}
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
