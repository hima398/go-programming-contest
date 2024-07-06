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

	n, a, b := nextInt(), nextInt(), nextInt()
	d := nextIntSlice(n)

	//ok := solveHonestly(n, a, b, d)
	ok := solve(n, a, b, d)

	if ok {
		Print("Yes")
	} else {
		Print("No")
	}
}

// O((a+b)n)~int(1e14)なので時間切れ
func solveHonestly(n, a, b int, d []int) bool {
	for i := 0; i < a+b; i++ {
		ok := true
		for _, di := range d {
			ok = ok && (di+i)%(a+b) < a
		}
		if ok {
			return true
		}
	}
	return false
}

func solve(n, a, b int, d []int) bool {
	m := make(map[int]struct{})
	for _, di := range d {
		m[di%(a+b)] = struct{}{}
	}
	var e []int
	for k := range m {
		e = append(e, k)
	}
	sort.Ints(e)

	nn := len(e)
	for i := 0; i < nn; i++ {
		e = append(e, e[i]+a+b)
	}

	for i := 0; i < nn; i++ {
		l, r := e[i], e[i+nn-1]
		if r-l+1 <= a {
			return true
		}
	}
	return false
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
