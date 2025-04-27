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

	n, s := nextInt(), nextInt()
	a := nextIntSlice(n)

	if solve(n, s, a) {
		Print("Yes")
	} else {
		Print("No")
	}
}

func solve(n, s int, a []int) bool {
	ss := make([]int, 2*n+1)
	for i := 0; i < 2*n; i++ {
		ss[i+1] = ss[i] + a[i%n]
	}
	var sum int
	for _, ai := range a {
		sum += ai
	}

	rem := s % sum

	for i := range ss {
		if ss[i] < rem {
			continue
		}
		idx := sort.Search(2*n+1, func(j int) bool {
			return ss[i]-ss[j] <= rem
		})
		if ss[i]-ss[idx] == rem {
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
