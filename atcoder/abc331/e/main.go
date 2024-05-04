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

	n, m, l := nextInt(), nextInt(), nextInt()
	a := nextIntSlice(n)
	b := nextIntSlice(m)
	var c, d []int
	for i := 0; i < l; i++ {
		c = append(c, nextInt()-1)
		d = append(d, nextInt()-1)
	}

	ans := solve(n, m, l, a, b, c, d)

	Print(ans)
}

func solve(n, m, l int, a, b, c, d []int) int {
	//組み合わせNGのマップを作る
	ngMap := make(map[int]map[int]struct{})
	for i := 0; i < l; i++ {
		if ngMap[c[i]] == nil {
			ngMap[c[i]] = make(map[int]struct{})
		}
		ngMap[c[i]][d[i]] = struct{}{}
	}
	type meal struct {
		i, v int
	}
	var mb []meal
	for i, v := range b {
		mb = append(mb, meal{i, v})
	}
	sort.Slice(mb, func(i, j int) bool {
		return mb[i].v > mb[j].v
	})
	var ans int
	for i, v := range a {
		for j := range mb {
			if _, found := ngMap[i][mb[j].i]; found {
				continue
			}
			ans = Max(ans, v+mb[j].v)
			break
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
