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
	a := nextIntSlice(n)
	b := nextIntSlice(m)

	ma := make(map[int]struct{})
	for _, v := range a {
		ma[v] = struct{}{}
	}
	var c []int
	c = append(c, a...)
	c = append(c, b...)

	sort.Ints(c)

	for i := 0; i < len(c)-1; i++ {
		if _, f1 := ma[c[i]]; f1 {
			if _, f2 := ma[c[i+1]]; f2 {
				Print("Yes")
				return
			}
		}
	}
	Print("No")
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
