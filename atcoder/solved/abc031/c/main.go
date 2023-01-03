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
	a := nextIntSlice(n)
	ans := -50*50 - 1
	for i := 0; i < n; i++ {
		aoki := -50*50 - 1
		idx := -1
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			s := 0
			for k := Min(i, j); k <= Max(i, j); k++ {
				if (k-Min(i, j))%2 == 1 {
					s += a[k]
				}
			}
			if aoki < s {
				idx = j
				aoki = s
			}
		}
		takahashi := 0
		for k := Min(i, idx); k <= Max(i, idx); k++ {
			if (k-Min(i, idx))%2 == 0 {
				takahashi += a[k]
			}
		}
		ans = Max(ans, takahashi)
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

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
