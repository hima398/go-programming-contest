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

	n, m := nextInt(), nextInt()
	b := nextIntSlice(n)
	w := nextIntSlice(m)

	ans := solve(n, m, b, w)

	Print(ans)
}

func solve(n, m int, b, w []int) int {
	sort.Slice(b, func(i, j int) bool {
		return b[i] > b[j]
	})
	sort.Slice(w, func(i, j int) bool {
		return w[i] > w[j]
	})

	//白を価値が高い順にx個選ぶ合計
	sw := make([]int, m+1)
	for i := 0; i < m; i++ {
		sw[i+1] = sw[i] + w[i]
	}
	//黒をx個以上選ぶ価値の最大値
	var s int
	for _, bi := range b {
		if bi > 0 {
			s += bi
		}
	}
	sb := make([]int, n+1)
	for i := 0; i <= n; i++ {
		sb[i] = s
	}
	for i, bi := range b {
		if bi < 0 {
			sb[i+1] = sb[i] + bi
		}
	}
	//Print(sb)
	//Print(sw)

	var ans int
	for i := 0; i <= Min(n, m); i++ {
		ans = Max(ans, sw[i]+sb[i])
	}

	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func nextIntSlice(n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = nextInt()
	}
	return s
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

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
