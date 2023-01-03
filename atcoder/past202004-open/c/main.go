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
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	s := make([]string, n)
	for i := 0; i < n; i++ {
		s[i] = nextString()
	}
	ans := solve(n, s)
	PrintVertically(ans)
}

func solve(n int, s []string) []string {
	t := make([][]string, n)
	for i := 0; i < n; i++ {
		t[i] = strings.Split(s[i], "")
	}
	dk := []int{-1, 0, 1}
	for i := n - 2; i >= 0; i-- {
		for j := 1; j < 2*(n-1); j++ {
			if t[i][j] != "#" {
				continue
			}
			isNearX := false
			for _, ki := range dk {
				ni, nj := i+1, j+ki
				isNearX = isNearX || t[ni][nj] == "X"
			}
			if isNearX {
				t[i][j] = "X"
			}
		}
	}
	ans := make([]string, n)
	for i, ti := range t {
		ans[i] = strings.Join(ti, "")
	}
	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func PrintVertically(x []string) {
	defer out.Flush()
	for _, v := range x {
		fmt.Fprintln(out, v)
	}
}
