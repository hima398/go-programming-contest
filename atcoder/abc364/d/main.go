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
	n, q := nextInt(), nextInt()
	a := nextIntSlice(n)
	var b, k []int
	for i := 0; i < q; i++ {
		b = append(b, nextInt())
		k = append(k, nextInt())
	}

	ans := solve(n, q, a, b, k)
	for _, v := range ans {
		Print(v)
	}
}

func solve(n, q int, a, b, k []int) []int {
	sort.Ints(a)
	var ans []int
	check := func(b, k, x int) bool {
		l := sort.Search(n, func(i int) bool {
			return b-x <= a[i]
		})
		r := sort.Search(n, func(i int) bool {
			return b+x < a[i]
		})
		return r-l >= k
	}
	for i := 0; i < q; i++ {
		//距離xにaの点がk以上あるか？
		ng, ok := -1, int(2e9)
		for ok-ng > 1 {
			mid := (ok + ng) / 2
			if check(b[i], k[i], mid) {
				ok = mid
			} else {
				ng = mid
			}
		}
		ans = append(ans, ok)
	}
	return ans
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
