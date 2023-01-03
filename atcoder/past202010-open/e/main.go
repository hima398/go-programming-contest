package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	s := nextString()
	ans := solve(n, s)
	PrintString(ans)
}

func solve(n int, s string) string {
	t := strings.Split(s, "")
	sort.Slice(t, func(i, j int) bool {
		return t[i] < t[j]
	})
	rs := ""
	for i := n - 1; i >= 0; i-- {
		rs += string(s[i])
	}
	for {
		isSame1 := true
		for i := 0; i < n; i++ {
			isSame1 = isSame1 && t[i] == string(s[i])
		}
		isSame2 := true
		for i := 0; i < n; i++ {
			isSame2 = isSame2 && t[i] == string(rs[i])
		}
		if !isSame1 && !isSame2 {
			return strings.Join(t, "")
		}
		if !NextPermutation(sort.StringSlice(t)) {
			break
		}
	}
	return "None"
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
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
