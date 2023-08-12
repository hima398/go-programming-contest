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

	h, w, n := nextInt(), nextInt(), nextInt()
	var r, c, a []int
	for i := 0; i < n; i++ {
		r = append(r, nextInt())
		c = append(c, nextInt())
		a = append(a, nextInt())
	}
	ans := solve(h, w, n, r, c, a)
	PrintVertically(ans)
}

func solve(h, w, n int, r, c, a []int) []int {
	type cell struct {
		i, r, c, a int
	}
	cs := make(map[int][]cell)
	for i := 0; i < n; i++ {
		cs[a[i]] = append(cs[a[i]], cell{i, r[i], c[i], a[i]})
	}
	//a[i]としてあり得る値を降順に持たせる
	var ks []int
	for k := range cs {
		ks = append(ks, k)
	}
	sort.Slice(ks, func(i, j int) bool {
		return ks[i] > ks[j]
	})
	ans := make([]int, n)
	//r行目から移動可能な数の最大値、c列目から移動可能な最大値
	rMax, cMax := make([]int, h+1), make([]int, w+1)
	for _, ai := range ks {
		for _, cl := range cs[ai] {
			ans[cl.i] = Max(rMax[cl.r], cMax[cl.c])
		}
		for _, cl := range cs[ai] {
			rMax[cl.r] = Max(rMax[cl.r], ans[cl.i]+1)
			cMax[cl.c] = Max(cMax[cl.c], ans[cl.i]+1)
		}
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

func PrintVertically(x []int) {
	defer out.Flush()
	for _, v := range x {
		fmt.Fprintln(out, v)
	}
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
