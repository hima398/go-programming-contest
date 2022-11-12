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
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	k := nextInt()
	m := make(map[int]struct{})
	var q []int
	for i := 1; i <= 9; i++ {
		m[i] = struct{}{}
		q = append(q, i)
	}
	for len(m) <= int(1e5) && len(q) > 0 {
		//fmt.Println(len(m))
		cur := q[0]
		q = q[1:]
		x := cur % 10
		for i := Max(x-1, 0); i <= Min(x+1, 9); i++ {
			next := cur*10 + i
			if _, found := m[next]; !found {
				m[next] = struct{}{}
				q = append(q, next)
			}
		}
	}
	var l []int
	for k := range m {
		l = append(l, k)
	}
	sort.Ints(l)
	ans := l[k-1]
	PrintInt(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
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
