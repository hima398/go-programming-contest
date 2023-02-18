package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	h, w := nextInt(), nextInt()
	s := make([][]string, h)
	for i := 0; i < h; i++ {
		s[i] = strings.Split(nextString(), "")
	}
	q := nextInt()
	var a, b []int
	for i := 0; i < q; i++ {
		a = append(a, nextInt())
		b = append(b, nextInt())
	}
	ans := solve(h, w, s, q, a, b)
	for _, v := range ans {
		PrintString(v)
	}
}

func solve(h, w int, s [][]string, q int, a, b []int) []string {
	var si, sj int
	for i := 0; i < q; i++ {
		if si < a[i] {
			si = a[i] - 1 - si
		} else {
			si = a[i] - 1 - si + h
		}
		if sj < b[i] {
			sj = b[i] - 1 - sj
		} else {
			sj = b[i] - 1 - sj + w
		}
	}
	t := make([][]string, h)
	for i := 0; i < h; i++ {
		t[i] = make([]string, w)
	}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			var ni, nj int
			if q%2 == 0 {
				ni = i + si
				nj = j + sj
			} else {
				ni = h - i + si
				nj = w - j + sj
			}
			if ni >= h {
				ni -= h
			}
			if nj >= w {
				nj -= w
			}
			t[ni][nj] = s[i][j]
		}
	}
	var ans []string
	for i := 0; i < h; i++ {
		ans = append(ans, strings.Join(t[i], ""))
	}
	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func PrintString(x string) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
