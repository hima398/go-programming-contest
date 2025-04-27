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
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	var a, b []int
	for i := 0; i < n; i++ {
		a = append(a, nextInt())
		b = append(b, nextInt())
	}
	ans := solve(n, a, b)
	PrintString(ans)
}

func solve(n int, a, b []int) string {
	type task struct {
		a, b int
	}
	var ts []task
	for i := range a {
		ts = append(ts, task{a[i], b[i]})
	}
	sort.Slice(ts, func(i, j int) bool {
		return ts[i].b < ts[j].b
	})
	s := 0
	for _, t := range ts {
		if s+t.a > t.b {
			return "No"
		}
		s += t.a
	}
	return "Yes"
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func PrintString(x string) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
