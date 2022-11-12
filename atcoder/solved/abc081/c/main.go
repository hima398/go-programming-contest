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

func solve(n, k int, a []int) int {
	type number struct {
		v, cnt int
	}
	m := make(map[int]int)
	for _, ai := range a {
		m[ai]++
	}
	var ns []number
	for k, v := range m {
		ns = append(ns, number{k, v})
	}
	sort.Slice(ns, func(i, j int) bool {
		return ns[i].cnt > ns[j].cnt
	})
	//fmt.Println(ns)
	idx := 0
	for i := 0; i < Min(k, len(ns)); i++ {
		idx++
	}
	var ans int
	for i := idx; i < len(ns); i++ {
		ans += ns[i].cnt
	}
	return ans
}
func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, k := nextInt(), nextInt()
	a := nextIntSlice(n)
	ans := solve(n, k, a)
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
