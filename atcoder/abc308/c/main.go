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
	type toss struct {
		i, a, b int
	}
	var ts []toss
	for i := 0; i < n; i++ {
		ts = append(ts, toss{i, a[i], b[i]})
	}
	sort.Slice(ts, func(i, j int) bool {
		if (ts[j].a+ts[j].b)*ts[i].a == (ts[i].a+ts[i].b)*ts[j].a {
			return ts[i].i < ts[j].i
		}
		return (ts[j].a+ts[j].b)*ts[i].a > (ts[i].a+ts[i].b)*ts[j].a
	})
	var ans []int
	for _, ti := range ts {
		ans = append(ans, ti.i+1)
	}
	PrintHorizonaly(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func PrintHorizonaly(x []int) {
	defer out.Flush()
	fmt.Fprintf(out, "%d", x[0])
	for i := 1; i < len(x); i++ {
		fmt.Fprintf(out, " %d", x[i])
	}
	fmt.Fprintln(out)
}
