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
	a := nextIntSlice(n)
	q := nextInt()
	t, idx, x := make([]int, q), make([]int, q), make([]int, q)
	for i := 0; i < q; i++ {
		t[i] = nextInt()
		switch t[i] {
		case 1:
			x[i] = nextInt()
		case 2:
			idx[i], x[i] = nextInt()-1, nextInt()
		case 3:
			idx[i] = nextInt() - 1
		}
	}
	ans := solve(n, a, q, t, idx, x)
	PrintVertically(ans)
}

func solve(n int, a []int, q int, t, idx, x []int) []int {
	var offset int
	diff := make(map[int]int)
	for i, v := range a {
		diff[i] = v
	}
	var ans []int
	for i, ti := range t {
		switch ti {
		case 1:
			offset = x[i]
			diff = make(map[int]int)
		case 2:
			diff[idx[i]] += x[i]
		case 3:
			ans = append(ans, offset+diff[idx[i]])
		}
	}
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
