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

	n := nextInt()
	var a []string
	for i := 0; i < n; i++ {
		a = append(a, nextString())
	}

	ans := solve(n, a)

	for _, v := range ans {
		Print(v)
	}
}

func rotate(n, i, j, cnt int) (int, int) {
	ni, nj := i, j
	for i := 0; i < cnt; i++ {
		ni, nj = nj, n-1-ni
	}
	return ni, nj
}

func solve(n int, a []string) []string {
	b := make([][]string, n)
	for i := range b {
		b[i] = make([]string, n)
		for j := range b[i] {
			b[i][j] = "."
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			cnt := (Min(i, Min(j, Min(n-1-i, n-1-j))) + 1) % 4
			ni, nj := rotate(n, i, j, cnt)
			b[ni][nj] = string(a[i][j])
		}
	}
	var ans []string
	for i := 0; i < n; i++ {
		ans = append(ans, strings.Join(b[i], ""))
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

func Print(x any) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
