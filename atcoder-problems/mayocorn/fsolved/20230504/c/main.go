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
	s := nextString()
	ans := solve(n, s)
	PrintHorizonaly(ans)
}

func solve(n int, s string) []int {
	var l, r []int
	r = append(r, n)
	for i := n - 1; i >= 0; i-- {
		switch s[i] {
		case 'L':
			//右にiを積む
			r = append(r, i)
		case 'R':
			//左にiを積む
			l = append(l, i)
		}
	}
	ans := make([]int, n+1)
	for i := len(l) - 1; i >= 0; i-- {
		ans[len(l)-i-1] = l[i]
	}
	for i := 0; i < len(r); i++ {
		ans[len(l)+i] = r[i]
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

func PrintHorizonaly(x []int) {
	defer out.Flush()
	fmt.Fprintf(out, "%d", x[0])
	for i := 1; i < len(x); i++ {
		fmt.Fprintf(out, " %d", x[i])
	}
	fmt.Fprintln(out)
}
