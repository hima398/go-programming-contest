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

func isPalindrome(t []int) bool {
	//fmt.Println(t)
	n := len(t)
	ok := true
	for i := 0; i < n/2; i++ {
		j := n - 1 - i
		ok = ok && (t[i] == t[j])
	}
	return ok
}

func main() {
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, k := nextInt(), nextInt()
	s := nextString()
	var p []int
	for i := 0; i < n; i++ {
		p = append(p, int(s[i]-'a'))
	}
	sort.Ints(p)
	var ans int
	for {
		var contains bool
		for i := 0; i <= n-k; i++ {
			contains = contains || isPalindrome(p[i:i+k])
		}
		if !contains {
			ans++
		}

		if !NextPermutation(sort.IntSlice(p)) {
			break
		}
	}

	Print(ans)
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

func Print(x any) {
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
