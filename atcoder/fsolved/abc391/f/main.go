package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"

	"github.com/liyue201/gostl/ds/priorityqueue"
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
	b := nextIntSlice(n)
	c := nextIntSlice(n)

	ans := solve(n, k, a, b, c)
	//ans := solveHonestly(n, k, a, b, c)

	Print(ans)
}

func Sort(s []int) []int {
	sort.Slice(s, func(i, j int) bool {
		return s[i] > s[j]
	})
	return s
}
func solve(n, k int, a, b, c []int) int {
	a, b, c = Sort(a), Sort(b), Sort(c)

	f := func(i, j, k int) int {
		return a[i]*b[j] + a[i]*c[k] + b[j]*c[k]
	}
	type node struct {
		i, j, k int
		v       int
	}

	q := priorityqueue.New[node](func(a, b node) int {
		if a.v == b.v {
			return 0
		}
		if a.v > b.v {
			return -1
		}
		return 1
	})
	q.Push(node{0, 0, 0, f(0, 0, 0)})

	visited := make(map[int]map[int]map[int]bool)

	var ranking []int
	for len(ranking) < k {
		cur := q.Pop()
		ranking = append(ranking, cur.v)

		ni, nj, nk := cur.i+1, cur.j+1, cur.k+1
		if ni < n {
			if visited[ni] == nil {
				visited[ni] = make(map[int]map[int]bool)
			}
			if visited[ni][cur.j] == nil {
				visited[ni][cur.j] = make(map[int]bool)
			}
			if !visited[ni][cur.j][cur.k] {
				q.Push(node{ni, cur.j, cur.k, f(ni, cur.j, cur.k)})
				visited[ni][cur.j][cur.k] = true
			}

		}
		if nj < n {
			if visited[cur.i] == nil {
				visited[cur.i] = make(map[int]map[int]bool)
			}
			if visited[cur.i][nj] == nil {
				visited[cur.i][nj] = make(map[int]bool)
			}
			if !visited[cur.i][nj][cur.k] {
				q.Push(node{cur.i, nj, cur.k, f(cur.i, nj, cur.k)})
				visited[cur.i][nj][cur.k] = true
			}
		}
		if nk < n {
			if visited[cur.i] == nil {
				visited[cur.i] = make(map[int]map[int]bool)
			}
			if visited[cur.i][cur.j] == nil {
				visited[cur.i][cur.j] = make(map[int]bool)
			}
			if !visited[cur.i][cur.j][nk] {
				q.Push(node{cur.i, cur.j, nk, f(cur.i, cur.j, nk)})
				visited[cur.i][cur.j][nk] = true
			}
		}
	}

	//fmt.Println(ranking)
	ans := ranking[k-1]
	return ans
}

func solveHonestly(n, k int, a, b, c []int) int {
	var s []int
	for x := 0; x < n; x++ {
		for y := 0; y < n; y++ {
			for z := 0; z < n; z++ {
				s = append(s, a[x]*b[y]+b[y]*c[z]+c[z]*a[x])
			}
		}
	}
	sort.Ints(s)
	ans := s[len(s)-k]
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
