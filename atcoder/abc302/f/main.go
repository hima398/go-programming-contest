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
	var a []int
	var s [][]int
	for i := 0; i < n; i++ {
		ai := nextInt()
		a = append(a, ai)
		s = append(s, nextIntSlice(ai))
		for j := range s[i] {
			s[i][j]--
		}
	}
	ans := solve(n, m, a, s)
	PrintInt(ans)
}

func solve(n, m int, a []int, s [][]int) int {
	e := make([][]int, n+m)
	for i := range a {
		for _, sij := range s[i] {
			e[i] = append(e[i], sij+n)
			e[sij+n] = append(e[sij+n], i)
		}
	}
	dist := make([]int, n+m)
	for i := range dist {
		dist[i] = -1
	}
	var q []int
	q = append(q, n)
	dist[n] = 0
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		for _, next := range e[cur] {
			//訪問済み
			if dist[next] >= 0 {
				continue
			}
			q = append(q, next)
			dist[next] = dist[cur] + 1
		}
	}
	//fmt.Println(dist)
	ans := dist[n+m-1]
	if ans < 0 {
		return ans
	}
	return ans/2 - 1
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

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
