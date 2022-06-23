package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func solve(n int) [][]int {
	ans := make([][]int, n)
	for i := 0; i < n; i++ {
		ans[i] = make([]int, n)
	}
	var s []int
	for i := 1; i <= n*n; i++ {
		s = append(s, i)
	}
	var cnt int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			var v int
			if cnt%2 == 0 {
				v = s[0]
				s = s[1:]
			} else {
				v = s[len(s)-1]
				s = s[:len(s)-1]
			}
			ans[i][j] = v
			cnt++
		}
	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	ans := solve(n)
	PrintVertically(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func PrintHorizonaly(x []int) {
	defer out.Flush()
	fmt.Fprintf(out, "%d", x[0])
	for i := 1; i < len(x); i++ {
		fmt.Fprintf(out, " %d", x[i])
	}
	fmt.Fprintln(out)
}

func PrintVertically(x [][]int) {
	defer out.Flush()
	for _, v := range x {
		PrintHorizonaly(v)
	}
}
