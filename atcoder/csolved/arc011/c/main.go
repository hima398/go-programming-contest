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

	first, last := nextString(), nextString()
	n := nextInt()
	var s []string
	for i := 0; i < n; i++ {
		s = append(s, nextString())
	}
	ans := solve(first, last, n, s)
	if len(ans) == 0 {
		PrintInt(-1)
		return
	}
	PrintInt(len(ans) - 2)
	PrintVertically(ans)
}

//文字列s, tとで異なる文字数を計測する。s, tの長さは同じ前提
func computeDist(s, t string) int {
	res := 0
	for i := range s {
		if s[i] != t[i] {
			res++
		}
	}
	return res
}

func solve(first, last string, n int, s []string) []string {
	if first == last {
		return []string{first, last}
	}
	nn := n + 2
	e := make([][]int, nn)
	ss := []string{first}
	ss = append(ss, s...)
	ss = append(ss, last)

	for i := 0; i < nn; i++ {
		for j := i + 1; j < nn; j++ {
			if i == j {
				continue
			}
			if computeDist(ss[i], ss[j]) == 1 {
				e[i] = append(e[i], j)
				e[j] = append(e[j], i)
			}
		}
	}
	visited := make([]bool, nn)
	type node struct {
		i int
		s []int
	}
	var q []node
	q = append(q, node{nn - 1, []int{nn - 1}})
	visited[nn-1] = true
	var res []int
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		if cur.i == 0 {
			res = cur.s
			break
		}
		for _, next := range e[cur.i] {
			if visited[next] {
				continue
			}
			ns := make([]int, len(cur.s))
			copy(ns, cur.s)
			ns = append(ns, next)
			q = append(q, node{next, ns})
			visited[next] = true
		}
	}
	var ans []string
	for i := len(res) - 1; i >= 0; i-- {
		ans = append(ans, ss[res[i]])
	}

	return ans
}

func firstsolve(first, last string, n int, s []string) []string {
	if first == last {
		return []string{first, last}
	}
	nn := n + 2
	e := make([][]int, nn)
	ss := []string{first}
	ss = append(ss, s...)
	ss = append(ss, last)

	for i := 0; i < nn; i++ {
		for j := i + 1; j < nn; j++ {
			if i == j {
				continue
			}
			if computeDist(ss[i], ss[j]) == 1 {
				e[i] = append(e[i], j)
				e[j] = append(e[j], i)
			}
		}
	}
	visited := make([]bool, nn)
	var dfs func(cur int) []string
	dfs = func(cur int) []string {
		visited[cur] = true
		if ss[cur] == last {
			visited[cur] = false
			return []string{ss[cur]}
		}
		var res []string
		for _, next := range e[cur] {
			if visited[next] {
				continue
			}
			candidate := dfs(next)
			if len(res) == 0 || (len(candidate) > 0 && len(res) > len(candidate)) {
				res = candidate
			}
		}
		visited[cur] = false
		if len(res) > 0 {
			res = append(res, ss[cur])
		}
		return res
	}
	ans := dfs(0)
	for i := 0; i < len(ans)/2; i++ {
		j := len(ans) - 1 - i
		ans[i], ans[j] = ans[j], ans[i]
	}
	return ans
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

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}

func PrintVertically(x []string) {
	defer out.Flush()
	for _, v := range x {
		fmt.Fprintln(out, v)
	}
}
