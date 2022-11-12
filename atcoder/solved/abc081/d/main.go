package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func solve(n int, a []int) ([][2]int, []int) {
	ix, in := -1, -1
	mx, mn := -int(1e6), int(1e6)
	for i, ai := range a {
		if mx < ai {
			mx = ai
			ix = i
		}
		if mn > ai {
			mn = ai
			in = i
		}
	}
	var ans [][2]int
	if Abs(mx) >= Abs(mn) {
		for i := 0; i < n; i++ {
			if i == ix {
				continue
			}
			ans = append(ans, [2]int{ix + 1, i + 1})
			a[i] += mx
		}
		ans = append(ans, [2]int{ix + 1, ix + 1})
		for i := 1; i < n; i++ {
			ans = append(ans, [2]int{i, i + 1})
			a[i] += a[i-1]
		}
	} else {
		for i := 0; i < n; i++ {
			if i == in {
				continue
			}
			ans = append(ans, [2]int{in + 1, i + 1})
			a[i] += mn
		}
		ans = append(ans, [2]int{in + 1, in + 1})
		for i := n - 1; i > 0; i-- {
			ans = append(ans, [2]int{i + 1, i})
			a[i-1] += a[i]
		}
	}
	return ans, a
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	a := nextIntSlice(n)
	ans, a := solve(n, a)
	//fmt.Println(a)
	PrintInt(len(ans))
	PrintVertically(ans)
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

func PrintHorizonaly(x [2]int) {
	defer out.Flush()
	fmt.Fprintf(out, "%d", x[0])
	for i := 1; i < len(x); i++ {
		fmt.Fprintf(out, " %d", x[i])
	}
	fmt.Fprintln(out)
}

func PrintVertically(x [][2]int) {
	defer out.Flush()
	for _, v := range x {
		PrintHorizonaly(v)
		//fmt.Fprintln(out, v)
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
