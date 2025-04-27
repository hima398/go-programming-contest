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

	n, m := nextInt(), nextInt()
	var a, b []int
	for i := 0; i < m; i++ {
		a = append(a, nextInt()-1)
		b = append(b, nextInt()-1)
	}

	ans := solve(n, m, a, b)

	Print(ans)
}

func solve(n, m int, a, b []int) int {
	lines := make([]int, n)
	for i := 0; i < m; i++ {
		lines[(a[i]+b[i])%n]++
	}
	//fmt.Println("# lines = ", lines)

	ans := m * (m - 1) / 2
	//fmt.Println("# ans = ", ans)
	for i := 0; i < n; i++ {
		ans -= lines[i] * (lines[i] - 1) / 2
	}

	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func Print(x any) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
