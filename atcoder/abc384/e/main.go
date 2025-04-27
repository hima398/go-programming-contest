package main

import (
	"bufio"
	"fmt"
	"os"
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

	h, w, x := nextInt(), nextInt(), nextInt()
	p, q := nextInt()-1, nextInt()-1
	var s [][]int
	for i := 0; i < h; i++ {
		s = append(s, nextIntSlice(w))
	}

	ans := solve(h, w, x, p, q, s)

	Print(ans)
}

func solve(h, w, x, p, q int, s [][]int) int {
	di := []int{-1, 0, 1, 0}
	dj := []int{0, -1, 0, 1}
	type node struct {
		i, j, s int
	}
	queue := priorityqueue.New[node](func(a, b node) int {
		if a.s == b.s {
			return 0
		}
		if a.s < b.s {
			return -1
		}
		return 1
	})
	visited := make([][]bool, h)
	for i := range visited {
		visited[i] = make([]bool, w)
	}
	queue.Push(node{p, q, s[p][q]})
	visited[p][q] = true
	//ans := s[p][q]
	var ans int
	for !queue.Empty() {
		cur := queue.Pop()
		//fmt.Println(cur.s, ans, ans/x)
		var connected bool
		if ans == 0 {
			ans += cur.s
			connected = true
		} else if ans%x == 0 {
			if cur.s < ans/x {
				ans += cur.s
				connected = true
			}
		} else {
			if cur.s <= ans/x {
				ans += cur.s
				connected = true
			}
		}
		if !connected {
			continue
		}
		for k := 0; k < 4; k++ {
			ni, nj := cur.i+di[k], cur.j+dj[k]
			if ni < 0 || ni >= h || nj < 0 || nj >= w {
				continue
			}
			if visited[ni][nj] {
				continue
			}
			queue.Push(node{ni, nj, s[ni][nj]})
			visited[ni][nj] = true

			/*
				if ans%x == 0 {
					if s[ni][nj] < ans/x {
						queue.Push(node{ni, nj, s[ni][nj]})
						ans += s[ni][nj]
						visited[ni][nj] = true
					}
				} else {
					if s[ni][nj] <= ans/x {
						queue.Push(node{ni, nj, s[ni][nj]})
						ans += s[ni][nj]
						visited[ni][nj] = true
					}
				}
			*/
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
