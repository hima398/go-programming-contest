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
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	a := nextIntSlice(1 << n)
	ans := solve(n, a)
	PrintVertically(ans)
}

func solve(n int, a []int) []int {
	type person struct {
		i, a, d int
	}
	nn := 1 << n
	tournament := make([]person, 2*nn)
	ans := make([]int, nn)
	for i := 0; i < nn; i++ {
		tournament[nn+i] = person{i, a[i], 1}
	}
	for i := 2*nn - 1; i > 3; i -= 2 {
		if tournament[i-1].a > tournament[i].a {
			tournament[i/2] = tournament[i-1]
			ans[tournament[i].i] = tournament[i].d
		} else {
			//t[i-1].a < t[i].a
			tournament[i/2] = tournament[i]
			ans[tournament[i-1].i] = tournament[i-1].d
		}
		tournament[i/2].d++
	}
	ans[tournament[2].i] = tournament[2].d
	ans[tournament[3].i] = tournament[3].d
	return ans
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextIntSlice(n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = nextInt()
	}
	return s
}

func PrintVertically(x []int) {
	defer out.Flush()
	for _, v := range x {
		fmt.Fprintln(out, v)
	}
}
