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
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, k := nextInt(), nextInt()
	a := nextIntSlice(n)

	ans := solve(n, k, a)

	for _, v := range ans {
		Print(v)
	}
}

func listDivisor(x int) []int {
	m := make(map[int]struct{})
	for i := 1; i*i <= x; i++ {
		if x%i == 0 {
			m[i] = struct{}{}
			m[x/i] = struct{}{}
		}
	}
	var res []int
	for k := range m {
		res = append(res, k)
	}
	//sort.Ints(res)
	return res
}

func solve(n, k int, a []int) []int {
	const MaxA = int(1e6)
	//約数xを含む個数
	cnt := make([]int, MaxA+1)
	for _, ai := range a {
		for i := 1; i*i <= ai; i++ {
			if ai%i == 0 {
				if ai/i == i {
					cnt[i]++
				} else {
					cnt[i]++
					cnt[ai/i]++
				}
			}
		}
	}
	var ans []int
	for _, ai := range a {
		ds := listDivisor(ai)
		var v int
		for i := 0; i <= len(ds)/2; i++ {
			if cnt[ds[i]] >= k {
				v = Max(v, ds[i])
			}
			if cnt[ai/ds[i]] >= k {
				v = Max(v, ai/ds[i])
			}
		}
		ans = append(ans, v)
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

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
