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

	n := nextInt()
	q := nextInt()
	t, a, b := make([]int, q), make([]int, q), make([]int, q)
	for i := 0; i < q; i++ {
		t[i] = nextInt()
		if t[i] != 3 {
			a[i] = nextInt() - 1
			b[i] = nextInt() - 1
		}
	}

	ans := solve(n, q, t, a, b)
	PrintVertically(ans)
}

func solve(n, q int, t, a, b []int) []int {
	r, c := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		r[i] = i
		c[i] = i
	}
	var isTransposed bool
	var ans []int
	for k := 0; k < q; k++ {
		switch t[k] {
		case 1:
			if isTransposed {
				c[a[k]], c[b[k]] = c[b[k]], c[a[k]]
			} else {
				r[a[k]], r[b[k]] = r[b[k]], r[a[k]]
			}
		case 2:
			if isTransposed {
				r[a[k]], r[b[k]] = r[b[k]], r[a[k]]
			} else {
				c[a[k]], c[b[k]] = c[b[k]], c[a[k]]
			}
		case 3:
			isTransposed = !isTransposed
		case 4:
			var i, j int
			if isTransposed {
				i, j = r[b[k]], c[a[k]]
			} else {
				i, j = r[a[k]], c[b[k]]
			}
			ans = append(ans, n*i+j)
		}
	}
	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func PrintVertically(x []int) {
	defer out.Flush()
	for _, v := range x {
		fmt.Fprintln(out, v)
	}
}
