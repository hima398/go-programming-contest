package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func solve(n, m int, a, b []int) int {
	ca := make([]int, n)
	copy(ca, a)
	mn, mx := 0, int(1e9)+1
	ca = append([]int{mn}, ca...)
	ca = append(ca, mx)

	type section struct {
		l, r int
	}
	c := make(map[section][]int)

	for j := 0; j < m; j++ {
		idx := sort.Search(len(ca), func(i int) bool {
			return b[j] <= ca[i]
		})
		l, r := ca[idx-1], ca[idx]
		c[section{l, r}] = append(c[section{l, r}], b[j])
	}

	var ans int
	for k, v := range c {
		for len(v) > 0 {
			if k.l == mn {
				ans += k.r - v[len(v)-1]
				k.r = v[len(v)-1]
				v = v[:len(v)-1]
			} else if k.r == mx {
				ans += v[0] - k.l
				k.l = v[0]
				v = v[1:]
			} else {
				dl, dr := v[0]-k.l, k.r-v[len(v)-1]
				if dl <= dr {
					ans += v[0] - k.l
					k.l = v[0]
					v = v[1:]
				} else {
					ans += k.r - v[len(v)-1]
					k.r = v[len(v)-1]
					v = v[:len(v)-1]
				}
				//38 55 64 68 70 92 97
				// 38:1, 97:2, 92:5, 55:17, 64:9, 68:4, 70:2
				//
			}
		}
	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, m := nextInt(), nextInt()
	a := nextIntSlice(n)
	b := nextIntSlice(m)
	ans := solve(n, m, a, b)
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

func PrintFloat64(x float64) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}

func PrintString(x string) {
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

func PrintVertically(x []int) {
	defer out.Flush()
	for _, v := range x {
		fmt.Fprintln(out, v)
	}
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
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

func Floor(x, y int) int {
	return x / y
}

func Ceil(x, y int) int {
	return (x + y - 1) / y
}
