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
	var s []string
	for i := 0; i < n; i++ {
		s = append(s, nextString())
	}
	var idxes []int
	for i := 0; i < n; i++ {
		idxes = append(idxes, i)
	}
	computeDist := func(s, t string) int {
		var res int
		for i := 0; i < m; i++ {
			if s[i] != t[i] {
				res++
			}
		}
		return res
	}
	for {
		ok := true
		for i := 0; i < n-1; i++ {
			ok = ok && (computeDist(s[idxes[i]], s[idxes[i+1]]) == 1)
		}
		if ok {
			PrintString("Yes")
			return
		}
		if !NextPermutation(sort.IntSlice(idxes)) {
			break
		}
	}
	PrintString("No")
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
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
