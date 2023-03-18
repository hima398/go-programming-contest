package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/emirpasic/gods/trees/avltree"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	//bufサイズ以上の文字列入力が必要な場合は拡張すること
	buf := make([]byte, 9*1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	var k []int
	t := make([][]int, n)
	for i := 0; i < n; i++ {
		k = append(k, nextInt())
		for j := 0; j < k[i]; j++ {
			t[i] = append(t[i], nextInt())
		}
	}
	m := nextInt()
	a := nextIntSlice(m)
	ans := solve(n, k, t, m, a)
	PrintVertically(ans)
}

func solve(n int, k []int, t [][]int, m int, a []int) []int {
	s := make([]*avltree.Tree, 2)
	for i := 0; i < 2; i++ {
		s[i] = avltree.NewWithIntComparator()
	}
	l := make([]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < Min(k[i], 2); j++ {
			s[j].Put(t[i][j], i)
			l[i]++
		}
	}
	var ans []int
	for i := 0; i < m; i++ {
		//消費期限が大きい商品を取る
		n1 := s[0].Right()
		n2 := s[1].Right()

		v1, i1 := n1.Key.(int), n1.Value.(int)
		if a[i] == 1 || n2 == nil || v1 > n2.Key.(int) {
			ans = append(ans, v1)
			s[0].Remove(v1)
			idx := l[i1] - 1
			v := t[i1][idx]
			if _, found := s[1].Get(v); found {
				s[1].Remove(v)
				s[0].Put(v, i1)
			}
			if l[i1] < k[i1] {
				s[1].Put(t[i1][l[i1]], i1)
				l[i1]++
			}
		} else {
			v2, i2 := n2.Key.(int), n2.Value.(int)
			ans = append(ans, v2)
			s[1].Remove(v2)
			if l[i2] < k[i2] {
				s[1].Put(t[i2][l[i2]], i2)
				l[i2]++
			}
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

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
