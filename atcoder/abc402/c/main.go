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

	n, m := nextInt(), nextInt()
	var k []int
	var a [][]int
	for i := 0; i < m; i++ {
		k = append(k, nextInt())
		a = append(a, nextIntSlice(k[i]))
	}
	b := nextIntSlice(n)

	ans := solve(n, m, k, a, b)

	for _, v := range ans {
		Print(v)
	}
}

func solve(n, m int, k []int, a [][]int, b []int) []int {
	for i := range a {
		for j := range a[i] {
			a[i][j]--
		}
	}
	for i := range b {
		b[i]--
	}
	c := make([][]int, n)

	for i := range a {
		for j := range a[i] {
			c[a[i][j]] = append(c[a[i][j]], i)
		}
	}

	ans := make([]int, n)
	ans[n-1] = m

	canNotEat := make([]bool, m)
	for i := n - 1; i > 0; i-- {
		var cnt int
		for _, v := range c[b[i]] {
			//vが食べられなくなった料理の番号
			if canNotEat[v] {
				continue
			}
			canNotEat[v] = true
			cnt++
		}
		ans[i-1] = ans[i] - cnt
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
