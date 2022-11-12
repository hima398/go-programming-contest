package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	s := nextString()
	ans := solve(s)
	PrintString(ans)
}

func solve(s string) string {
	m := make(map[string]int)
	for _, r := range s {
		m[string(r)]++
	}
	//fmt.Println(m)
	t := []int{int('a'), int('b'), int('c')}
	for {
		var pattern []string
		for i := 0; i < len(s); i++ {
			pattern = append(pattern, string(byte(t[i%3])))
		}
		m2 := make(map[string]int)
		for _, pi := range pattern {
			m2[pi]++
		}
		if m["a"] == m2["a"] && m["b"] == m2["b"] && m["c"] == m2["c"] {
			return "YES"
		}
		if !NextPermutation(sort.IntSlice(t)) {
			break
		}
	}
	return "NO"
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func PrintString(x string) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}

func NextPermutation(x sort.Interface) bool {
	n := x.Len() - 1
	if n < 1 {
		return false
	}
	j := n - 1
	for ; !x.Less(j, j+1); j-- {
		if j == 0 {
			return false
		}
	}
	l := n
	for !x.Less(j, l) {
		l--
	}
	x.Swap(j, l)
	for k, l := j+1, n; k < l; {
		x.Swap(k, l)
		k++
		l--
	}
	return true
}
