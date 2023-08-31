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
	ok := solve(n, a, b)
	if ok {
		PrintString("Yes")
	} else {
		PrintString("No")
	}
}

func solve(n int, a, b []int) bool {
	type task struct {
		a, b int
	}
	var ts []task
	for i := 0; i < n; i++ {
		ts = append(ts, task{a[i], b[i]})
	}
	sort.Slice(ts, func(i, j int) bool {
		if ts[i].b == ts[j].b {
			return ts[i].a < ts[j].a
		}
		return ts[i].b < ts[j].b
	})
	var cur int
	for _, t := range ts {
		cur += t.a
		if cur > t.b {
			return false
		}
	}
	return true
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
