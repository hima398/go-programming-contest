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
	var s []string
	for i := 0; i < n; i++ {
		s = append(s, nextString())
	}

	ans := solve(n, s)

	Print(ans)
}

func solve(n int, s []string) int {
	//i行目のoの数
	row := make([]int, n)
	//j列目のoの数
	col := make([]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if s[i][j] == 'o' {
				row[i]++
			}
		}
	}
	for j := 0; j < n; j++ {
		for i := 0; i < n; i++ {
			if s[i][j] == 'o' {
				col[j]++
			}
		}
	}

	var ans int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if s[i][j] == 'o' {
				ans += (row[i] - 1) * (col[j] - 1)
			}
		}
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
