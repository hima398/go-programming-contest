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
	l := nextIntSlice(n)

	ans := solve(n, m, l)

	Print(ans)
}

func solve(n, m int, l []int) int {
	var max int
	for _, v := range l {
		max = Max(max, v)
	}
	ok, ng := 1<<60, max-1
	//横幅をxとしたときm行に収まるかを返す
	check := func(x int) bool {
		w := l[0]
		row := 1
		for i := 1; i < n; i++ {
			if w+l[i]+1 <= x {
				w += l[i] + 1
			} else {
				row++
				w = l[i]
			}
		}
		return row <= m
	}
	for ok-ng > 1 {
		mid := (ok + ng) / 2
		if check(mid) {
			ok = mid
		} else {
			ng = mid
		}
	}
	return ok
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

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
