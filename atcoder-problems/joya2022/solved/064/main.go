package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	h1, w1 := nextInt(), nextInt()
	a := make([][]int, h1)
	for i := 0; i < h1; i++ {
		a[i] = nextIntSlice(w1)
	}
	h2, w2 := nextInt(), nextInt()
	b := make([][]int, h2)
	for i := 0; i < h2; i++ {
		b[i] = nextIntSlice(w2)
	}
	ans := solve(h1, w1, a, h2, w2, b)
	PrintString(ans)
}

func solve(h1, w1 int, a [][]int, h2, w2 int, b [][]int) string {
	mh, mw := 1<<h1-1, 1<<w1-1
	for ph := 0; ph <= mh; ph++ {
		for pw := 0; pw <= mw; pw++ {
			if bits.OnesCount(uint(ph)) != h2 || bits.OnesCount(uint(pw)) != w2 {
				continue
			}
			ih := 0
			ok := true
			for i := 0; i < h1; i++ {
				if (ph>>i)&1 == 0 {
					continue
				}
				iw := 0
				for j := 0; j < w1; j++ {
					if (pw>>j)&1 == 0 {
						continue
					}
					ok = ok && a[i][j] == b[ih][iw]
					iw++
				}
				ih++
			}
			if ok {
				return "Yes"
			}
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
