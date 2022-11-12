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

func solve(n int, a []int) int {
	sort.Ints(a)
	var q []int
	var ans int
	for len(q)+len(a) > 1 {
		var min, max int
		if len(q) > 0 {
			min = q[len(q)-1]
		} else {
			min = a[0]
		}
		if len(a) > 0 {
			max = a[len(a)-1]
			a = a[:len(a)-1]
		} else {
			max = q[0]
			q = q[1:]
		}
		if max%min != 0 {
			q = append(q, max%min)
		}
		ans++
	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	a := nextIntSlice(n)
	ans := solve(n, a)
	PrintInt(ans)
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

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
