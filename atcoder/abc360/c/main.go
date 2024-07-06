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
	a := nextIntSlice(n)
	w := nextIntSlice(n)

	m := make([][]int, int(1e5)+1)
	for i := range a {
		m[a[i]] = append(m[a[i]], w[i])
	}
	var ans int
	for _, s := range m {
		if len(s) <= 1 {
			continue
		}
		sort.Ints(s)
		for j := 0; j < len(s)-1; j++ {
			ans += s[j]
		}
	}
	Print(ans)
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
