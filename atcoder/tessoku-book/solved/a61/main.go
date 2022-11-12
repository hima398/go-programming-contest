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

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, m := nextInt(), nextInt()
	var a, b []int
	for i := 0; i < m; i++ {
		a = append(a, nextInt()-1)
		b = append(b, nextInt()-1)
	}
	ans := solve(n, m, a, b)
	for i := 0; i < n; i++ {
		fmt.Printf("%d: ", i+1)
		PrintHorizonaly(ans[i])
	}
}

func solve(n, m int, a, b []int) [][]int {
	e := make([][]int, n)
	for i := 0; i < m; i++ {
		e[a[i]] = append(e[a[i]], b[i])
		e[b[i]] = append(e[b[i]], a[i])
	}
	ans := make([][]int, n)
	for i := 0; i < n; i++ {
		for _, to := range e[i] {
			ans[i] = append(ans[i], to+1)
		}
		sort.Ints(ans[i])
	}
	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func PrintHorizonaly(x []int) {
	defer out.Flush()
	fmt.Fprintf(out, "{")
	if len(x) > 0 {
		fmt.Fprintf(out, "%d", x[0])
		for i := 1; i < len(x); i++ {
			fmt.Fprintf(out, ", %d", x[i])
		}
	}
	fmt.Fprintf(out, "}")
	fmt.Fprintln(out)
}
