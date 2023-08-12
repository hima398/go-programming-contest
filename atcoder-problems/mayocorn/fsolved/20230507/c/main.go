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

	h, w := nextInt(), nextInt()
	var a [][]int
	for i := 0; i < h; i++ {
		a = append(a, nextIntSlice(w))
	}
	ans := solve(h, w, a)
	PrintInt(ans)
}

func solve(h, w int, a [][]int) int {
	di := []int{0, 1}
	dj := []int{1, 0}
	var dfs func(i, j int, route map[int]int) int
	dfs = func(i, j int, route map[int]int) int {
		if i >= h || j >= w {
			return 0
		}
		route[a[i][j]]++
		if i == h-1 && j == w-1 {
			if len(route) == h+w-1 {
				route[a[i][j]]--
				if route[a[i][j]] == 0 {
					delete(route, a[i][j])
				}
				return 1
			} else {
				route[a[i][j]]--
				if route[a[i][j]] == 0 {
					delete(route, a[i][j])
				}
				return 0
			}
		}
		var res int
		for k := 0; k < 2; k++ {
			ni, nj := i+di[k], j+dj[k]
			res += dfs(ni, nj, route)
		}

		route[a[i][j]]--
		if route[a[i][j]] == 0 {
			delete(route, a[i][j])
		}
		return res
	}
	ans := dfs(0, 0, make(map[int]int))
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

func PrintInt(x int) {
	defer out.Flush()
	fmt.Fprintln(out, x)
}
