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

	n, q := nextInt(), nextInt()
	x := nextIntSlice(q)
	ans := solve(n, q, x)
	PrintHorizonaly(ans)
}

func solve(n, q int, x []int) []int {
	for i := range x {
		x[i]--
	}
	var ans []int
	var pos []int
	for i := 0; i < n; i++ {
		ans = append(ans, i)
		pos = append(pos, i)
	}
	for i := 0; i < q; i++ {
		i, j := pos[x[i]], pos[x[i]]+1
		if i == n-1 {
			j = n - 2
		}
		ans[i], ans[j] = ans[j], ans[i]
		pos[ans[i]], pos[ans[j]] = pos[ans[j]], pos[ans[i]]
	}
	for i := range ans {
		ans[i]++
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

func PrintHorizonaly(x []int) {
	defer out.Flush()
	fmt.Fprintf(out, "%d", x[0])
	for i := 1; i < len(x); i++ {
		fmt.Fprintf(out, " %d", x[i])
	}
	fmt.Fprintln(out)
}
